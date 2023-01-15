package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func TestGetValidate(t *testing.T) {
	res, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(conf.GetMenuloomTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "menu#" + "himalayantasteusa.com"},
			"sk": &types.AttributeValueMemberS{Value: "valid"},
		},
	})
	if err != nil {
		fmt.Print(err)
	}
	valid, ok := res.Item["valid"].(*types.AttributeValueMemberBOOL)
	if !ok {
		fmt.Print(ok)
	}
	fmt.Println(valid)

}
