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
	ApiKey *string
}

func buildUrl(appId string, path string) string {
	return fmt.Sprintf("%s/%s/%s", queueApiBaseUrl, appId, path)
}

func NewQueueHTTPClient(options *QueueHTTPClientOptions) (*QueueHTTPClient, error) {
	httpClient, err := httpclient.NewHTTPClient(&httpclient.HTTPClientOptions{
		ApiKey: options.ApiKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize queue client: %w", err)
	}
	return &QueueHTTPClient{
		httpClient: httpClient,
	}, nil
}
