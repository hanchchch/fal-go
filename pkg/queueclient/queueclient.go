package queueclient

import (
	"fmt"

	"github.com/fal-ai/fal-go/pkg/httpclient"
)

const queueApiBaseUrl = "https://queue.fal.run"

type QueueHTTPClient struct {
	httpClient *httpclient.HTTPClient
}

type QueueHTTPClientOptions struct {
	ApiKey string
}

type SubmitOptions struct {
	// The function input. It will be submitted either as query params
	Input      map[string]interface{} `json:"input,omitempty"`
	WebhookUrl string                 `json:"WebhookUrl,omitempty"`
}

type SubmitResponse struct {
	RequestId   string `json:"request_id"`
	ResponseUrl string `json:"response_url"`
	StatusUrl   string `json:"status_url"`
	CancelUrl   string `json:"cancel_url"`
}

func NewQueueHTTPClient(options *QueueHTTPClientOptions) *QueueHTTPClient {
	httpClient := httpclient.NewHTTPClient(&httpclient.HTTPClientOptions{
		BaseUrl: queueApiBaseUrl,
		ApiKey:  options.ApiKey,
	})
	return &QueueHTTPClient{
		httpClient: httpClient,
	}
}

func (q *QueueHTTPClient) Submit(appId string, options SubmitOptions) (*SubmitResponse, error) {
	res := &SubmitResponse{}
	err := q.httpClient.PostJson(appId, "submit", options, res)
	if err != nil {
		return nil, fmt.Errorf("failed to submit queue: %w", err)
	}
	return res, nil
}
