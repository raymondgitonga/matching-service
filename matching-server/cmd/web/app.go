package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching-server/internal/adapters/db"
	"github.com/raymondgitonga/matching-server/internal/adapters/httpserver"
)

type AppConfigs struct {
	dbURL   string
	dbName  string
	baseURL string
}

func NewAppConfigs(dbURL, dbName, baseURL string) (*AppConfigs, error) {
	if dbURL == "" {
		return nil, fmt.Errorf("kindly provide dbURL")
	}
	if dbName == "" {
		return nil, fmt.Errorf("kindly provide dbName")
	}
	if baseURL == "" {
		return nil, fmt.Errorf("kindly provide baseURL")
	}
	return &AppConfigs{
		dbURL:   dbURL,
		dbName:  dbName,
		baseURL: baseURL,
	}, nil
}
func (c *AppConfigs) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	baseURL := c.baseURL
	dbClient, err := db.NewClient(context.Background(), c.dbURL)
	if err != nil {
		return nil, fmt.Errorf("error running migration: %w", err)
	}

	err = db.RunMigrations(dbClient, c.dbName)
	if err != nil {
		return nil, fmt.Errorf("error running migration: %w", err)
	}

	handler, err := httpserver.NewHandler(dbClient)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	r.HandleFunc(fmt.Sprintf("%s/health-check", baseURL), handler.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partner", baseURL), handler.GetPartnerDetails).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("%s/partners", baseURL), handler.GetMatchingPartners).Methods(http.MethodPost)

	return r, nil
}
