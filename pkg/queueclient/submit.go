package queueclient

import (
	"fmt"
)

type SubmitOptions struct {
	// The function input. It will be submitted either as query params
	Input      interface{} `json:"input,omitempty"`
	WebhookUrl string      `json:"WebhookUrl,omitempty"`
}

type SubmitResponse struct {
	RequestId   string `json:"request_id"`
	ResponseUrl string `json:"response_url"`
	StatusUrl   string `json:"status_url"`
	CancelUrl   string `json:"cancel_url"`
}

func (q *QueueHTTPClient) Submit(appId string, options SubmitOptions) (*SubmitResponse, error) {
	res := &SubmitResponse{}
	err := q.httpClient.PostJson(buildUrl(appId, "submit"), options, res)
	if err != nil {
		return nil, fmt.Errorf("failed to submit queue: %w", err)
	}
	return res, nil
}
