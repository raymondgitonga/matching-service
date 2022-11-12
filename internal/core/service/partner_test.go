package service_test

import (
	"context"
	"encoding/json"
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"github.com/raymondgitonga/matching-service/internal/core/service"
	"github.com/raymondgitonga/matching-service/internal/core/service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetPartnerDetails(t *testing.T) {
	t.Run("Test partner details is mapped correctly to partner dto", func(t *testing.T) {
		materialMap := map[string]bool{"carpet": true, "tiles": true, "wood": false}
		materialByte, _ := json.Marshal(materialMap)
		repo := &mocks.RepositoryMock{
			GetPartnerFunc: func(ctx context.Context, partnerID int) (*dormain.Partner, error) {
				return &dormain.Partner{
					Name:     "Name",
					Location: "(51.73212999999999,-1.0831176441976451)",
					Material: materialByte,
					Radius:   5,
					Rating:   4.5,
				}, nil
			},
		}

		partnerDetails := service.NewPartnerService(repo)
		partnerDTO, err := partnerDetails.GetPartnerDetails(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, "51.73212999999999,-1.0831176441976451", partnerDTO.Location)
		assert.Equal(t, []string{"tiles", "carpet"}, partnerDTO.Material)
	})
}

func TestService_GetMatchingPartners(t *testing.T) {
	customerLat := 51.73213
	customerLon := -1.156285162957502

	request := dormain.CustomerRequest{
		Material: "wood",
		Lat:      customerLat,
		Long:     customerLon,
	}

	t.Run("Test partner details are being filtered and sorted correctly", func(t *testing.T) {
		materialMap := map[string]bool{"carpet": true, "tiles": true, "wood": false}
		specByte, err := json.Marshal(materialMap)
		assert.NoError(t, err)

		repo := &mocks.RepositoryMock{
			GetPartnersFunc: func(ctx context.Context, material string) (*[]dormain.Partner, error) {
				return &[]dormain.Partner{
					{
						Name:     "Business1",
						Location: "(51.73212999999999,-1.0831176441976451)",
						Material: specByte,
						Radius:   5,
						Rating:   5.0,
					},
					{
						Name:     "Business2",
						Location: "(51.73213,-1.0877933247235136)",
						Material: specByte,
						Radius:   1,
						Rating:   4.5,
					},
					{
						Name:     "Business3",
						Location: "(51.73213,-1.156285162957502)",
						Material: specByte,
						Radius:   1,
						Rating:   3.5,
					},
				}, nil
			},
		}

		partnerDetails := service.NewPartnerService(repo)
		partnersDTO, err := partnerDetails.GetMatchingPartners(context.Background(), request)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(*partnersDTO), "Only two partners match")
		assert.Equal(t, float64(5), (*partnersDTO)[0].Rating, "Partner with the highest rating shown first")
	})

	t.Run("Test location unmarshalling error being caught", func(t *testing.T) {
		materialMap := map[string]bool{"carpet": true, "tiles": true, "wood": false}
		specByte, err := json.Marshal(materialMap)
		assert.NoError(t, err)

		repo := &mocks.RepositoryMock{
			GetPartnersFunc: func(ctx context.Context, material string) (*[]dormain.Partner, error) {
				return &[]dormain.Partner{
					{
						Name:     "Business1",
						Location: "(51.73212999999999, -1.0831176441976451)",
						Material: specByte,
						Radius:   5,
						Rating:   5.0,
					},
				}, nil
			},
		}

		partnerDetails := service.NewPartnerService(repo)
		partnersDTO, err := partnerDetails.GetMatchingPartners(context.Background(), request)
		assert.Error(t, err)
		assert.Nil(t, partnersDTO)
	})
}

func TestService_ComputeDistance(t *testing.T) {
	customerLat := 51.73213
	customerLon := -1.1116500381594543

	t.Run("Distance computed correctly", func(t *testing.T) {
		partnerLocation := "51.73213,-1.12645200882915"
		distance, err := service.ComputeDistance(partnerLocation, customerLat, customerLon)

		assert.Equal(t, 1, distance)
		assert.NoError(t, err)
	})

	t.Run("Distance computation failed due to marshalling error", func(t *testing.T) {
		partnerLocation := "51.73213, -1.12645200882915"
		distance, err := service.ComputeDistance(partnerLocation, customerLat, customerLon)

		assert.Equal(t, -1, distance)
		assert.Error(t, err)
	})

}
