package booking_db

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/business/core/model"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (bookingdb BookingDb) GetBooking(ctx context.Context, quoteId string) (*v1.BookingResponse, error) {
	res, err := bookingdb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.GetFirstShipperTableName()),
		IndexName:              aws.String("booking_index"),
		KeyConditionExpression: aws.String("#booking_pk = :booking_pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":booking_pk": &types.AttributeValueMemberS{Value: quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#booking_pk": "booking_pk",
		},
	})
	if err != nil {
		return nil, err
	}
	bookingReq := &model.QuoteRequest{}
	if len(res.Items) < 1 {
		return nil, errors.New("booking not found")
	}
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
