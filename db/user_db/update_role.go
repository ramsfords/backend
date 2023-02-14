package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) UpdateUserRole(ctx context.Context, email string, data []*v1.Role) error {
	roles := []v1.Role{
		v1.Role_ADMIN,
		v1.Role_MANAGER,
	}
	newRoles, err := attributevalue.MarshalList(roles)
	if err != nil {
		return err
	}
	_, err = user.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "user"},
			"sk": &types.AttributeValueMemberS{Value: email},
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
