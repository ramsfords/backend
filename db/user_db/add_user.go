package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) AddUser(ctx context.Context, input *v1.User, businessId string) error {
	mData, err := attributevalue.Marshal(input)
	if err != nil {
		return err
	}
	_, err = user.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + input.Email},
		},
		ExpressionAttributeNames: map[string]string{
			"#newUser": input.Email,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user": mData,
		},

		UpdateExpression: aws.String("SET #newUser = :user"),
	})
	return err
}
