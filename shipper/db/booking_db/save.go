package booking_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/shipper/business/core/model"
)

func (bookingdb BookingDb) SaveBooking(ctx context.Context, bookingRes *model.QuoteRequest) error {
	bookingMarshlled, err := attributevalue.Marshal(bookingRes.QuoteRequest)
	if err != nil {
		return err
	}
	rapidSavedMarshalled, err := attributevalue.Marshal(bookingRes.RapidSaveQuote)
	if err != nil {
		return err
	}
	bookingInfoMarshed, err := attributevalue.Marshal(bookingRes.BookingInfo)
	if err != nil {
		return err
	}
	bidsMarshed, err := attributevalue.Marshal(bookingRes.Bids)
	if err != nil {
		return err
	}
	rapidBookingMarshalled, err := attributevalue.Marshal(bookingRes.RapidBooking)
	if err != nil {
		return err
	}
	rapidSaveQuoteResponseMarshalled, err := attributevalue.Marshal(bookingRes.SaveQuoteResponse)
	if err != nil {
		return err
	}
	bidMarshalled, err := attributevalue.Marshal(bookingRes.Bid)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(bookingdb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":                &types.AttributeValueMemberS{Value: "pk#" + bookingRes.QuoteRequest.BusinessId},
			"sk":                &types.AttributeValueMemberS{Value: "quote#" + bookingRes.QuoteRequest.QuoteId},
			"quote_pk":          &types.AttributeValueMemberS{Value: bookingRes.Bid.QuoteId},
			"booking_pk":        &types.AttributeValueMemberS{Value: bookingRes.QuoteRequest.QuoteId},
			"quoteRequest":      bookingMarshlled,
			"rapidSaveQuote":    rapidSavedMarshalled,
			"bids":              bidsMarshed,
			"saveQuoteResponse": rapidSaveQuoteResponseMarshalled,
			"rapidBooking":      rapidBookingMarshalled,
			"bookingInfo":       bookingInfoMarshed,
			"bid":               bidMarshalled,
		},
	}
	_, err = bookingdb.Client.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
