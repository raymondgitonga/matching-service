package service

import (
	"encoding/json"
	"fmt"
	"github.com/raymondgitonga/matching_client/internal/core/dormain"
)

type MatchingClient interface {
	GetPartner(partnerID string) ([]byte, error)
}

type PartnerService struct {
	matchingClient MatchingClient
}

func NewMatchingPartner(matchingClient MatchingClient) *PartnerService {
	return &PartnerService{matchingClient}
}

func (m *PartnerService) GetMatchingPartner(partnerID string) (*dormain.Partner, error) {
	fmt.Println(partnerID)
	var partner dormain.Partner

	response, err := m.matchingClient.GetPartner(partnerID)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response, &partner)
	if err != nil {
		return nil, err
	}
	return &partner, nil
}
