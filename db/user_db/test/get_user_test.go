package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestGetRestaurant(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		res, err := db.Client.Query(context.Background(), &dynamodb.QueryInput{
			TableName:              aws.String(conf.GetFirstShipperTableName()),
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
			t.Fatal(err, res)
		}

		userData := &v1.User{}
		data, ok := res.Items[0]["users"]
		if !ok {
			t.Fatal("no data found")
		}
		err = attributevalue.Unmarshal(data, &userData)
		if err != nil {
			t.Fatal(err, res, data)
		}

	})
}
