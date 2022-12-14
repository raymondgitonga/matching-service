package service

import (
	"fmt"
	"github.com/raymondgitonga/matching_client/internal/adapters/httpclient"
	"github.com/raymondgitonga/matching_client/internal/core/dormain"
)

type MatchingClient interface {
	GetPartner(partnerID string) (*httpclient.Partner, error)
}

type PartnerService struct {
	matchingClient MatchingClient
}

func NewMatchingPartner(matchingClient MatchingClient) *PartnerService {
	return &PartnerService{matchingClient}
}

func (m *PartnerService) GetMatchingPartner(partnerID string) (*dormain.Partner, error) {
	var partner dormain.Partner
	resp, err := m.matchingClient.GetPartner(partnerID)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if len(resp.Result) < 1 {
		return nil, fmt.Errorf("no results found")
	}
	result := resp.Result[0]
	partner = dormain.Partner{
		Result: []dormain.Result{
			{
				Name:     result.Name,
				Location: result.Location,
				Material: result.Material,
				Radius:   result.Radius,
				Rating:   result.Rating,
			},
		},
		Message: resp.Message,
	}

	return &partner, nil
}
