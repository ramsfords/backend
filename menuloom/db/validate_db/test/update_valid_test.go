package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func TestUpdateValidate(t *testing.T) {
	_, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(conf.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + "himalayantasteusa.com"},
			"sk": &types.AttributeValueMemberS{Value: "valid"},
		},
		UpdateExpression: aws.String("SET #valid = :valid"),
		ExpressionAttributeNames: map[string]string{
			"#valid": "valid",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":valid": &types.AttributeValueMemberBOOL{Value: false},
		},
	})
	fmt.Print(err)

}
