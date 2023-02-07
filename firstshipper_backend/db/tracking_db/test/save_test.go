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

func TestSave(t *testing.T) {
	trakkingData := v1.ShipmentTracking{
		ShipmentId: "123",
		ShipperId:  "123",
		ReceiverId: "123",
		ShipmentStatusData: []*v1.ShipmentStatusData{
			{
				CurrentLocation:      "123",
				Date:                 "123",
				Comment:              "booked",
				ShipmentStatusOption: v1.ShipmentStatusOption_booked,
			},
		},
	}
	marshalled, err := attributevalue.Marshal(trakkingData)
	if err != nil {
		t.Error(err)
	}
	res, err := trackingDb.DB.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("shipmentTracking_dev"),
		Item: map[string]types.AttributeValue{
			"shipmentId": &types.AttributeValueMemberS{
				Value: trakkingData.ShipmentId,
			},
			"shipmentStatus": marshalled,
		},
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
