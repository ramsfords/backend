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

func TestUpdateRestaurantOpenHours(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		data := map[string]*v1.Hours{
			"sunday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"monday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"tuesday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"wednesday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"thrusday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"friday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
			"saturday": {
				OpenHours: []*v1.OpenWindow{
					{
						From: "10:00",
						To:   "14:00",
					},
					{
						From: "17:00",
						To:   "21:00",
					},
				},
			},
		}
		mData, err := attributevalue.Marshal(&data)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#restaurant": "restaurant",
				"#openHours":  "openHours",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":openHours": mData,
			},
			UpdateExpression: aws.String("SET #restaurant.#openHours = :openHours"),
			ReturnValues:     types.ReturnValueAllNew,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
