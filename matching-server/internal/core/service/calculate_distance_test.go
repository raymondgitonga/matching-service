package service_test

import (
	"github.com/raymondgitonga/matching-service/internal/core/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_CalculateDistance(t *testing.T) {
	partnerCoordinates := service.NewCoordinates(51.73213, -1.20631)

	testCases := []struct {
		customerCoordinates *service.Coordinates
		expected            float32
	}{
		{
			customerCoordinates: service.NewCoordinates(51.73213, -1.1765110351270853),
			expected:            2.0521753,
		},
		{
			customerCoordinates: service.NewCoordinates(51.73212999999999, -1.1692650258255377),
			expected:            2.5511887,
		},
		{
			customerCoordinates: service.NewCoordinates(53.73213, -1.0877933247235136),
			expected:            222.5329,
		},
	}

	for _, tc := range testCases {
		t.Run("Test distance between coordinates is correct", func(t *testing.T) {
			distance := service.Distance(*tc.customerCoordinates, *partnerCoordinates)

			assert.Equal(t, tc.expected, float32(distance))
		})
	}
}
