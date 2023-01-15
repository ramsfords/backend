package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db RestaurantDb) UpdateAddress(ctx context.Context, data *v1.RestaurantAddress) error {
	dataNew, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	_, err = db.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
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
		return err
	}
	return nil
}
