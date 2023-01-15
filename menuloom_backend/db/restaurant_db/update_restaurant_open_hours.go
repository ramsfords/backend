package restaurant_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db RestaurantDb) UpdateRestaurantOpenHours(ctx context.Context, hours map[string]*v1.Hours) error {
	data, err := attributevalue.Marshal(&hours)
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
			"#openHours":  "openHours",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":openHours": data,
		},
		UpdateExpression: aws.String("SET #restaurant.#openHours = :openHours"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}
	return err
}
