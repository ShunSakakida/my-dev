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

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// リクエストボディをMyEvent構造体にパース
	var event MyEvent
	if err := json.Unmarshal([]byte(request.Body), &event); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       fmt.Sprintf("Invalid request: %v", err),
		}, nil
	}

	// メッセージ生成
	message := fmt.Sprintf("Hello %s!", event.Name)
	body, err := json.Marshal(map[string]string{"message": message})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to generate response",
		}, nil
	}

	// 正常なレスポンス
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
