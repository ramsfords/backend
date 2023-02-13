package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/business/core/model"
)

// "TableName": "firstshipper-dev",
// "IndexName": "quote_index",
// "KeyConditionExpression": "#83340 = :83340",
// "ExpressionAttributeNames": {"#83340":"quote_pk"},
// "ExpressionAttributeValues": {":83340": {"S":"50099"}}

func (quoteDb QuoteDb) GetQuoteByQuoteId(ctx context.Context, quoteId string) (*model.QuoteRequest, error) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		IndexName:              aws.String("quote_index"),
		KeyConditionExpression: aws.String("#quote_pk = :quote_pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quote_pk": &types.AttributeValueMemberS{Value: quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#quote_pk": "quote_pk",
		},
	})
	if err != nil {
		return nil, err
	}
	bidsRequest := &model.QuoteRequest{}
	if len(res.Items) == 0 {
		return bidsRequest, nil
	}
	err = attributevalue.UnmarshalMap(res.Items[0], bidsRequest)
	if err != nil {
		return nil, err
	}
	quoteWithBidsRes := &model.QuoteRequest{
		QuoteRequest:      bidsRequest.QuoteRequest,
		Bids:              bidsRequest.Bids,
		BookingInfo:       bidsRequest.BookingInfo,
		RapidSaveQuote:    bidsRequest.RapidSaveQuote,
		SaveQuoteResponse: bidsRequest.SaveQuoteResponse,
		RapidBooking:      bidsRequest.RapidBooking,
		Business:          bidsRequest.Business,
	}

	return quoteWithBidsRes, nil
}
