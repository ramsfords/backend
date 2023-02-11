package validatedb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type ValidateDb struct {
	dynamo.DB
	*configs.Config
}

func (db ValidateDb) CreateValidate(ctx context.Context, id string, isValid bool) error {

	_, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Item: map[string]types.AttributeValue{
			"pk":    &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + id},
			"sk":    &types.AttributeValueMemberS{Value: "valid"},
			"valid": &types.AttributeValueMemberBOOL{Value: isValid},
		},
	})
	return err

}
