package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestUpdatePhoneNumber(t *testing.T) {
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
		url := "+17135162836"
		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#restaurant":  "restaurant",
				"#phoneNumber": "phoneNumber",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":phoneNumber": &types.AttributeValueMemberS{Value: url},
			},
			UpdateExpression: aws.String("SET #restaurant.#phoneNumber = :phoneNumber"),
			ReturnValues:     types.ReturnValueAllNew,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
