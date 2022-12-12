package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const matchingURL = "https:"

type MatchingClient struct {
	client http.Client
}

func NewMatchingClient(client http.Client) (*MatchingClient, error) {
	return &MatchingClient{client: client}, nil
}

func (m MatchingClient) GetPartner(partnerID string) (io.ReadCloser, error) {
	partnerURL := strings.TrimSpace(fmt.Sprintf("%spartner?id=%s", matchingURL, partnerID))
	resp, err := m.client.Get(partnerURL)
	if err != nil {
		return nil, fmt.Errorf("error making partner call: %w", err)
	}
	return resp.Body, err
}
