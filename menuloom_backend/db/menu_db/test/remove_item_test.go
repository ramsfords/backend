package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestRemoveItem(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {

		res, err := db.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayen.menuloom.com"},
				"sk": &types.AttributeValueMemberS{Value: "item#" + "chicken tikka"},
			},
		})
		fmt.Println(res, err)
	})
}
