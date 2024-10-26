package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	// AWS Configのロード
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// DynamoDBクライアントの作成
	svc := dynamodb.NewFromConfig(cfg)

	// DynamoDBのテーブル名
	tableName := "kst-skida-sample"

	putItem(svc, tableName)
	getItem(svc, tableName)
	getQuery(svc, tableName)
	getQueryGSI(svc, tableName)

}

func putItem(svc *dynamodb.Client, tableName string) {
	fmt.Println("------------ putItem ------------")

	// アイテムの追加（PutItem）
	item := map[string]types.AttributeValue{
		"PK":     &types.AttributeValueMemberS{Value: "PK2"},
		"SK":     &types.AttributeValueMemberS{Value: "SK3"},
		"Name":   &types.AttributeValueMemberS{Value: "Sample Name23"},
		"GSIPK1": &types.AttributeValueMemberS{Value: "GSIPK1"},
		"GSISK1": &types.AttributeValueMemberS{Value: "GSISK1"},
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err := svc.PutItem(context.TODO(), putItemInput)
	if err != nil {
		log.Fatalf("failed to put item, %v", err)
	}
	fmt.Println("Item successfully added to DynamoDB")
}

func getItem(svc *dynamodb.Client, tableName string) {
	fmt.Println("------------ getItem ------------")

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "PK2"},
			"SK": &types.AttributeValueMemberS{Value: "SK1"},
		},
	}

	result, err := svc.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Fatalf("failed to get item, %v", err)
	}

	if result.Item == nil {
		fmt.Println("Item not found")
		return
	}
	fmt.Println("Retrieved item:")
	fmt.Println("PK:", result.Item["PK"].(*types.AttributeValueMemberS).Value)
	fmt.Println("SK:", result.Item["SK"].(*types.AttributeValueMemberS).Value)
	fmt.Println("Name:", result.Item["Name"].(*types.AttributeValueMemberS).Value)
	// fmt.Println("Score:", result.Item["Score"].(*types.AttributeValueMemberN).Value)
}

func getQuery(svc *dynamodb.Client, tableName string) {
	fmt.Println("------------ getQuery ------------")

	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("PK = :pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "PK2"},
		},
	}

	result, err := svc.Query(context.TODO(), queryInput)
	if err != nil {
		log.Fatalf("failed to query items, %v", err)
	}

	if len(result.Items) == 0 {
		fmt.Println("No items found")
		return
	}

	fmt.Printf("Found %d items:\n", len(result.Items))
	for _, item := range result.Items {
		fmt.Println("PK:", item["PK"].(*types.AttributeValueMemberS).Value)
		fmt.Println("SK:", item["SK"].(*types.AttributeValueMemberS).Value)
		fmt.Println("Name:", item["Name"].(*types.AttributeValueMemberS).Value)
		fmt.Println("---")
	}

}

func getQueryGSI(svc *dynamodb.Client, tableName string) {
	fmt.Println("------------ getQueryGSI ------------")

	queryInput := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String("GSIPK1-GSISK1-index"),
		KeyConditionExpression: aws.String("GSIPK1 = :gsipk1"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":gsipk1": &types.AttributeValueMemberS{Value: "GSIPK1"},
		},
	}

	result, err := svc.Query(context.TODO(), queryInput)
	if err != nil {
		log.Fatalf("failed to query GSIPK1-GSISK1-index, %v", err)
	}

	if len(result.Items) == 0 {
		fmt.Println("No items found in GSIPK1-GSISK1-index")
		return
	}

	fmt.Printf("Found %d items in GSIPK1-GSISK1-index:\n", len(result.Items))
	for _, item := range result.Items {
		fmt.Println("GSIPK1:", item["GSIPK1"].(*types.AttributeValueMemberS).Value)
		fmt.Println("GSISK1:", item["GSISK1"].(*types.AttributeValueMemberS).Value)
		fmt.Println("PK:", item["PK"].(*types.AttributeValueMemberS).Value)
		fmt.Println("SK:", item["SK"].(*types.AttributeValueMemberS).Value)
		fmt.Println("Name:", item["Name"].(*types.AttributeValueMemberS).Value)
		fmt.Println("---")
	}
}

func Add(a, b int) int {
	return a + b
}
