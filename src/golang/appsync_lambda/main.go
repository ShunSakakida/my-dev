package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// AppSyncからのリクエストイベント構造体
type AppSyncEvent struct {
	Arguments GreetArguments `json:"arguments"`
}

// greetクエリの引数構造体
type GreetArguments struct {
	Name string `json:"name"`
}

// レスポンス構造体（スキーマのGreetResponseに対応）
type GreetResponse struct {
	Message string `json:"message"`
	Length  int    `json:"length"`
}

// ハンドラ関数
func HandleRequest(ctx context.Context, event AppSyncEvent) (GreetResponse, error) {
	// 引数からnameを取得
	name := event.Arguments.Name

	// メッセージを生成
	message := fmt.Sprintf("Hello, %s!", name)
	length := len(message)

	// レスポンスを生成
	return GreetResponse{
		Message: message,
		Length:  length,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
