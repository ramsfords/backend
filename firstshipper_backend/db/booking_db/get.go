package booking_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

type QuoteRequest struct {
	Bids              []*v1.Bid                 `json:"bids" dynamodbav:"bids"`
	Bid               *v1.Bid                   `json:"bid" dynamodbav:"bid"`
	QuoteRequest      *v1.QuoteRequest          `json:"quoteRequest" dynamodbav:"quoteRequest"`
	RapidSaveQuote    *models.SaveQuote         `json:"SaveQuote" dynamodbav:"rapidSaveQuote"`
	SaveQuoteResponse *models.SaveQuoteResponse `json:"saveQuoteResponse" dynamodbav:"saveQuoteResponse"`
	RapidBooking      *models.DispatchResponse  `json:"Booking" dynamodbav:"rapidBooking"`
	BookingInfo       *v1.BookingInfo           `json:"bookingInfo" dynamodbav:"bookingInfo"`
}

func (bookingdb BookingDb) GetBooking(ctx context.Context, quoteId string) (*v1.BookingResponse, error) {
	res, err := bookingdb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.GetFirstShipperTableName()),
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
	bookingReq := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Items[0], bookingReq)
	if err != nil {
		return nil, err
	}
	resQuote := &v1.BookingResponse{
		Bid:          bookingReq.Bid,
		QuoteRequest: bookingReq.QuoteRequest,
		BookingInfo:  bookingReq.BookingInfo,
	}
	return resQuote, nil
}
