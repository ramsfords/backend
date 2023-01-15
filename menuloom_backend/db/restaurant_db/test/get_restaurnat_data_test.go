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

// "KeyConditionExpression": "#69240 = :69240",
// "ExpressionAttributeNames": {"#69240":"pk"},
// "ExpressionAttributeValues": {":69240": {"S":"pk#127.0.0.1:3000"}}

func TestGetRestaurantData(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#himalayantasteusa.com"},
				"sk": &types.AttributeValueMemberS{Value: "restaurant#himalayantasteusa.com"},
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		data := &v1.CreateRestaurantData{}
		dbItem, ok := res.Item["restaurant"]
		if ok {
			v, ok := dbItem.(*types.AttributeValueMemberM)
			if ok {
				value := v.Value
				err = attributevalue.UnmarshalMap(value, data)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	})
}
