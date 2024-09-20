package httpclient

import (
	"fmt"
	"io"
	"net/http"
)

func (h *HTTPClient) Request(method string, url string, body io.Reader) (*http.Response, error) {
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

func (h *HTTPClient) Get(url string) (*http.Response, error) {
	return h.Request("GET", url, nil)
}

func (h *HTTPClient) Post(url string, body io.Reader) (*http.Response, error) {
	return h.Request("POST", url, body)
}

func (h *HTTPClient) Put(url string, body io.Reader) (*http.Response, error) {
	return h.Request("PUT", url, body)
}
