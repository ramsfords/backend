package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/shipper/business/core/model"
)

func (quotedb QuoteDb) SaveQuote(ctx context.Context, qtReq *model.QuoteRequest) error {
	marshlledQtReq, err := attributevalue.Marshal(qtReq.QuoteRequest)
	if err != nil {
		return err
	}
	marshlledRapidSaveQuote, err := attributevalue.Marshal(qtReq.RapidSaveQuote)
	if err != nil {
		return err
	}
	marshlledSaveQuoteResponose, err := attributevalue.Marshal(qtReq.SaveQuoteResponse)
	if err != nil {
		return err
	}
	marshalledBids, err := attributevalue.Marshal(qtReq.Bids)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":                &types.AttributeValueMemberS{Value: "pk#" + qtReq.QuoteRequest.BusinessId},
			"sk":                &types.AttributeValueMemberS{Value: "quote#" + qtReq.QuoteRequest.QuoteId},
			"quote_pk":          &types.AttributeValueMemberS{Value: qtReq.QuoteRequest.QuoteId},
			"quoteRequest":      marshlledQtReq,
			"rapidSaveQuote":    marshlledRapidSaveQuote,
			"bids":              marshalledBids,
			"saveQuoteResponse": marshlledSaveQuoteResponose,
			"rapidBooking":      &types.AttributeValueMemberM{Value: nil},
			"bookingInfo":       &types.AttributeValueMemberM{Value: nil},
			"bid":               &types.AttributeValueMemberM{Value: nil},
		},
	}
	_, err = quotedb.Client.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
