package user_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (user UserDb) UpdateUserConfirmEmail(ctx context.Context, businessId, email string) (bool, error) {
	_, err := user.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		ExpressionAttributeNames: map[string]string{
			"#user":           "user",
			"#email_verified": "email_verified",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email_verified": &types.AttributeValueMemberBOOL{Value: true},
		},
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + email},
		},
		UpdateExpression:    aws.String("SET #v1.#email_verified = :email_verified"),
		ConditionExpression: aws.String(fmt.Sprintf("attribute_exists(%s)", "sk")),
		ReturnValues:        types.ReturnValueUpdatedNew,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
