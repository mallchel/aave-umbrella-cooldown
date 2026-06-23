package main

import (
	"context"
	"log"

	"1-task/internal/appconfig"
	"1-task/internal/backend"
	"1-task/internal/storage/postgres"
)

func main() {
	cfg := appconfig.LoadServiceConfig()

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

	errCh := make(chan error, 2)

	log.Printf("backend API listening on %s", cfg.HTTPAddr)
	go func() {
		errCh <- server.Run(cfg.HTTPAddr)
	}()

	log.Printf("swagger UI listening on %s", cfg.SwaggerAddr)
	go func() {
		errCh <- server.RunSwaggerUI(cfg.SwaggerAddr)
	}()

	log.Fatal(<-errCh)
}
