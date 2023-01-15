package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func TestPutValidate(t *testing.T) {
	_, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(conf.GetMenuloomTableName()),
		Item: map[string]types.AttributeValue{
			"pk":    &types.AttributeValueMemberS{Value: "menu#" + "himalayantasteusa.com"},
			"sk":    &types.AttributeValueMemberS{Value: "valid"},
			"valid": &types.AttributeValueMemberBOOL{Value: false},
		},
	})
	fmt.Println(err)

}
