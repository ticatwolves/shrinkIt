package db

import (
	"context"
	"fmt"
	"log"
	"os"
	schema "shrinkIt/internal/schemas"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func getdynamoDBClient() (*dynamodb.Client, error) {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	dynamoDBEndpointURL := os.Getenv("DYNAMODB_ENDPOINT")
	if dynamoDBEndpointURL != "" {
		return dynamodb.NewFromConfig(config, func(o *dynamodb.Options) {
			o.BaseEndpoint = &dynamoDBEndpointURL
		}), nil
	} else {
		return dynamodb.NewFromConfig(config), nil
	}
}

func Insert(hash_url *string, actual_url *string) {
	db_client, _ := getdynamoDBClient()
	item := map[string]types.AttributeValue{
		"UrlHash":   &types.AttributeValueMemberS{Value: *hash_url},
		"ActualUrl": &types.AttributeValueMemberS{Value: *actual_url},
	}
	tableName := os.Getenv("DYNAMODB_TABLE")
	putItemInput := &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      item,
	}
	_, err := db_client.PutItem(context.TODO(), putItemInput)
	if err != nil {
		log.Fatalf("failed to put item, %v", err)
	}
	log.Println("Successfully added item to DynamoDB.")
}

func GetByHash(hash_url *string) schema.ShrinkIt {
	db_client, _ := getdynamoDBClient()
	var tableName string = os.Getenv("DYNAMODB_TABLE")
	fmt.Println(*hash_url)
	input := &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"UrlHash": &types.AttributeValueMemberS{Value: *hash_url},
		},
		ConsistentRead: aws.Bool(false),
	}
	item, _ := db_client.GetItem(context.TODO(), input)
	var hrul string
	var aurl string
	if urlHash, ok := item.Item["UrlHash"]; ok {
		if s, ok := urlHash.(*types.AttributeValueMemberS); ok {
			hrul = s.Value
		}
	}
	if actualUrl, ok := item.Item["ActualUrl"]; ok {
		if s, ok := actualUrl.(*types.AttributeValueMemberS); ok {
			aurl = s.Value
		}
	}
	return schema.ShrinkIt{UrlHash: hrul, ActualUrl: aurl}
}

// from boto3 import client, resource
// endpoint_url = "http://localhost:8000"
// client = resource(service_name="dynamodb", endpoint_url=endpoint_url)
// table = client.Table("ShrinkItUrlMapping")
// client.create_table(TableName="ShrinkItUrlMapping", KeySchema=[{"AttributeName": "UrlHash", "KeyType": "HASH"}], AttributeDefinitions=[{"AttributeName": "UrlHash", "AttributeType": "S"}], Provisioned\
// Throughput={"ReadCapacityUnits": 1, "WriteCapacityUnits": 1})
