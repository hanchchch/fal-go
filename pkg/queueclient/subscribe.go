package queueclient

import (
	"fmt"
	"time"
)

const defaultPollingInterval = 500 * time.Millisecond
const defaultTimeout = 300 * time.Second

type SubscribeOptions struct {
	SubmitOptions
	Timeout         time.Duration
	PollingInterval time.Duration
}

func (q *QueueHTTPClient) Subscribe(appId string, options *SubscribeOptions) (*any, error) {
	if options.PollingInterval == 0 {
		options.PollingInterval = defaultPollingInterval
	}

	if options.Timeout == 0 {
		options.Timeout = defaultTimeout
	}

	res, err := q.Submit(appId, &options.SubmitOptions)
	if err != nil {
		return nil, err
	}

	for alive := true; alive; {
		timer := time.NewTimer(options.Timeout)
		tick := time.NewTicker(options.PollingInterval)

		select {
		case <-timer.C:
			alive = false
			return nil, fmt.Errorf("timeout exceeded (duration: %v)", options.Timeout.String())
		case <-tick.C:
			queueStatus, err := q.Status(appId, res.RequestId)
			if err != nil {
				return nil, err
			} else if queueStatus.Status == QueueStatusCompleted {
				timer.Stop()
				alive = false
			}
		}
	}

	result, err := q.Result(appId, res.RequestId)
	if err != nil {
		return nil, err
	}

	return result, nil
}
