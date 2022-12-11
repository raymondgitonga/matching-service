package cmd

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raymondgitonga/matching_integration/internal/adapters/httpserver"
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

	return r, nil
}
