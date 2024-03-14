package adapter_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/martoc/ipstack-cli/adapter"
	"github.com/stretchr/testify/assert"
)

type mockRoundTripper struct {
	handler func(req *http.Request) (*http.Response, error)
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.handler(req)
}

func TestIPStackServiceAdapter_GetCoordinates_RoundTripper(t *testing.T) {
	t.Parallel()

	// Test case 1: Valid IP and access key
	validIP := "validIP"
	validAccessKey := "validAccessKey"
	expectedCoordinates := "latitude,longitude"

	// Create a round tripper with a custom handler
	roundTripper := &mockRoundTripper{
		handler: func(req *http.Request) (*http.Response, error) {
			assert.Equal(t, validAccessKey, req.URL.Query().Get("access_key"))

			resp := &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(expectedCoordinates)),
			}

			return resp, nil
		},
	}

	httpClient := &http.Client{Transport: roundTripper}

	// Set the round tripper as the transport for the adapter
	ipstackService := adapter.NewIPStackServiceAdapterBuilder().SetClient(httpClient).Build()

	coordinates, err := ipstackService.GetCoordinates(context.Background(), validIP, validAccessKey)
	assert.NoError(t, err)
	assert.Equal(t, expectedCoordinates, coordinates)
}
