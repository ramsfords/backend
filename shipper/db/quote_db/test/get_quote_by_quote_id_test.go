package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/shipper/business/core/model"
)

func TestQuoteByQuoteId(t *testing.T) {
	res, err := quoteDb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + "kandelsuren@gmail.com"},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + "23015"},
		},
	})
	if err != nil {
		fmt.Println(err)
	}

	quoteData := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Item, quoteData)
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Error(err)
}
