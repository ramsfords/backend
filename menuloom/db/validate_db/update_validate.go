package validatedb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (db ValidateDb) UpdateValidate(ctx context.Context, id string, isValid bool) error {
	_, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix() + id},
			"sk": &types.AttributeValueMemberS{Value: "valid"},
		},
		UpdateExpression: aws.String("SET #valid = :valid"),
		ExpressionAttributeNames: map[string]string{
			"#valid": "valid",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":valid": &types.AttributeValueMemberBOOL{Value: isValid},
		},
	})
	return err

}
