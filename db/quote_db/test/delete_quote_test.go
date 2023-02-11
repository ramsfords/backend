package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestDeleteQuote(t *testing.T) {
	input := v1.QuoteRequest{
		QuoteId:    "1",
		BusinessId: "1cc284",
	}
	res, err := quoteDb.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + input.BusinessId},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + input.QuoteId},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
