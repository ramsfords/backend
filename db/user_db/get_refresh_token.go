package user_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/foundations/errs"
)

// "TableName": "firstshipper-dev",
// "IndexName": "refreshTokenIndex",
// "KeyConditionExpression": "#16011 = :16011",
// "ProjectionExpression": "#16010",
// "ExpressionAttributeNames": {"#16010":"refreshToken","#16011":"refreshTokenPk"},
// "ExpressionAttributeValues": {":16011": {"S":"1345448f-87f5-4e98-a5ef-9301eb250a01"}}
func (userdb UserDb) GetRefreshToken(ctx context.Context, userId string) (string, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(userdb.GetFirstShipperTableName()),
		IndexName:              aws.String("refreshTokenIndex"),
		KeyConditionExpression: aws.String("#refreshTokenPK = :refreshTokenPK"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":refreshTokenPK": &types.AttributeValueMemberS{Value: userId},
		},
		ExpressionAttributeNames: map[string]string{
			"#refreshToken":   "refreshToken",
			"#refreshTokenPK": "refreshTokenPk",
		},
		ProjectionExpression: aws.String("#refreshToken"),
	}
	refreshToken, err := userdb.Client.Query(ctx, input)
	// getToken from response
	var tokenStr string
	if refreshToken != nil && len(refreshToken.Items) > 0 {
		token, ok := refreshToken.Items[0]["refreshToken"]
		if !ok {
			return "", err
		}
		// unmarshal token
		tokenStr = token.(*types.AttributeValueMemberS).Value
		if err != nil || tokenStr == "" {
			return "", err
		}
	} else {
		return "", errs.ErrDataNotFound
	}
	return tokenStr, nil
}
