package queueclient

import (
	"fmt"

	http "github.com/fal-ai/fal-go/pkg/httpclient"
)

func (q *QueueHTTPClient) Cancel(appId string, requestId string) error {
	_, err := http.NewJsonHttpRequest[any, any](q.httpClient).Put(
		buildUrl(appId, fmt.Sprintf("requests/%s/cancel", requestId)),
		&http.JsonHttpRequestOptions[any]{},
	)
	if err != nil {
		return fmt.Errorf("failed to get queue status: %w", err)
	}
	return nil
}
