package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db MenuDb) UpdateItems(ctx context.Context, data []*v1.Item, restaurantPk string) error {
	var errs error
	for _, item := range data {
		mData, err := attributevalue.Marshal(data)
		if err != nil {
			errs = err
		}
		_, err = db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(db.GetMenuloomTableName()),
			Item: map[string]types.AttributeValue{
				"pk":    &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
				"sk":    &types.AttributeValueMemberS{Value: db.GetMenuloomSKPrefix() + item.Name},
				"items": mData,
			},
		})
		if err != nil {
			errs = err
		}
	}
	if errs != nil {
		return errs
	}

	return nil
}
