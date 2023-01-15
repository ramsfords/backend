package validatedb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (db ValidateDb) GetValidate(ctx context.Context, id string) (bool, error) {

	res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + id},
			"sk": &types.AttributeValueMemberS{Value: "valid"},
		},
	})
	if err != nil {
		return false, err
	}
	valid, ok := res.Item["valid"].(*types.AttributeValueMemberBOOL)
	if !ok {
		return false, nil
	}
	return valid.Value, nil

}
