package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/adapters/handler"
	"net/http"
)

func StartApp() {
	r := mux.NewRouter()

	r.HandleFunc("/health_check", handler.HealthCheck).Methods(http.MethodGet)

	fmt.Printf("starting server on :8080")

	_, err := db.NewClient(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Printf("error running migration: %s", err)
		return
	}

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
		return
	}
}
