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

type Items struct {
	Items []v1.Item `dynamodbav:"items"`
}

func TestGetItems(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		restaurantPk := "himalayen.menuloom.com"
		res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
			TableName:              aws.String(conf.GetMenuloomTableName()),
			KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :begins_with)"),
			ExpressionAttributeNames: map[string]string{
				"#pk":    "pk",
				"#sk":    "sk",
				"#items": "items",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":pk":          &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + restaurantPk},
				":begins_with": &types.AttributeValueMemberS{Value: "items#"},
			},
			ProjectionExpression:   aws.String("#items"),
			ReturnConsumedCapacity: types.ReturnConsumedCapacityTotal,
		})
		if err != nil {
			fmt.Println("no items")
		}
		itemsCollection := []v1.Item{}
		for _, items := range res.Items {
			datas, ok := items["items"]
			if ok {
				data := v1.Item{}
				v, ok := datas.(*types.AttributeValueMemberM)
				if ok {
					value := v.Value
					err = attributevalue.UnmarshalMap(value, &data)
					if err == nil {
						itemsCollection = append(itemsCollection, data)
					}
					continue
				}
			} else {
				continue
			}
		}
		fmt.Println(itemsCollection)

	})
}
