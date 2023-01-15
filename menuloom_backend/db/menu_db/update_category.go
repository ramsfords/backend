package menu_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db MenuDb) UpdateCategory(ctx context.Context, data *v1.Category, restaurantPk string) error {
	marshallCategory, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	_, err = db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Item: map[string]types.AttributeValue{
			"pk":         &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + restaurantPk},
			"sk":         &types.AttributeValueMemberS{Value: "category#" + "category#" + data.Name},
			"categories": marshallCategory,
		},
	})
	if err != nil {
		return err
	}

	return nil

}
