package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db MenuDb) AddItem(ctx context.Context, data *v1.Item, restaurantUrl string) error {
	mData, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	_, err = db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Item: map[string]types.AttributeValue{
			"pk":    &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantUrl},
			"sk":    &types.AttributeValueMemberS{Value: db.GetMenuloomSKPrefix() + data.Name},
			"items": mData,
		},
		ConditionExpression: aws.String("attribute_not_exists(sk)"),
	})
	return err
}
