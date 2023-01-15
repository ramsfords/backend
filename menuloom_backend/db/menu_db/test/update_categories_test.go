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

func TestUpdateCategories(t *testing.T) {
	t.Run("create restaurant", func(t *testing.T) {
		rawData := v1.Category{
			Name:        "dinners",
			ServingTime: v1.ServingTime_ALLDAY,
		}
		data, err := attributevalue.Marshal(&rawData)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#"},
				"sk": &types.AttributeValueMemberS{Value: "category#"},
			},
			ExpressionAttributeNames: map[string]string{
				"#categories": "categories",
				"#name":       rawData.Name,
			},
			ExpressionAttributeValues: map[string]types.AttributeValue{
				":categories": data,
			},
			UpdateExpression: aws.String("SET #categories.#name = :categories"),
		})
		fmt.Println(res, err)
	})
}
