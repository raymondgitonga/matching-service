package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching_client/internal/adapters/httpserver"
)

type AppConfigs struct {
	baseURL string
}

func NewAppConfigs(baseURL string) *AppConfigs {
	return &AppConfigs{baseURL: baseURL}
}
func (c *AppConfigs) StartApp() (*mux.Router, error) {
	r := mux.NewRouter()
	handler := httpserver.Handler{}
	r.HandleFunc(fmt.Sprintf("%s/health-check", c.baseURL), handler.HealthCheck)
	r.HandleFunc(fmt.Sprintf("%s/partner", c.baseURL), handler.GetPartnerDetails)

	return r, nil
}
