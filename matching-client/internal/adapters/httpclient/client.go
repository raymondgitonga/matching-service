package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

const matchingURL = "https:"

type MatchingClient struct {
	client http.Client
}

func NewMatchingClient(client http.Client) (*MatchingClient, error) {
	return &MatchingClient{client: client}, nil
}

func (m MatchingClient) GetPartner(partnerID string) ([]byte, error) {
	partnerURL := fmt.Sprintf("%spartner?id=%s", matchingURL, partnerID)
	resp, err := m.client.Get(partnerURL)

	if err != nil {
		return nil, fmt.Errorf("error making partner call: %w", err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading partner call: %w", err)
	}

	return b, nil
}
