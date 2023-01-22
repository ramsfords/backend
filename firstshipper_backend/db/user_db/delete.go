package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (userDb UserDb) DeleteUser(ctx context.Context, userId string, businessId string) error {
	_, err := userDb.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(userDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + userId},
		},
	})
	if err != nil {
		return err
	}
	return nil
}
