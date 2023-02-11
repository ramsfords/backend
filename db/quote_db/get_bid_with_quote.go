package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quoteDb QuoteDb) GetBidByQuoteId(ctx context.Context, businessId string, quoteId string, bidId string) (*model.QuoteRequest, error) {
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
	quoteData := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Item, quoteData)
	if err != nil {
		return nil, err
	}
	quoteWithBid := &model.QuoteRequest{
		QuoteRequest: quoteData.QuoteRequest,
		Bid:          getBidFromBids(quoteData.Bids, bidId),
		BookingInfo:  quoteData.BookingInfo,
	}
	return quoteWithBid, nil
}
func getBidFromBids(bids []*v1.Bid, bidId string) *v1.Bid {
	for _, bid := range bids {
		if bid.BidId == bidId {
			return bid
		}
	}
	return nil
}
