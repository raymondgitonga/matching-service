package httpclient_test

import (
	"github.com/raymondgitonga/matching_client/internal/adapters/httpclient"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMatchingClient_GetPartner(t *testing.T) {

	t.Run("Successful matching call, partner returned", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			_, err := rw.Write([]byte(`{
							"result": [
								{
									"name": "Cummerata, Wolff and Hauck",
									"location": "51.73212999999999,-1.0831176441976451",
									"material": [
										"tiles",
										"carpet"
									],
									"radius": 1,
									"rating": 4.5
								}
							],
							"error": false,
							"message": "success"
						}`))
			assert.NoError(t, err)
		}))
		defer server.Close()

		config := httpclient.NewConfig(server.URL)
		matchingClient := httpclient.NewMatchingClient(*server.Client(), *config)
		partner, err := matchingClient.GetPartner("1")

		assert.NoError(t, err)
		assert.Equal(t, 1, len(partner.Result))
		assert.Equal(t, "Cummerata, Wolff and Hauck", partner.Result[0].Name)
	})

	t.Run("Successful matching call, no partner returned", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			_, err := rw.Write([]byte(`{
							"result": [],
							"error": false,
							"message": "success"
						}`))
			assert.NoError(t, err)
		}))

		defer server.Close()

		config := httpclient.NewConfig(server.URL)
		matchingClient := httpclient.NewMatchingClient(*server.Client(), *config)
		partner, err := matchingClient.GetPartner("1")

		assert.NoError(t, err)
		assert.Equal(t, 0, len(partner.Result))
	})

	t.Run("Matching request BadRequest", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusBadRequest)
			_, err := rw.Write([]byte(`{
							"result": [],
							"error": true,
							"message": "malformed request"
						}`))
			assert.NoError(t, err)
		}))

		defer server.Close()

		config := httpclient.NewConfig(server.URL)
		matchingClient := httpclient.NewMatchingClient(*server.Client(), *config)
		partner, err := matchingClient.GetPartner("1")

		assert.Error(t, err)
		assert.Nil(t, partner)
	})

	t.Run("Matching request InternalServerError", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusInternalServerError)
			_, err := rw.Write([]byte(``))
			assert.NoError(t, err)
		}))

		defer server.Close()

		config := httpclient.NewConfig(server.URL)
		matchingClient := httpclient.NewMatchingClient(*server.Client(), *config)
		partner, err := matchingClient.GetPartner("1")

		assert.Error(t, err)
		assert.Nil(t, partner)
	})
}
