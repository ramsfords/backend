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
		res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
			TableName: aws.String(conf.GetFirstShipperTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#127.0.0.1:3000"},
				"sk": &types.AttributeValueMemberS{Value: "users"},
			},
			ProjectionExpression: aws.String("#email"),
			ExpressionAttributeNames: map[string]string{
				"#email": "kandelsuren@gmail.com",
			},
		})
		if err != nil {
			t.Fatal(err, res)
		}

		userData := &v1.User{}
		data, ok := res.Item["kandelsuren@gmail.com"]
		err = attributevalue.Unmarshal(data, &userData)
		if err != nil {
			t.Fatal(err, res, data, ok)
		}

	})
}
