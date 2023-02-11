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

func TestCreateUser(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		item := &v1.User{
			Email:           "rakeshneupane2045@gmail.coM",
			Name:            "surendra kandel",
			Password:        "",
			ConfirmPassword: "",
		}
		data, err := attributevalue.Marshal(item)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(conf.GetFirstShipperTableName()),
			Item: map[string]types.AttributeValue{
				"pk":      &types.AttributeValueMemberS{Value: "pk#" + "1cc284"},
				"sk":      &types.AttributeValueMemberS{Value: "user#" + item.Email},
				"user_sk": &types.AttributeValueMemberS{Value: item.Email},
				"users":   data,
			},
			ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", "pk")),
		})
		fmt.Println(res, err)
	})
}
