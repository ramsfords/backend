package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestNewQ(t *testing.T) {
	input := v1.QuoteRequest{
		QuoteId:    "1",
		BusinessId: "1cc284",
	}
	marshaledQuote, err := attributevalue.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	saveQute := models.SaveQuote{
		SavedQuoteID: "2",
	}

	marshaledsaveQute, err := attributevalue.Marshal(saveQute)
	if err != nil {
		t.Fatal(err)
	}
	bids := map[string]v1.Bid{}
	bids["roadrunner"] = v1.Bid{
		CarrierName: "roadrunner",
	}

	marshaledBid, err := attributevalue.Marshal(bids)
	if err != nil {
		t.Fatal(err)
	}

	res, err := quoteDb.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":             &types.AttributeValueMemberS{Value: "pk#" + input.BusinessId},
			"sk":             &types.AttributeValueMemberS{Value: "quote#" + input.QuoteId},
			"quoteRequest":   marshaledQuote,
			"rapidSaveQuote": marshaledsaveQute,
			"bids":           marshaledBid,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
