package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
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
	bidMap := make(map[string]v1.Bid)
	for _, bid := range qtReq.Bids {
		bidMap[bid.CarrierName] = bid
	}
	marshalledBids, err := attributevalue.Marshal(bidMap)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":             &types.AttributeValueMemberS{Value: "pk#" + qtReq.QuoteRequest.BusinessId},
			"sk":             &types.AttributeValueMemberS{Value: "quote#" + qtReq.QuoteRequest.QuoteId},
			"quoteRequest":   marshlledQtReq,
			"rapidSaveQuote": marshlledRapidSaveQuote,
			"bids":           marshalledBids,
		},
	}
	_, err = quotedb.Client.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
