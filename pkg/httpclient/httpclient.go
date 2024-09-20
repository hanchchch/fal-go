package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient struct {
	client  *http.Client
	baseUrl string
	apiKey  string
}

type HTTPClientOptions struct {
	BaseUrl string
	ApiKey  string
}

func NewHTTPClient(options *HTTPClientOptions) *HTTPClient {
	return &HTTPClient{
		client:  &http.Client{},
		baseUrl: options.BaseUrl,
		apiKey:  options.ApiKey,
	}
}

func (h *HTTPClient) buildUrl(appId string, path string) string {
	return fmt.Sprintf("%s/%s/%s", h.baseUrl, appId, path)
}

func (h *HTTPClient) Request(method string, appId string, path string, body io.Reader) (*http.Response, error) {
	url := h.buildUrl(appId, path)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Key %s", h.apiKey))

	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	return res, nil
}

func (h *HTTPClient) Get(appId string, path string) (*http.Response, error) {
	return h.Request("GET", appId, path, nil)
}

func (h *HTTPClient) Post(appId string, path string, body io.Reader) (*http.Response, error) {
	return h.Request("POST", appId, path, body)
}

// TODO generic
func (h *HTTPClient) GetJson(appId string, path string) (interface{}, error) {
	res, err := h.Request("GET", appId, path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	var jsonRes interface{}
	err = json.NewDecoder(res.Body).Decode(&jsonRes)

	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return jsonRes, nil
}

// TODO generic
func (h *HTTPClient) PostJson(appId string, path string, jsonBody interface{}, jsonRes interface{}) error {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(jsonBody)
	if err != nil {
		return fmt.Errorf("failed to encode json: %w", err)
	}

	body := bytes.NewReader(buf.Bytes())
	res, err := h.Request("POST", appId, path, body)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&jsonRes)

	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}
