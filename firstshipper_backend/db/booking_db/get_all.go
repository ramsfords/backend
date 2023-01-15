package booking_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "first-shipper-dev",
// "KeyConditionExpression": "#pk = :pk And begins_with(#sk, :sk)",
// "FilterExpression": "attribute_exists(#booking_sk)",
// "ExpressionAttributeNames": {"#pk":"pk","#sk":"sk","#booking_sk":"booking_sk"},
// "ExpressionAttributeValues": {":pk": {"S":"business#7e3da8"},":sk": {"S":"quote#"}}
func (bookingdb BookingDb) GetAllBookings(ctx context.Context) ([]*v1.BookingResponse, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.Config.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#booking_pk = :pk"),
		FilterExpression:       aws.String("attribute_exists(#booking_sk)"),
		IndexName:              aws.String("book_index"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "business"},
		},
		ExpressionAttributeNames: map[string]string{
			"#booking_pk": "booking_sk",
		},
		ScanIndexForward: aws.Bool(true),
	}

	res, err := bookingdb.Client.Query(ctx, qryInput)
	if err != nil {
		return nil, err
	}
	if len(res.Items) == 0 {
		return nil, nil
	}
	bookingsData := []*v1.BookingResponse{}
	err = attributevalue.UnmarshalListOfMaps(res.Items, &bookingsData)
	if err != nil {
		return nil, err
	}
	return bookingsData, nil
}
