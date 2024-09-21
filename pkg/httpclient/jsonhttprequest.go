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

type JsonHttpRequestOptions[ReqBody any] struct {
	Query map[string]string
	Body  *ReqBody
}

func NewJsonHttpRequest[ReqBody any, ResBody any](httpClient *HttpClient) *JsonHttpRequest[ReqBody, ResBody] {
	return &JsonHttpRequest[ReqBody, ResBody]{
		httpClient: httpClient,
	}
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Request(method string, url string, options *JsonHttpRequestOptions[ReqBody]) (*ResBody, error) {
	var reader io.Reader
	reqOptions := &RequestOptions{
		Body:    &reader,
		Headers: map[string]string{},
		Query:   options.Query,
	}

	if options.Body != nil {
		buf := bytes.NewBuffer(nil)
		err := json.NewEncoder(buf).Encode(options.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to encode json: %w", err)
		}
		reader = bytes.NewReader(buf.Bytes())
		reqOptions.Headers["Content-Type"] = "application/json"
	}

	res, err := j.httpClient.Request(method, url, reqOptions)
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

func (j *JsonHttpRequest[ReqBody, ResBody]) Get(url string, options *JsonHttpRequestOptions[ReqBody]) (*ResBody, error) {
	return j.Request("GET", url, options)
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Post(url string, options *JsonHttpRequestOptions[ReqBody]) (*ResBody, error) {
	return j.Request("POST", url, options)
}

func (j *JsonHttpRequest[ReqBody, ResBody]) Put(url string, options *JsonHttpRequestOptions[ReqBody]) (*ResBody, error) {
	return j.Request("PUT", url, options)
}
