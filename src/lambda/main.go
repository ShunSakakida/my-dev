package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type AppSyncEvent struct {
	Arguments map[string]interface{} `json:"arguments"`
}

type Response struct {
	Message string `json:"message"`
	Length  int    `json:"length"`
}

func HandleRequest(ctx context.Context, event AppSyncEvent) (Response, error) {
	name, ok := event.Arguments["name"].(string)
	if !ok {
		return Response{}, fmt.Errorf("invalid argument")
	}

	message := fmt.Sprintf("Hello, %s!", name)
	length := len(message)

	return Response{
		Message: message,
		Length:  length,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
