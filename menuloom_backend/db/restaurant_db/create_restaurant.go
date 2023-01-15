package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db RestaurantDb) CreateRestaurant(ctx context.Context, data *v1.CreateRestaurantData) error {
	mData, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	_, err = db.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Item: map[string]types.AttributeValue{
			"pk":         &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + data.RestaurantWebUrl},
			"sk":         &types.AttributeValueMemberS{Value: "restaurant#" + data.RestaurantWebUrl},
			"restaurant": mData},
		ConditionExpression: aws.String("attribute_not_exists(sk)"),
	})

	if err != nil {
		return err
	}
	return nil
}
