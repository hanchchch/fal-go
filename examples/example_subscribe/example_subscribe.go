package main

import (
	"fmt"

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

	result, err := queue.Subscribe("fal-ai/fast-lightning-sdxl", &queueclient.SubscribeOptions{
		SubmitOptions: queueclient.SubmitOptions{
			Input: ModelInput{
				Prompt: "photo of a girl smiling during a sunset, with lightnings in the background",
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %+v\n", *result)
}
