package service_test

import (
	"fmt"
	"github.com/raymondgitonga/matching_client/internal/adapters/httpclient"
	"github.com/raymondgitonga/matching_client/internal/core/service"
	"github.com/raymondgitonga/matching_client/internal/core/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartnerService_GetMatchingPartner(t *testing.T) {
	t.Run("Test Matching partner returned successfully", func(t *testing.T) {
		matchingClient := &mocks.MatchingClientMock{
			GetPartnerFunc: func(parnterID string) (*httpclient.Partner, error) {
				return &httpclient.Partner{
					Result: []httpclient.Result{
						{
							Name:     "Cummerata, Wolff and Hauck",
							Location: "51.73212999999999,-1.0831176441976451",
							Material: []string{"tiles", "carpet"},
							Radius:   1,
							Rating:   4.5,
						},
					},
					Error:   false,
					Message: "success",
				}, nil
			},
		}
		partnerService := service.NewMatchingPartner(matchingClient)
		partner, err := partnerService.GetMatchingPartner("1")

		assert.NoError(t, err)
		assert.Equal(t, 1, len(partner.Result))
		assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Result[0].Name)
	})

	t.Run("Test Matching partner fails", func(t *testing.T) {
		matchingClient := &mocks.MatchingClientMock{
			GetPartnerFunc: func(parnterID string) (*httpclient.Partner, error) {
				return nil, fmt.Errorf("error reading partner call: No partner found")
			},
		}
		partnerService := service.NewMatchingPartner(matchingClient)
		partner, err := partnerService.GetMatchingPartner("1")

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "error reading partner call: No partner found")
		assert.Nil(t, partner)
	})

	t.Run("Test Matching no partner returned", func(t *testing.T) {
		matchingClient := &mocks.MatchingClientMock{
			GetPartnerFunc: func(parnterID string) (*httpclient.Partner, error) {
				return &httpclient.Partner{
					Result:  []httpclient.Result{},
					Error:   false,
					Message: "success",
				}, nil
			},
		}
		partnerService := service.NewMatchingPartner(matchingClient)
		partner, err := partnerService.GetMatchingPartner("1")

		assert.Nil(t, partner)
		assert.Equal(t, "no results found", err.Error())
	})
}
