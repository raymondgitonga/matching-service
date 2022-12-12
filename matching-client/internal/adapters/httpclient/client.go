package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

const matchingURL = "localhost:8080/matching-service/partner"

type Matching struct {
	client *http.Client
}

func NewMatchingClient(client *http.Client) (*Matching, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is nil")
	}
	return &Matching{client: client}, nil
}

func (m Matching) GetPartner(partnerID string) (io.ReadCloser, error) {
	resp, err := m.client.Get(fmt.Sprintf("%s?%s", matchingURL, partnerID))
	if err != nil {
		return nil, fmt.Errorf("error making partner call: %w", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error closing body: %w", err)
	}
	return resp.Body, err
}
