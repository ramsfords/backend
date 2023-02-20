package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (userdb UserDb) GetRefreshToken(ctx context.Context, userId string) (string, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(userdb.GetFirstShipperTableName()),
		IndexName:              aws.String("refreshToken_index"),
		KeyConditionExpression: aws.String("#refreshTokenPK = :refreshTokenPK"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":refreshTokenPK": &types.AttributeValueMemberS{Value: userId},
		},
		ExpressionAttributeNames: map[string]string{
			"#refreshTokenPK": "refreshTokenPk",
		},
		ProjectionExpression: aws.String("#refreshToken"),
	}
	refreshToken, err := userdb.Client.Query(ctx, input)
	// getToken from response
	token, ok := refreshToken.Items[0]["refreshToken"]
	if !ok {
		return "", err
	}
	// unmarshal token
	tokenString := token.(*types.AttributeValueMemberS).Value
	if err != nil || tokenString == "" {
		return "", err
	}
	return tokenString, nil
}
