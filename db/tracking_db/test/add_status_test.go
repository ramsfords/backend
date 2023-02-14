package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestAddSave(t *testing.T) {
	trakkingData := &v1.ShipmentStatusData{
		CurrentLocation:      "123",
		Date:                 "123",
		Comment:              "booked",
		ShipmentStatusOption: v1.ShipmentStatusOption_booked,
	}
	marshalled, err := attributevalue.Marshal(trakkingData)
	if err != nil {
		t.Error(err)
	}

	// expr := expression.NewBuilder().WithUpdate(
	// 	expression.Add(expression.Name("shipmentStatus.shipmentStatusData"), expression.Value(marshalled)),
	// )
	// exprB, err := expr.Build()
	// if err != nil {
	// 	t.Error(err)
	// }
	res, err := trackingDb.DB.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String("shipmentTracking_dev"),
		Key: map[string]types.AttributeValue{
			"shipmentId": &types.AttributeValueMemberS{
				Value: "123",
			},
		},
		UpdateExpression:          aws.String("SET shipmentStatus.shipmentStatusData = list_append(shipmentStatus.shipmentStatusData, :val)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{":val": &types.AttributeValueMemberL{Value: []types.AttributeValue{marshalled}}},

		// ExpressionAttributeNames: map[string]string{
		// 	"#shipmentStatus": "shipmentStatus",
		// },
	})

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
