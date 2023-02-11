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

func TestCreateitem(t *testing.T) {
	t.Run("create create item", func(t *testing.T) {
		datas := &v1.Item{
			Name:        "tandoor",
			Description: "test",
			Price:       "1.00",
			SpiceLevel:  "mild",
			Images:      []*v1.Image{},
			Categories:  []string{"test"},
		}
		data, err := attributevalue.Marshal(&datas)
		if err != nil {
			t.Fatal(err)
		}
		res, err := db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
			TableName: aws.String(conf.GetMenuloomTableName()),
			Item: map[string]types.AttributeValue{
				"pk":    &types.AttributeValueMemberS{Value: conf.GetMenuloomPKPrefix() + "himalayen.menuloom.com"},
				"sk":    &types.AttributeValueMemberS{Value: "item#" + datas.Name},
				"items": data,
			},
			ConditionExpression: aws.String("attribute_not_exists(sk)"),
		})
		fmt.Println(res, err)
	})
}
