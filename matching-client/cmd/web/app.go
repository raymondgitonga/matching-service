package main

import (
	"fmt"
	"github.com/raymondgitonga/matching_client/internal/adapters/httpclient"

	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching_client/internal/adapters/httpserver"
)

type AppConfigs struct {
	baseURL  string
	matchURL string
}

func NewAppConfigs(baseURL string, matchURL string) *AppConfigs {
	return &AppConfigs{baseURL: baseURL, matchURL: matchURL}
}
func (c *AppConfigs) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	handler := httpserver.NewHandler(httpclient.NewConfig(c.matchURL))
	r.HandleFunc(fmt.Sprintf("%s/health-check", c.baseURL), handler.HealthCheck)
	r.HandleFunc(fmt.Sprintf("%s/partner", c.baseURL), handler.GetPartnerDetails)

	return r, nil
}
