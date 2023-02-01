package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
)

func (quoteDb QuoteDb) GetBidsWithQuoteByQuoteId(ctx context.Context, businessId string, quoteId string) (*model.QuoteRequest, error) {
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
	quoteWithBids := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Item, quoteWithBids)
	if err != nil {
		return nil, err
	}
	quoteWithBidsRes := &model.QuoteRequest{
		QuoteRequest: quoteWithBids.QuoteRequest,
		Bids:         quoteWithBids.Bids,
	}
	return quoteWithBidsRes, nil
}
