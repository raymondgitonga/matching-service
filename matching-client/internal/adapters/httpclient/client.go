package httpclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	matchingURL string
}

type MatchingClient struct {
	client http.Client
	config Config
}

func NewConfig(matchURL string) *Config {
	return &Config{matchingURL: matchURL}
}

func NewMatchingClient(client http.Client, config Config) *MatchingClient {
	return &MatchingClient{client: client, config: config}
}

func (m MatchingClient) GetPartner(partnerID string) (*Partner, error) {
	var partner Partner
	partnerURL := fmt.Sprintf("%s/matching-service/partner?id=%s", m.config.matchingURL, partnerID)
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
		fmt.Println(partner.Error)
		return nil, fmt.Errorf("error reading partner call: %s", partner.Message)
	}

	return &partner, nil
}
