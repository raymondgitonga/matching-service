package httpserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/repository"
	"github.com/raymondgitonga/matching-service/internal/core/service"
	"net/http"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	response, err := json.Marshal("Healthy")
	if err != nil {
		fmt.Printf("error writing marshalling response: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		fmt.Printf("error writing httpserver response: %s", err)
	}
}

func (h *Handler) GetPartnerDetails(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	partnerID, err := strconv.Atoi(ID)
	if err != nil {
		// do something
	}

	if err != nil {
		// do something
	}

	partnerService := service.NewPartnerDetails(repository.NewRepository(h.DB))
	partner, err := partnerService.GetPartnerDetails(context.Background(), partnerID)
	if err != nil {
		// do something
	}

	response, err := json.Marshal(partner)
	if err != nil {
		// do something
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		fmt.Printf("error writing httpserver response: %s", err)
	}
}

func (h *Handler) GetPartners(w http.ResponseWriter, r *http.Request) {
	speciality := r.URL.Query().Get("speciality")
	lat, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		// do something
	}
	lon, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		// do something
	}

	partnerService := service.NewPartnerDetails(repository.NewRepository(h.DB))
	partners, err := partnerService.GetMatchingPartners(context.Background(), speciality, lat, lon)
	if err != nil {
		// do something
	}

	response, err := json.Marshal(*partners)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		fmt.Printf("error writing httpserver response: %s", err)
	}
}
