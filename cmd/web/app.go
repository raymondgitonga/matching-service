package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching-service/internal/adapters/db"
	"github.com/raymondgitonga/matching-service/internal/adapters/httpserver"
	"net/http"
)

type AppConfigs struct {
	dbURL   string
	dbName  string
	baseURL string
}

func NewAppConfigs(dbURL, dbName, baseURL string) *AppConfigs {
	return &AppConfigs{
		dbURL:   dbURL,
		dbName:  dbName,
		baseURL: baseURL,
	}
}
func (c *AppConfigs) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	baseURL := c.baseURL
	dbClient, err := db.NewClient(context.Background(), c.dbURL)

	err = db.RunMigrations(dbClient, c.dbName)
	if err != nil {
		return nil, fmt.Errorf("error running migration: %w", err)
	}

	handler := httpserver.NewHandler(dbClient)
	r.HandleFunc(fmt.Sprintf("%s/health-check", baseURL), handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partner", baseURL), handler.GetPartnerDetails).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partners", baseURL), handler.GetMatchingPartners).Methods(http.MethodPost)

	fmt.Printf("starting server on :8080")

	return r, nil
}
