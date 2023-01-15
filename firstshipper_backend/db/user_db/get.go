package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) GetUser(ctx context.Context, userId string) (*v1.User, error) {
	urEmailInput := &dynamodb.QueryInput{
		TableName:              aws.String(user.Config.GetFirstShipperTableName()),
		IndexName:              aws.String("user_index"),
		KeyConditionExpression: aws.String("#user_pk = :pkey and #user_sk = :skey"),

		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: "user"},
			":skey": &types.AttributeValueMemberS{Value: userId},
		},
		ExpressionAttributeNames: map[string]string{
			"#user_pk": "user_pk",
			"#user_sk": "user_sk",
		},
		ScanIndexForward: aws.Bool(true),
	}
	res, err := user.Client.Query(ctx, urEmailInput)
	if err != nil {
		return &v1.User{}, err
	}
	if len(res.Items) == 0 {
		return &v1.User{}, nil
	}
	return &v1.User{}, nil
}
