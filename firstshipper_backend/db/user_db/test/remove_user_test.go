package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestRemoveUser(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		user := &v1.User{
			Email: "kandelsuren@gmail.com",
		}

		res, err := db.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
			TableName: aws.String(conf.GetFirstShipperTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#" + "1cc284"},
				"sk": &types.AttributeValueMemberS{Value: "user#" + user.Email},
			},
		})
		fmt.Println(res, err)
	})
}
