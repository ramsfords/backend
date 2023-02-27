package user_db

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (userdb UserDb) Getuser(ctx context.Context, email string) (*v1.User, error) {
	res, err := userdb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(userdb.GetFirstShipperTableName()),
		IndexName:              aws.String("user_index"),
		KeyConditionExpression: aws.String("#user_pk = :user_pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":user_pk": &types.AttributeValueMemberS{Value: "kandelsuren@gmail.com"},
		},
		ExpressionAttributeNames: map[string]string{
			"#user_pk": "user_pk",
			"#users":   "users",
		},
		ProjectionExpression: aws.String("#users"),
	})
	if err != nil {
		return nil, err
	}
	if len(res.Items) == 0 {
		return nil, errors.New("no data found")
	}
	userData := &v1.User{}
	data, ok := res.Items[0]["users"]
	if !ok {
		return nil, errors.New("no data found")
	}
	err = attributevalue.Unmarshal(data, &userData)
	if err != nil {
		return nil, errors.New("no data found")
	}
	return userData, nil
}
