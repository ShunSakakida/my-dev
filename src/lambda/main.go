package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response struct {
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (*Response, error) {
	// リクエストボディをMyEvent構造体にパース
	var event MyEvent
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request body: %v", err)
	}

	// メッセージ生成
	message := fmt.Sprintf("Hello %s!", event.Name)
	body, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		return nil, err
	}

	// レスポンス生成
	return &Response{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
