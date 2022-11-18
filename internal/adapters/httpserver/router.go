package httpserver

import (
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
	dB *sql.DB
}

func NewHandler(db *sql.DB) (*Handler, error) {
	if db == nil {
		return nil, fmt.Errorf("db is null")
	}
	return &Handler{
		db,
	}, nil
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

	partnerRepo, err := repository.NewPartnerRepository(h.dB)
	if err != nil {
		processResponse(w, nil, err, http.StatusInternalServerError)
	}

	partnerService := service.NewPartnerService(partnerRepo)
	partner, err := partnerService.GetPartnerDetails(r.Context(), partnerID)
	if err != nil {
		processResponse(w, nil, err, http.StatusInternalServerError)
	}

	partners = append(partners, *partner)
	processResponse(w, partners, nil, http.StatusOK)
}

func (h *Handler) GetMatchingPartners(w http.ResponseWriter, r *http.Request) {
	var request dormain.CustomerRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, fmt.Errorf("error decoding body: %w", err), http.StatusBadRequest)
	}

	err = ValidateCustomerRequest(request)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, err, http.StatusBadRequest)
		return
	}

	partnerRepo, err := repository.NewPartnerRepository(h.dB)
	if err != nil {
		processResponse(w, nil, err, http.StatusInternalServerError)
	}

	partnerService := service.NewPartnerService(partnerRepo)
	partners, err := partnerService.GetMatchingPartners(r.Context(), request)
	if err != nil {
		processResponse(w, []dormain.PartnerDTO{}, err, http.StatusInternalServerError)
		return
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
			Result:  []dormain.PartnerDTO{},
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("error marshaling response: %s", err)
		}

		w.WriteHeader(status)
		_, _ = w.Write(jsonResponse)
		return
	}

	response = dormain.PartnerResponse{
		Error:   false,
		Message: "success",
		Result:  partner,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("error marshaling response: %s", err)
		return
	}
	_, _ = w.Write(jsonResponse)
}
