package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestUpdateItem(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		item := v1.Item{
			Name:  "tandoori chicken",
			Price: "200",
		}
		mData, err := attributevalue.Marshal(&item)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Item: map[string]types.AttributeValue{
				"pk":    &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayen.menuloom.com"},
				"sk":    &types.AttributeValueMemberS{Value: "item#" + item.Name},
				"items": mData,
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
