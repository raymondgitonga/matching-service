package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/adapters/httpserver"
	"net/http"
)

func StartApp() {
	r := mux.NewRouter()

	dbClient, err := db.NewClient(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	handler := httpserver.Handler{DB: dbClient}

	r.HandleFunc("/health_check", handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/partner", handler.GetPartnerDetails).Methods(http.MethodGet)

	fmt.Printf("starting server on :8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
		return
	}
}
