package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/adapters/httpserver"
	"log"
	"net/http"
	"os"
)

func StartApp() {
	r := mux.NewRouter()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbURL := os.Getenv("DB_CONNECTION_URL")
	dbName := os.Getenv("DB_NAME")
	baseURL := os.Getenv("BASE_URL")
	dbClient, err := db.NewClient(context.Background(), dbURL)

	err = db.RunMigrations(dbClient, dbName)
	if err != nil {
		log.Fatal("Error running migration")
		return
	}

	handler := httpserver.Handler{DB: dbClient}
	r.HandleFunc(fmt.Sprintf("%s/health-check", baseURL), handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partner", baseURL), handler.GetPartnerDetails).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partners", baseURL), handler.GetMatchingPartners).Methods(http.MethodPost)

	fmt.Printf("starting server on :8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("error starting server: %s", err)
		return
	}
}
