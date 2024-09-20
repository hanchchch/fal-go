package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// TODO generic
type JsonHttpRequest[ReqBody any, ResBody any] struct {
	httpClient *HttpClient
}

func NewJsonHttpRequest[ReqBody any, ResBody any](httpClient *HttpClient) *JsonHttpRequest[ReqBody, ResBody] {
	return &JsonHttpRequest[ReqBody, ResBody]{
		httpClient: httpClient,
	}
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Request(method string, url string, jsonBody *ReqBody) (*ResBody, error) {
	var reader io.Reader
	if jsonBody != nil {
		buf := bytes.NewBuffer(nil)
		err := json.NewEncoder(buf).Encode(jsonBody)
		if err != nil {
			return nil, fmt.Errorf("failed to encode json: %w", err)
		}
		reader = bytes.NewReader(buf.Bytes())
	}

	res, err := j.httpClient.Request(method, url, reader, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()

	var resBody ResBody
	err = json.NewDecoder(res.Body).Decode(&resBody)

	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &resBody, nil
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Get(url string) (*ResBody, error) {
	return j.Request("GET", url, nil)
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Post(url string, jsonBody *ReqBody) (*ResBody, error) {
	return j.Request("POST", url, jsonBody)
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Put(url string, jsonBody *ReqBody) (*ResBody, error) {
	return j.Request("PUT", url, jsonBody)
}
