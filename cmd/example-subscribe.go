package main

import (
	"fmt"
	"time"

	"github.com/fal-ai/fal-go/pkg/queueclient"
)

type ModelInput struct {
	Prompt string `json:"prompt"`
}

func main() {
	queue, err := queueclient.NewQueueHTTPClient(&queueclient.QueueHTTPClientOptions{})
	if err != nil {
		panic(err)
	}

	res, err := queue.Submit("fal-ai/fast-lightning-sdxl", &queueclient.SubmitOptions{
		Input: &ModelInput{
			Prompt: "photo of a girl smiling during a sunset, with lightnings in the background",
		},
	})
	if err != nil {
		panic(err)
	}

	for {
		queueStatus, err := queue.Status("fal-ai/fast-lightning-sdxl", res.RequestId)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Status: %v\n", queueStatus.Status)
		time.Sleep(500 * time.Millisecond)

		if queueStatus.Status == queueclient.QueueStatusCompleted {
			break
		}
	}

	result, err := queue.Result("fal-ai/fast-lightning-sdxl", res.RequestId)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %+v\n", *result)
}
