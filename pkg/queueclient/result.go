package queueclient

import (
	"fmt"
)

type ResultWrapper[T any] struct {
	Result *T
}

func (q *QueueHTTPClient) Result(appId string, requestId string) (*interface{}, error) {
	res := &ResultWrapper[interface{}]{}
	err := q.httpClient.GetJson(buildUrl(appId, fmt.Sprintf("requests/%s", requestId)), res.Result)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue status: %w", err)
	}

	return res.Result, nil
}
