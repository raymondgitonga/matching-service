package httpserver

import (
	"github.com/raymondgitonga/matching-service/internal/core/dormain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_ValidateCustomerRequest(t *testing.T) {
	t.Run("Test valid phone number", func(t *testing.T) {
		testCases := []struct {
			desc        string
			phoneNumber string
			valid       bool
		}{
			{
				phoneNumber: "(+351) 282 43 50 50",
				valid:       true,
			},
			{
				phoneNumber: "90191919908",
				valid:       true,
			},
			{
				phoneNumber: "001 6867684",
				valid:       true,
			},
			{
				phoneNumber: "0729320243",
				valid:       true,
			},
			{
				phoneNumber: "2547293u20243",
				valid:       false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				request := dormain.CustomerRequest{Phone: tc.phoneNumber, Material: "tiles"}
				err := ValidateCustomerRequest(request)
				assert.Equal(t, tc.valid, err == nil)
			})
		}
	})

	t.Run("Test valid material type", func(t *testing.T) {
		testCases := []struct {
			desc     string
			material string
			valid    bool
		}{
			{
				material: "tiles",
				valid:    true,
			},
			{
				material: "wood",
				valid:    true,
			},
			{
				material: "carpet",
				valid:    true,
			},
			{
				material: "tile",
				valid:    false,
			},
			{
				material: "ceiling",
				valid:    false,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.desc, func(t *testing.T) {
				request := dormain.CustomerRequest{Material: tc.material}
				err := ValidateCustomerRequest(request)
				assert.Equal(t, tc.valid, err == nil)
			})
		}
	})
}
