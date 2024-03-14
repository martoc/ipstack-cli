package adapter

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const host = "api.ipstack.com"

var (
	errIndalidIP        = fmt.Errorf("ip address cannot be empty")
	errInvalidAccessKey = fmt.Errorf("access key cannot be empty")
)

// IPStackService represents a service for retrieving coordinates based on an IP address.
type IPStackService interface {
	GetCoordinates(ctx context.Context, ip string, accessKey string) (string, error)
}

// IPStackServiceAdapter is an implementation of the IPStackService interface.
// It provides methods for retrieving coordinates based on an IP address using the IPStack API.
type IPStackServiceAdapter struct {
	IPStackService IPStackService
	client         *http.Client
	host           string
}

// IPStackServiceAdapterBuilder is a builder for creating instances of IPStackServiceAdapter.
type IPStackServiceAdapterBuilder struct {
	client *http.Client
	host   string
}

// SetClient sets the HTTP client to be used by the IPStackServiceAdapter.
// It takes a pointer to an http.Client as input and returns a pointer to the IPStackServiceAdapterBuilder.
func (b *IPStackServiceAdapterBuilder) SetClient(client *http.Client) *IPStackServiceAdapterBuilder {
	b.client = client

	return b
}

// SetHost sets the host for the IPStackServiceAdapterBuilder.
// It takes a string parameter `host` and returns a pointer to the IPStackServiceAdapterBuilder.
func (b *IPStackServiceAdapterBuilder) SetHost(host string) *IPStackServiceAdapterBuilder {
	b.host = host

	return b
}

// Build creates a new instance of IPStackServiceAdapter using the configured values from the builder.
// If the client is not set, a default http.Client will be used.
// If the host is not set, the default host "http://api.ipstack.com" will be used.
// It returns a pointer to the newly created IPStackServiceAdapter.
func (b *IPStackServiceAdapterBuilder) Build() *IPStackServiceAdapter {
	if b.client == nil {
		b.client = &http.Client{}
	}

	if b.host == "" {
		b.host = host
	}

	return &IPStackServiceAdapter{
		client: b.client,
		host:   b.host,
	}
}

// NewIPStackServiceAdapterBuilder creates a new instance of IPStackServiceAdapterBuilder.
// It returns a pointer to the newly created IPStackServiceAdapterBuilder.
func NewIPStackServiceAdapterBuilder() *IPStackServiceAdapterBuilder {
	return &IPStackServiceAdapterBuilder{}
}

// GetCoordinates retrieves the coordinates for a given IP address using the IPStack API.
// It takes the IP address and access key as parameters and returns the response body as a string.
// If an error occurs during the API request or response processing, it returns an empty string and the error.
func (a *IPStackServiceAdapter) GetCoordinates(ctx context.Context, ip, accessKey string) (string, error) { //nolint:varnamelen
	if ip == "" {
		return "", errIndalidIP
	}

	if accessKey == "" {
		return "", errInvalidAccessKey
	}

	url := fmt.Sprintf("http://%s/%s?access_key=%s", host, ip, accessKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return "", err
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
