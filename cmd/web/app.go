package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/adapters/handler"
	"github.com/raymondgitonga/matching-service/internal/core/repository"
	"net/http"
)

func StartApp() {
	r := mux.NewRouter()

	r.HandleFunc("/health_check", handler.HealthCheck).Methods(http.MethodGet)

	fmt.Printf("starting server on :8080")

	dbClient, err := db.NewClient(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Printf("error running migration: %s", err)
		return
	}

	repo := repository.NewRepository(dbClient)

	x, err := repo.GetPartner(context.Background(), 1)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println(x.Location[1 : len(x.Location)-1])

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
		return
	}
}
