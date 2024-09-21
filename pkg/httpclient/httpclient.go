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

type RequestOptions struct {
	Headers map[string]string
	Query   map[string]string
	Body    *io.Reader
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

func appendQuery(url string, query map[string]string) string {
	if query != nil {
		url = fmt.Sprintf("%s?", url)
		for key, value := range query {
			url = fmt.Sprintf("%s%s=%s&", url, key, value)
		}
		url = url[:len(url)-1]
	}
	return url
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

func (h *HttpClient) Request(method string, url string, options *RequestOptions) (*http.Response, error) {
	processedUrl := url
	if options.Query != nil {
		processedUrl = appendQuery(url, options.Query)
	}

	req, err := http.NewRequest(method, processedUrl, *options.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Key %s", h.apiKey))
	for key, value := range options.Headers {
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
