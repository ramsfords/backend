package user_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (user UserDb) Getuser(ctx context.Context, email string) (*v1.User, error) {
	res, err := user.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(user.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: email},
		},
		ProjectionExpression: aws.String("#email"),
		ExpressionAttributeNames: map[string]string{
			"#email": email,
		},
	})
	if err != nil {
		return nil, err
	}
	userData := &v1.User{}
	data, ok := res.Item["kandelsuren@gmail.com"]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	err = attributevalue.Unmarshal(data, &userData)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
