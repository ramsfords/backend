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

func TestQuoteByQuoteId(t *testing.T) {
	res, err := quoteDb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + "1cc284"},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + "2"},
		},
		ProjectionExpression: aws.String("quotes"),
	})
	if err != nil {
		t.Error(err)
	}

	quoteData := &v1.QuoteRequest{}
	data, ok := res.Item["quotes"]
	if !ok {
		t.Error(err)
	}
	err = attributevalue.Unmarshal(data, &quoteData)
	if err != nil {
		t.Error(err)
	}
	t.Error(err)
}
