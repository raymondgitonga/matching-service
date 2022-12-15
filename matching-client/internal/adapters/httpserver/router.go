package httpserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/raymondgitonga/matching_client/internal/adapters/httpclient"
	"github.com/raymondgitonga/matching_client/internal/core/dormain"
	"github.com/raymondgitonga/matching_client/internal/core/service"
)

type Handler struct {
	*httpclient.Config
}

func NewHandler(httpClient *httpclient.Config) *Handler {
	return &Handler{httpClient}
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
	partnerID := r.URL.Query().Get("id")
	client := http.Client{}

	matchingCLient := httpclient.NewMatchingClient(client, *h.Config)
	partnerService := service.NewMatchingPartner(matchingCLient)

	partners, err := partnerService.GetMatchingPartner(partnerID)
	if err != nil {
		processResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	processResponse(w, partners.Result, nil, http.StatusOK)
}

func processResponse(w http.ResponseWriter, results []dormain.Result, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	var response dormain.Partner

	if err != nil {
		response = dormain.Partner{
			Message: err.Error(),
			Result:  []dormain.Result{},
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("error marshaling response: %s", err)
		}

		w.WriteHeader(status)
		_, _ = w.Write(jsonResponse)
		return
	}

	response = dormain.Partner{
		Message: "success",
		Result:  results,
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatalf("error marshaling response: %s", err)
		return
	}
	_, _ = w.Write(jsonResponse)
}
