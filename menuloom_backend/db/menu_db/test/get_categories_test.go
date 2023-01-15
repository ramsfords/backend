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

func TestGetCategories(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "category#"},
			},
		})
		if err != nil {
			fmt.Println("no items")
		}
		categories := []*v1.Category{}
		for k, v := range res.Item {
			cat, ok := v.(*types.AttributeValueMemberM)
			if ok {
				fmt.Println(k, cat)
				data := &v1.Category{}
				err = attributevalue.Unmarshal(cat, data)
				if err != nil {
					fmt.Println(err)
				}
				categories = append(categories, data)
			}
		}
		fmt.Println(res, err, res)
	})
}
