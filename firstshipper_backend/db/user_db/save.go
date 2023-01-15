package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) SaveUser(ctx context.Context, usr v1.User, businessId string) error {
	usr.Type = "user"
	marshalledUser, err := attributevalue.MarshalMap(usr)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":      &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk":      &types.AttributeValueMemberS{Value: "sk#" + usr.Email},
			"user_pk": &types.AttributeValueMemberS{Value: "user"},
			"user_sk": &types.AttributeValueMemberS{Value: usr.Email},
			"user":    &types.AttributeValueMemberM{Value: marshalledUser},
		},
	}
	_, err = user.Client.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
