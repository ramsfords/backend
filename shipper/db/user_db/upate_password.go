package user_db

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
	"golang.org/x/crypto/bcrypt"
)

func (user UserDb) UpdateUserPassword(ctx context.Context, userData v1.User, businessId string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = user.Client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		ExpressionAttributeNames: map[string]string{
			"#user":                  "user",
			"#hashed_password":       "hashed_password",
			"#new_password_required": "new_password_required",
			"#password_changed_on":   "password_changed_on",
			"#updated_on":            "updated_on",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":new_password":          &types.AttributeValueMemberS{Value: string(hashedPassword)},
			":new_password_required": &types.AttributeValueMemberBOOL{Value: false},
			":password_changed_on":   &types.AttributeValueMemberS{Value: time.Now().String()},
			":updated_on":            &types.AttributeValueMemberS{Value: time.Now().String()}},
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "user#" + userData.Email},
		},
		UpdateExpression:    aws.String("SET #v1.#hashed_password = :new_password, #v1.#new_password_required = :new_password_required, #v1.#password_changed_on = :password_changed_on, #v1.#updated_on = :updated_on"),
		ConditionExpression: aws.String("attribute_exists(sk)"),
		ReturnValues:        types.ReturnValueAllNew,
	})
	if err != nil {
		return err
	}

	return nil
}
