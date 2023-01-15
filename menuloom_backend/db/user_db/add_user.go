package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (db UserDb) AddUser(ctx context.Context, input *v1.User) error {
	mData, err := attributevalue.Marshal(input)
	if err != nil {
		return err
	}
	_, err = db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "users#"},
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
