package booking_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (bookingdb BookingDb) GetBooking(ctx context.Context, bookingId string) (*v1.BookingResponse, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.Config.GetFirstShipperTableName()),
		IndexName:              aws.String("book_index"),
		KeyConditionExpression: aws.String("#booking_pk = :pkey and #booking_sk = :skey"),

		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: "booking"},
			":skey": &types.AttributeValueMemberS{Value: bookingId},
		},
		ExpressionAttributeNames: map[string]string{
			"#booking_pk": "booking_pk",
			"#booking_sk": "booking_sk",
		},
		ScanIndexForward: aws.Bool(true),
	}

	res, err := bookingdb.Client.Query(ctx, qryInput)
	if err != nil {
		return &v1.BookingResponse{}, err
	}
	if len(res.Items) == 0 {
		return &v1.BookingResponse{}, nil
	}
	bookingData := &v1.BookingResponse{}
	err = attributevalue.UnmarshalMap(res.Items[0], &bookingData)
	if err != nil {
		return &v1.BookingResponse{}, err
	}

	return bookingData, nil
}
