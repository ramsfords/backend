package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestRemoveUser(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {

		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "users#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#name": "reakeshneupane2045@gmail.com",
			},
			UpdateExpression: aws.String("REMOVE #name"),
		})
		fmt.Println(res, err)
	})
}
