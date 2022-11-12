package httpserver

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"github.com/raymondgitonga/matching-service/internal/core/repository"
	"github.com/raymondgitonga/matching-service/internal/core/service"
	"log"
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
	partners := make([]dormain.PartnerDTO, 0)
	ID := r.URL.Query().Get("id")

	partnerID, err := strconv.Atoi(ID)
	if err != nil {
		processResponse(w, nil, err, http.StatusBadRequest)
	}

	partnerService := service.NewPartnerService(repository.NewPartnerRepository(h.DB))
	partner, err := partnerService.GetPartnerDetails(context.Background(), partnerID)
	if err != nil {
		processResponse(w, nil, err, http.StatusInternalServerError)
	}

	partners = append(partners, *partner)
	processResponse(w, partners, nil, http.StatusOK)
}

func (h *Handler) GetMatchingPartners(w http.ResponseWriter, r *http.Request) {
	speciality := r.URL.Query().Get("speciality")
	latitude, err := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, err, http.StatusBadRequest)
	}
	longitude, err := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, err, http.StatusBadRequest)
	}

	partnerService := service.NewPartnerService(repository.NewPartnerRepository(h.DB))
	partners, err := partnerService.GetMatchingPartners(context.Background(), speciality, latitude, longitude)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, err, http.StatusInternalServerError)
	}

	processResponse(w, *partners, nil, http.StatusOK)
}

func processResponse(w http.ResponseWriter, partner []dormain.PartnerDTO, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	var response dormain.PartnerResponse

	if err != nil {
		response = dormain.PartnerResponse{
			Error:   true,
			Message: err.Error(),
			Result:  nil,
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("error marshaling response: %s", err)
		}

		w.WriteHeader(status)
		_, _ = w.Write(jsonResponse)
	}

	response = dormain.PartnerResponse{
		Error:   false,
		Message: "success",
		Result:  partner,
	}

	jsonResponse, _ := json.Marshal(response)

	w.WriteHeader(status)
	_, _ = w.Write(jsonResponse)
}
