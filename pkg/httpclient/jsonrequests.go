package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// TODO generic
func (h *HTTPClient) RequestJson(method string, url string, jsonBody interface{}, jsonRes interface{}) error {
	var reader io.Reader
	if jsonBody != nil {
		buf := bytes.NewBuffer(nil)
		err := json.NewEncoder(buf).Encode(jsonBody)
		if err != nil {
			return fmt.Errorf("failed to encode json: %w", err)
		}
		reader = bytes.NewReader(buf.Bytes())
	}

	res, err := h.Request(method, url, reader)
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

func (h *HTTPClient) GetJson(url string, jsonRes interface{}) error {
	return h.RequestJson("GET", url, nil, jsonRes)
}

func (h *HTTPClient) PostJson(url string, jsonBody interface{}, jsonRes interface{}) error {
	return h.RequestJson("POST", url, jsonBody, jsonRes)
}

func (h *HTTPClient) PutJson(url string, jsonBody interface{}, jsonRes interface{}) error {
	return h.RequestJson("PUT", url, jsonBody, jsonRes)
}
