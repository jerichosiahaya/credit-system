package main

import (
	"context"
	"kredit-plus/src/config"
	delivery "kredit-plus/src/credit/delivery"
	postgresRepository "kredit-plus/src/credit/repository/postgres"
	usecase "kredit-plus/src/credit/usecase"

	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    ctx := context.Background()
    router := chi.NewRouter()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgxCfg, err := pgxpool.ParseConfig(config.PostgresUri)
	if err != nil {
		panic(err)
	}
	pgxCfg.MaxConns = 200
	pgPool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		panic(err)
	}

	db := postgresRepository.New(pgPool)
	uc := usecase.New(db)
	router = delivery.New(router, ctx, uc)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Default().Println("Server running on port :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Default().Fatal("Failed to start server: ", err)
	}
}