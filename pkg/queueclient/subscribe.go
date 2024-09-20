package queueclient

import (
	"time"
)

const defaultPollingInterval = 500 * time.Millisecond

type SubscribeOptions struct {
	SubmitOptions
	PollingInterval *time.Duration
}

func (q *QueueHTTPClient) Subscribe(appId string, options *SubscribeOptions) (*any, error) {
	pollingInterval := defaultPollingInterval
	if options.PollingInterval != nil {
		pollingInterval = defaultPollingInterval
	}

	res, err := q.Submit(appId, &options.SubmitOptions)
	if err != nil {
		return nil, err
	}

	for {
		queueStatus, err := q.Status(appId, res.RequestId)
		if err != nil {
			return nil, err
		}
		time.Sleep(pollingInterval)

		if queueStatus.Status == QueueStatusCompleted {
			break
		}
	}

	result, err := q.Result(appId, res.RequestId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
