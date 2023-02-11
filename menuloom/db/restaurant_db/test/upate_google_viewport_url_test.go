package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestUpdateGoogleViewportUrl(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		url := "https://www.google.com/maps/embed/v1/place?key=AIzaSyB0Z1Z0Z0Z0Z0Z0Z0Z0Z0Z0Z0Z0Z0Z0Z0&q=place_id:ChIJN1t_tDeuEmsRUsoyG83frY4"
		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#restaurant":           "restaurant",
				"#googleMapViewportUrl": "googleMapViewportUrl",
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":googleMapViewportUrl": &types.AttributeValueMemberS{Value: url},
			},
			UpdateExpression: aws.String("SET #restaurant.#googleMapViewportUrl = :googleMapViewportUrl"),
			ReturnValues:     types.ReturnValueAllNew,
		})
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(res, err)
	})
}
