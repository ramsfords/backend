package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (db UserDb) RemoveUser(ctx context.Context, email string) error {
	_, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "#users"},
		},
		ExpressionAttributeNames: map[string]string{
			"#name": email,
		},
		UpdateExpression: aws.String("REMOVE #name"),
	})
	return err
}
