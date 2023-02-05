package booking_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
)

// "TableName": "first-shipper-dev",
// "KeyConditionExpression": "#pk = :pk And begins_with(#sk, :sk)",
// "FilterExpression": "attribute_exists(#booking_sk)",
// "ExpressionAttributeNames": {"#pk":"pk","#sk":"sk","#booking_sk":"booking_sk"},
// "ExpressionAttributeValues": {":pk": {"S":"business#7e3da8"},":sk": {"S":"quote#"}}
func (bookingdb BookingDb) GetAllBookingsByBusinessId(ctx context.Context, businessId string) ([]*model.QuoteRequest, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(bookingdb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "quote"},
		},

		FilterExpression: aws.String("attribute_exists(booking_pk)"),
		ExpressionAttributeNames: map[string]string{
			"#pk": "pk",
			"#sk": "sk",
		},

		ScanIndexForward: aws.Bool(false),
	}

	res, err := bookingdb.Client.Query(ctx, qryInput)
	if err != nil {
		return []*model.QuoteRequest{}, err
	}
	if len(res.Items) == 0 {
		return []*model.QuoteRequest{}, nil
	}
	bookingsData := []*model.QuoteRequest{}
	err = attributevalue.UnmarshalListOfMaps(res.Items, &bookingsData)
	if err != nil {
		return []*model.QuoteRequest{}, err
	}
	return bookingsData, nil
}
