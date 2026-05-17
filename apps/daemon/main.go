package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"1-task/internal/envutil"
	"1-task/internal/indexer"
	"1-task/internal/storage/postgres"

	daemon "github.com/sevlyar/go-daemon"
)

func main() {
	if os.Getenv("DAEMON_FOREGROUND") == "1" {
		runLoop()
		return
	}

	cntxt := &daemon.Context{
		PidFileName: envutil.Get("DAEMON_PID_FILE", "./tmp/umbrella-daemon.pid"),
		PidFilePerm: 0o644,
		LogFileName: envutil.Get("DAEMON_LOG_FILE", "./tmp/umbrella-daemon.log"),
		LogFilePerm: 0o640,
		WorkDir:     "./",
		Umask:       0o027,
		Args:        []string{os.Args[0], "-daemon"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatalf("failed to daemonize: %v", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	runLoop()
}

func runLoop() {
	ctx := context.Background()

	repo, svc, err := initIndexer(ctx)
	if err != nil {
		log.Fatalf("init indexer: %v", err)
	}
	defer svc.Close()
	defer func() {
		if err := repo.Close(); err != nil {
			log.Printf("close postgres repository: %v", err)
		}
	}()

	log.Print("daemon started")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(sigs)

	// First run is unbounded so historical catch-up is not interrupted by timeout.
	runIndexCycle(svc, true)

	ticker := time.NewTicker(12 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			runIndexCycle(svc, false)
		case s := <-sigs:
			log.Printf("signal received: %s", s)
			log.Print("daemon stopped")
			return
		}
	}
}

func runIndexCycle(svc *indexer.Service, isFirstRun bool) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	if isFirstRun {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	}
	defer cancel()

	processed, err := svc.RunCycle(ctx, cancel)
	if err != nil {
		log.Printf("indexer cycle failed: %v", err)
		return
	}
	if processed {
		log.Print("indexer cycle completed: logs processed")
		return
	}
	log.Print("indexer cycle completed: no new finalized logs")
}

func initIndexer(ctx context.Context) (*postgres.Repository, *indexer.Service, error) {
	rpcURL := envutil.Get("RPC_URL", "https://ethereum-rpc.publicnode.com")
	dsn := envutil.Get("POSTGRES_DSN", "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable")
	configPath := envutil.Get("UMBRELLA_CONFIG_PATH", "./configs/umbrella/mainnet.json")
	batchRange := uint64(2000)
	finalityDepth := uint64(12)

	cfg, err := indexer.LoadConfig(configPath, rpcURL, batchRange, finalityDepth)
	if err != nil {
		return nil, nil, fmt.Errorf("load indexer config: %w", err)
	}

	repo, err := postgres.New(ctx, dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("init postgres repository: %w", err)
	}

	svc, err := indexer.NewService(ctx, cfg, repo)
	if err != nil {
		_ = repo.Close()
		return nil, nil, fmt.Errorf("init indexer service: %w", err)
	}

	return repo, svc, nil
}
