package service

import (
	"encoding/json"
	"io"

	"github.com/raymondgitonga/matching_client/internal/core/dormain"
)

type MatchingClient interface {
	GetPartner(partnerID string) (io.ReadCloser, error)
}

type MatchingPartner struct {
	matchingClient MatchingClient
}

func NewMatchingPartner(matchingClient MatchingClient) *MatchingPartner {
	return &MatchingPartner{matchingClient}
}

func (m *MatchingPartner) GetMatchingPartner(partnerID string) (*dormain.Partner, error) {
	var partner dormain.Partner

	response, err := m.matchingClient.GetPartner(partnerID)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(response).Decode(&partner)
	if err != nil {
		return nil, err
	}

	return &partner, nil
}
