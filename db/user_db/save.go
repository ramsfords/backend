package user_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (userdb UserDb) SaveUser(ctx context.Context, usr *v1.User, businessId string) error {
	usr.Type = "user"
	marshalledUser, err := attributevalue.Marshal(usr)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(userdb.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":      &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk":      &types.AttributeValueMemberS{Value: "user#" + usr.Email},
			"user_pk": &types.AttributeValueMemberS{Value: usr.Email},
			"users":   marshalledUser,
		},
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", "pk")),
	}
	_, err = userdb.Client.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
