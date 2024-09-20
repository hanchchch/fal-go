package queueclient

import "fmt"

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

type InProgressQueueStatus struct {
	Status      string       `json:"status"`
	ResponseUrl *string      `json:"response_url,omitempty"`
	Logs        []RequestLog `json:"logs,omitempty"`
}

type CompletedQueueStatus struct {
	Status      string        `json:"-"`
	ResponseUrl *string       `json:"-"`
	Logs        *[]RequestLog `json:"-"`
	Metrics     *Metrics      `json:"metrics,omitempty"`
}

type EnqueuedQueueStatus struct {
	Status        string  `json:"-"`
	ResponseUrl   *string `json:"-"`
	QueuePosition *int    `json:"queue_position,omitempty"`
}

// QueueStatusResponse is the response of the status endpoint
type QueueStatusResponse struct {
	InProgressQueueStatus
	EnqueuedQueueStatus
	CompletedQueueStatus
}

func (q *QueueHTTPClient) Status(appId string, requestId string) (*QueueStatusResponse, error) {
	res := &QueueStatusResponse{}
	err := q.httpClient.GetJson(appId, fmt.Sprintf("requests/%s/status", requestId), res)
	if err != nil {
		return nil, fmt.Errorf("failed to get queue status: %w", err)
	}

	return res, nil
}
