package dynamo

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func createTable(cli *dynamodb.Client, tablename string) {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: types.ScalarAttributeType("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: types.ScalarAttributeType("S"),
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       types.KeyType("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       types.KeyType("RANGE"),
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tablename),
	}

	_, err := cli.CreateTable(context.Background(), input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", tablename)
}
