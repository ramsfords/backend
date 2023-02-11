package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db UserDb) UpdateUserRole(ctx context.Context, email string, data []*v1.Role) error {
	newRoles, err := attributevalue.MarshalList(data)
	if err != nil {
		return err
	}
	_, err = db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "#users"},
		},
		ExpressionAttributeNames: map[string]string{
			"#Role":     "Role",
			"#username": email,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":role": &types.AttributeValueMemberL{Value: newRoles},
		},
		UpdateExpression: aws.String("SET #username.#Role = :role"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}
	return err
}
