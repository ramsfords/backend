package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (userdb UserDb) SaveRefreshToken(ctx context.Context, businessId string, userId string, token string) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(userdb.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":             &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk":             &types.AttributeValueMemberS{Value: "sk#" + userId},
			"refreshTokenPk": &types.AttributeValueMemberS{Value: userId},
			"refreshToken":   &types.AttributeValueMemberS{Value: token},
		},
	}
	_, err := userdb.Client.PutItem(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
