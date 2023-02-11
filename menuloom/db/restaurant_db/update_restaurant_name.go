package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db RestaurantDb) UpdateRestaurantName(ctx context.Context, name string) error {
	_, err := db.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
		},
		ExpressionAttributeNames: map[string]string{
			"#restaurant":     "restaurant",
			"#restaurantName": "restaurantName",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":restaurantName": &types.AttributeValueMemberS{Value: name},
		},
		UpdateExpression: aws.String("SET #restaurant.#restaurantName = :restaurantName"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}
	// update another name field with same name for the restaurant
	go db.UpdateName(ctx, name)
	return nil
}
func (db RestaurantDb) UpdateName(ctx context.Context, name string) error {
	_, err := db.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
		},
		ExpressionAttributeNames: map[string]string{
			"#restaurant": "restaurant",
			"#name":       "name",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":name": &types.AttributeValueMemberS{Value: name},
		},
		UpdateExpression: aws.String("SET #restaurant.#name = :name"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}

	return nil
}
