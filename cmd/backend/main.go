package main

import (
	"context"
	"log"

	"1-task/internal/backend"
	"1-task/internal/envutil"
	"1-task/internal/storage/postgres"
)

func main() {
	cfg := loadConfig()

	repo, err := postgres.New(context.Background(), cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("init postgres repository: %v", err)
	}
	defer func() {
		if err := repo.Close(); err != nil {
			log.Printf("close postgres repository: %v", err)
		}
	}()

	server := backend.NewServer(repo)

	log.Printf("backend listening on %s", cfg.HTTPAddr)
	if err := server.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}

type config struct {
	HTTPAddr    string
	PostgresDSN string
}

func loadConfig() config {
	return config{
		HTTPAddr:    envutil.Get("HTTP_ADDR", ":8888"),
		PostgresDSN: envutil.Get("POSTGRES_DSN", "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable"),
	}
}
