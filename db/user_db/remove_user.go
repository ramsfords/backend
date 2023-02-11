package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (user UserDb) RemoveUser(ctx context.Context, email string) error {
	_, err := user.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "user"},
			"sk": &types.AttributeValueMemberS{Value: email},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": email,
		},
		UpdateExpression: aws.String("REMOVE #name"),
	})
	return err
}
