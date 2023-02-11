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

// "ProjectionExpression": "#67ad0"
//
//	"ExpressionAttributeNames": {"#67ad0":"items","#67ad1":"pk"},
//	"ExpressionAttributeValues": {":67ad1": {"S":"pk#himalayen.menuloom.com"}}

func TestGetItem(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		restaurantPk := "himalayen.menuloom.com"
		itemName := "biryani"
		res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + restaurantPk},
				"sk": &types.AttributeValueMemberS{Value: "items#" + itemName},
			},
		})
		if err != nil {
			fmt.Println("no items")
		}
		data := v1.Item{}
		dbItem, ok := res.Item["items"]
		if ok {
			v, ok := dbItem.(*types.AttributeValueMemberM)
			if ok {
				value := v.Value
				err = attributevalue.UnmarshalMap(value, &data)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

	})
}
