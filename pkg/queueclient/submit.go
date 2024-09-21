package queueclient

import (
	"fmt"

	http "github.com/fal-ai/fal-go/pkg/httpclient"
)

// TODO generic
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

func (q *QueueHTTPClient) Submit(appId string, options *SubmitOptions) (*SubmitResponse, error) {
	reqOptions := &http.JsonHttpRequestOptions[interface{}]{
		Body: &options.Input,
	}

	if options.WebhookUrl != "" {
		reqOptions.Query = map[string]string{
			"fal_webhook": options.WebhookUrl,
		}
	}

	res, err := http.NewJsonHttpRequest[interface{}, SubmitResponse](q.httpClient).Post(
		buildUrl(appId, ""),
		reqOptions,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to submit queue: %w", err)
	}
	return res, nil
}
