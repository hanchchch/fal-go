package queueclient

import (
	"fmt"

	http "github.com/fal-ai/fal-go/pkg/httpclient"
)

// TODO generic
func (q *QueueHTTPClient) Result(appId string, requestId string) (*any, error) {
	res, err := http.NewJsonHttpRequest[any, any](q.httpClient).Get(
		buildUrl(appId, fmt.Sprintf("requests/%s", requestId)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue result: %w", err)
	}

	return res, nil
}
