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

func TestUpdateRestaurantWebUrl(t *testing.T) {
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
		roles := []v1.Role{
			v1.Role_ADMIN,
		}
		newRoles, err := attributevalue.MarshalList(roles)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "users#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#Role":     "Role",
				"#username": "kandelsuren@gmail.com",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":role": &types.AttributeValueMemberL{Value: newRoles},
			},
			UpdateExpression: aws.String("SET #username.#Role = :role"),
			ReturnValues:     types.ReturnValueAllNew,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
