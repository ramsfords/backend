package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/business/core/model"
)

// "TableName": "firstshipper-dev",
// "KeyConditionExpression": "#3da90 = :3da90 And begins_with(#3da91, :3da91)",
// "FilterExpression": "attribute_exists(#3da92)",
// "ExpressionAttributeNames": {"#3da90":"pk","#3da91":"sk","#3da92":"booking_pk"},
// "ExpressionAttributeValues": {":3da90": {"S":"pk#kandelsuren@gmail.com"},":3da91": {"S":"quote"}}
func TestGetBooking(t *testing.T) {
	businessId := "kandelsuren@gmail.com"
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(conf.GetFirstShipperTableName()),
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

	res, err := db.Client.Query(context.Background(), qryInput)
	if err != nil {
		fmt.Println(err.Error())
	}
	if len(res.Items) == 0 {
		fmt.Println(err)
	}
	bookingsData := []*model.QuoteRequest{}
	err = attributevalue.UnmarshalListOfMaps(res.Items, &bookingsData)
	if err != nil {
		fmt.Println(err)
	}
}
