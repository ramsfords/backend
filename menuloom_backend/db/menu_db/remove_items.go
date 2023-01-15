package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (db MenuDb) RemoveItems(ctx context.Context, itemsNames []string, restaurantPk string) error {
	var errs error
	for _, item := range itemsNames {
		_, err := db.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
			TableName: aws.String(db.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
				"sk": &types.AttributeValueMemberS{Value: db.GetMenuloomSKPrefix() + item},
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
