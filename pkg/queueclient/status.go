package queueclient

import (
	"fmt"

	http "github.com/fal-ai/fal-go/pkg/httpclient"
)

type RequestLogLevel string

const (
	RequestLogLevelStderr RequestLogLevel = "STDERR"
	RequestLogLevelStdout RequestLogLevel = "STDOUT"
	RequestLogLevelError  RequestLogLevel = "ERROR"
	RequestLogLevelInfo   RequestLogLevel = "INFO"
	RequestLogLevelWarn   RequestLogLevel = "WARN"
	RequestLogLevelDebug  RequestLogLevel = "DEBUG"
)

type RequestLogSource string

const (
	RequestLogSourceUser RequestLogSource = "USER"
)

type RequestLog struct {
	Message   string           `json:"message"`
	Level     RequestLogLevel  `json:"level"`
	Source    RequestLogSource `json:"source"`
	Timestamp string           `json:"timestamp"`
}

type Metrics struct {
	InferenceTime *float32 `json:"inference_time"`
}

type QueueStatus string

const (
	QueueStatusInProgress QueueStatus = "IN_PROGRESS"
	QueueStatusCompleted  QueueStatus = "COMPLETED"
	QueueStatusInQueue    QueueStatus = "IN_QUEUE"
)

// QueueStatusResponse is the response of the status endpoint
type QueueStatusResponse struct {
	Status      QueueStatus `json:"status"`
	ResponseUrl string      `json:"response_url"`

	Logs          *[]RequestLog `json:"logs,omitempty"`
	Metrics       *Metrics      `json:"metrics,omitempty"`
	QueuePosition *int          `json:"queue_position,omitempty"`
}

func (q *QueueHTTPClient) Status(appId string, requestId string) (*QueueStatusResponse, error) {
	res, err := http.NewJsonHttpRequest[any, QueueStatusResponse](q.httpClient).Get(
		buildUrl(appId, fmt.Sprintf("requests/%s/status", requestId)),
		&http.JsonHttpRequestOptions[any]{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue status: %w", err)
	}

	return res, nil
}
