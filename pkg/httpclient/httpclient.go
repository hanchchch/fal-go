package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type HttpClient struct {
	client *http.Client
	apiKey string
}

type HttpClientOptions struct {
	ApiKey *string
}

func getApiKey(options *HttpClientOptions) string {
	apiKey := ""

	if os.Getenv("FAL_KEY") != "" {
		apiKey = os.Getenv("FAL_KEY")
	}

	if options.ApiKey != nil {
		apiKey = *options.ApiKey
	}

	return apiKey
}

func NewHTTPClient(options *HttpClientOptions) (*HttpClient, error) {
	apiKey := getApiKey(options)
	if apiKey == "" {
		return nil, fmt.Errorf("failed to initialize http client: api key is required")
	}
	return &HttpClient{
		client: &http.Client{},
		apiKey: apiKey,
	}, nil
}

func (h *HttpClient) Request(method string, url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Key %s", h.apiKey))
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		errMsg := ""
		if res.Body != nil {
			buf := bytes.NewBuffer(nil)
			buf.ReadFrom(res.Body)
			errMsg = buf.String()
		}
		return nil, fmt.Errorf("failed to send request: %s %s", res.Status, errMsg)
	}

	return res, nil
}
