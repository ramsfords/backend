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
)

func TestUpdateAddress(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		// restaurantData := restaurant.CreateRestaurantData{
		// 	Id:                  "test",
		// 	Name:                "",
		// 	RestaurantName:      "",
		// 	RestaurantWebUrl:    "",
		// 	RestaurantS3DevUrl:  "",
		// 	RestaurantS3ProdUrl: "",
		// 	Address:             &v1.Address{},
		// 	PhoneNumber:         "",
		// 	Email:               "",
		// 	OwnerId:             "",
		// 	OpenHours:           map[string]*v1.Hours{},
		// 	Created:             "",
		// 	Updated:             "",
		// 	Type:                "restaurant",
		// 	Pk:                  "pk#",
		// 	Sk:                  conf.RestaurantSk,
		// }
		// data, err := attributevalue.Marshal(&restaurantData)
		// if err != nil {
		// 	t.Fatal(err)
		// }
		data := restaurant.RestaurantAddress{
			Street1: "street1",
			City:    "city",
			State:   "state",
			ZipCode: "zip",
			Country: "usa",
		}
		dataNew, err := attributevalue.Marshal(&data)
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
				"#address":    "address",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":address": dataNew,
			},
			UpdateExpression: aws.String("SET #restaurant.#address = :address"),
			ReturnValues:     types.ReturnValueAllNew,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
