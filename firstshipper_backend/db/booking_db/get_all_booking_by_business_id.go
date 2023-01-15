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
func (bookingdb BookingDb) GetAllBookingsByBusinessId(ctx context.Context, businessId string) ([]*v1.BookingResponse, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.Config.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		FilterExpression:       aws.String("attribute_exists(#booking_sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "quote#"},
		},
		ExpressionAttributeNames: map[string]string{
			"#booking_sk": "booking_sk",
			"#pk":         "pk",
			"#sk":         "sk",
		},
		ScanIndexForward: aws.Bool(true),
	}

	res, err := bookingdb.Client.Query(ctx, qryInput)
	if err != nil {
		return []*v1.BookingResponse{}, err
	}
	if len(res.Items) == 0 {
		return []*v1.BookingResponse{}, nil
	}
	bookingsData := []*v1.BookingResponse{}
	err = attributevalue.UnmarshalListOfMaps(res.Items, &bookingsData)
	if err != nil {
		return []*v1.BookingResponse{}, err
	}
	return bookingsData, nil
}
