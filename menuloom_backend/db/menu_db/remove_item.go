package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (db MenuDb) RemoveItem(ctx context.Context, itemName string, restaurantPk string) error {
	_, err := db.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
			"sk": &types.AttributeValueMemberS{Value: db.GetMenuloomSKPrefix() + itemName},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
