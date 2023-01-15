package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db RestaurantDb) UpdateS3ProdUrl(ctx context.Context, url string) error {
	_, err := db.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "restaurant#"},
		},
		ExpressionAttributeNames: map[string]string{
			"#restaurant":          "restaurant",
			"#restaurantS3ProdUrl": "restaurantS3ProdUrl",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":restaurantS3ProdUrl": &types.AttributeValueMemberS{Value: url},
		},
		UpdateExpression: aws.String("SET #restaurant.#restaurantS3ProdUrl = :restaurantS3ProdUrl"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}
	return err
}
