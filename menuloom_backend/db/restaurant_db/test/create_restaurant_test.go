package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	restaurant "github.com/ramsfords/types_gen/v1"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestCreateRestaurant(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		restaurantData := restaurant.CreateRestaurantData{
			Id:                        "test",
			Name:                      "",
			RestaurantName:            "",
			RestaurantWebUrl:          "himalayen.menuloom.com",
			RestaurantS3DevUrl:        "",
			RestaurantS3StaticProdUrl: "",
			Address:                   &v1.RestaurantAddress{},
			PhoneNumber:               "",
			Email:                     "",
			OwnerId:                   "",
			OpenHours:                 map[string]*v1.Hours{},
			Created:                   "",
			Updated:                   "",
			Type:                      "restaurant",
			Pk:                        "pk#",
			Sk:                        "restaurant#",
		}
		data, err := attributevalue.Marshal(&restaurantData)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Item: map[string]types.AttributeValue{
				"pk":         &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + restaurantData.RestaurantWebUrl},
				"sk":         &types.AttributeValueMemberS{Value: "restaurant#" + restaurantData.RestaurantWebUrl},
				"restaurant": data,
			},
			ConditionExpression: aws.String("attribute_not_exists(pk)"),
		})
		fmt.Println(res, err)
	})
}
