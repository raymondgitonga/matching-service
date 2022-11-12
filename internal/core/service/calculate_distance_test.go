package service

import (
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CalculateDistance(t *testing.T) {
	partnerCoordinates := NewCoordinates(51.73213, -1.20631)

	testCases := []struct {
		customerCoordinates dormain.Coordinates
		expected            int
	}{
		{
			customerCoordinates: dormain.Coordinates{Latitude: 51.73213, Longitude: -1.1765110351270853},
			expected:            2052,
		},
		{
			customerCoordinates: dormain.Coordinates{Latitude: 51.73212999999999, Longitude: -1.1692650258255377},
			expected:            2551,
		},
		{
			customerCoordinates: dormain.Coordinates{Latitude: 53.73213, Longitude: -1.0877933247235136},
			expected:            222532,
		},
	}

	for _, tc := range testCases {
		t.Run("Test distance between coordinates is correct", func(t *testing.T) {
			distance := distance(tc.customerCoordinates, *partnerCoordinates)
			assert.Equal(t, tc.expected, int(distance))
		})
	}
}
