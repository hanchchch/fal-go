package httpclient

import (
	"fmt"
	"net/http"
	"os"
)

type HTTPClient struct {
	client *http.Client
	apiKey string
}

type HTTPClientOptions struct {
	ApiKey *string
}

func getApiKey(options *HTTPClientOptions) string {
	apiKey := ""

	if os.Getenv("FAL_KEY") != "" {
		apiKey = os.Getenv("FAL_KEY")
	}

	if options.ApiKey != nil {
		apiKey = *options.ApiKey
	}

	return apiKey
}

func NewHTTPClient(options *HTTPClientOptions) (*HTTPClient, error) {
	apiKey := getApiKey(options)
	if apiKey == "" {
		return nil, fmt.Errorf("failed to initialize http client: api key is required")
	}
	return &HTTPClient{
		client: &http.Client{},
		apiKey: apiKey,
	}, nil
}
