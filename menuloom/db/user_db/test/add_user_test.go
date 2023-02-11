package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestAddUser(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		item := &v1.User{
			Email:           "kandelsuren@gmail.com",
			Name:            "surendra kandel",
			Password:        "",
			ConfirmPassword: "",
		}
		data, err := attributevalue.Marshal(item)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Item: map[string]types.AttributeValue{
				"pk":    &types.AttributeValueMemberS{Value: "pk#"},
				"sk":    &types.AttributeValueMemberS{Value: "users#" + item.Email},
				"users": data,
			},
			ConditionExpression: aws.String("attribute_not_exists(sk)"),
		})
		fmt.Println(res, err)
	})
}
