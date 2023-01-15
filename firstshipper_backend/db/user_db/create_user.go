package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) CreateUser(ctx context.Context, input *v1.User, businessId string) error {
	data, err := attributevalue.Marshal(input)
	if err != nil {
		return err
	}
	mData, err := attributevalue.Marshal(&input)
	if err != nil {
		return err
	}
	_, err = user.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":        &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk":        &types.AttributeValueMemberS{Value: "user#" + input.Email},
			input.Email: data,
			"user": &types.AttributeValueMemberM{
				Value: map[string]types.AttributeValue{
					input.Email: mData,
				},
			},
		},
	})
	return err

}
