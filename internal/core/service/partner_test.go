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

func Test_GetPartnerDetails(t *testing.T) {
	t.Run("Test partner details is mapped correctly to partner dto", func(t *testing.T) {
		specialityMap := map[string]bool{"carpet": true, "tiles": true, "wood": false}
		specByte, _ := json.Marshal(specialityMap)
		repo := &mocks.RepositoryMock{
			GetPartnerFunc: func(ctx context.Context, partnerID int) (*dormain.Partner, error) {
				return &dormain.Partner{
					Name:       "Name",
					Location:   "(51.73212999999999,-1.0831176441976451)",
					Speciality: specByte,
					Radius:     5,
					Rating:     4.5,
				}, nil
			},
		}

		partnerDetails := service.NewPartnerDetails(1, repo)

		partnerDTO, _ := partnerDetails.GetPartnerDetails(context.Background())

		assert.Equal(t, "51.73212999999999,-1.0831176441976451", partnerDTO.Location)
		assert.Equal(t, dormain.Speciality{"carpet", "tiles"}, partnerDTO.Speciality)
	})

}
