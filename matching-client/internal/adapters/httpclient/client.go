package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const matchingURL = "https://dadf-188-80-46-56.eu.ngrok.io/"

type Config struct {
	matchingURL string
}

type MatchingClient struct {
	client http.Client
	config Config
}

func NewConfig() *Config {
	return &Config{os.Getenv("MATCHING_URL")}
}
func NewMatchingClient(client http.Client, config Config) (*MatchingClient, error) {
	return &MatchingClient{client: client, config: config}, nil
}

func (m MatchingClient) GetPartner(partnerID string) (*Partner, error) {
	var partner Partner
	partnerURL := fmt.Sprintf("%smatching-service/partner?id=%s", matchingURL, partnerID)
	resp, err := m.client.Get(partnerURL)

	if err != nil {
		return nil, fmt.Errorf("error making partner call: %w", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&partner)
	if err != nil {
		return nil, fmt.Errorf("error decoding partner call: %w", err)
	}

	if partner.Error {
		return nil, fmt.Errorf("error reading partner call: %s", partner.Message)
	}

	return &partner, nil
}
