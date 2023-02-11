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

func (db UserDb) Getuser(ctx context.Context, email string) (*v1.User, error) {
	res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(db.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: db.GetMenuloomPKPrefix()},
			"sk": &types.AttributeValueMemberS{Value: "#users"},
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
