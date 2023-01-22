package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quoteDb QuoteDb) GetBidWithQuoteByQuoteId(ctx context.Context, businessId string, quoteId string, bidId string) (*model.BidWithQuote, error) {
	res, err := quoteDb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
		},
	})
	if err != nil {
		return nil, err
	}
	quote, ok := res.Item["quoteRequest"]
	quoteData := &v1.QuoteRequest{}
	if ok {
		err := attributevalue.Unmarshal(quote, &quoteData)
		if err != nil {
			return nil, err
		}
	}
	if !ok {
		return nil, nil
	}
	bids, ok := res.Item["bids"]
	if !ok {
		return nil, nil
	}
	unMarshlledBids := map[string]v1.Bid{}
	if ok {

		err := attributevalue.Unmarshal(bids, &unMarshlledBids)
		if err != nil {
			return nil, err
		}
	}
	bidsData := &v1.Bid{}
	for _, bid := range unMarshlledBids {
		if bid.BidId == bidId {
			bidsData = &bid
		}
	}
	quoteWithBid := &model.BidWithQuote{
		QuoteRequest: quoteData,
		Bid:          bidsData,
	}
	return quoteWithBid, nil
}
