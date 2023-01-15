package booking_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (bookingdb BookingDb) SaveBooking(ctx context.Context, bookingRes *v1.BookingResponse) error {
	bookingResMarashalled, err := attributevalue.MarshalMap(bookingRes)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(bookingdb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":         &types.AttributeValueMemberS{Value: "business#" + bookingRes.BusinessId},
			"sk":         &types.AttributeValueMemberS{Value: fmt.Sprintf("quote#%s", bookingRes.QuoteId+"#booking")},
			"booking_pk": &types.AttributeValueMemberS{Value: "booking"},
			"booking_sk": &types.AttributeValueMemberS{Value: bookingRes.QuoteId},
			"book_index": &types.AttributeValueMemberM{Value: bookingResMarashalled},
		},
	}
	_, err = bookingdb.Client.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
