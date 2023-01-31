package booking_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func (bookingdb BookingDb) GetBooking(ctx context.Context, bookingId string, businessId string) (*v1.BookingResponse, error) {
	res, err := bookingdb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(bookingdb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + bookingId},
		},
	})
	if err != nil {
		return nil, err
	}

	quoteData := &QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Item, quoteData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	resQuote := &v1.BookingResponse{
		Bid:          quoteData.Bid,
		QuoteRequest: quoteData.QuoteRequest,
		BookingInfo:  quoteData.BookingInfo,
	}
	return resQuote, nil
}
