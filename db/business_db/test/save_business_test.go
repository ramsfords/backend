package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func TestSaveBusiness(t *testing.T) {
	t.Log("TestSaveBusiness")
	business := &v1.Business{
		BusinessId:                        "1cc284",
		BusinessName:                      "Ramsfords",
		BusinessEmail:                     "ramsford@gmail.com",
		AccountingEmail:                   "ramsford@gmail.com",
		NeedsDefaultPickupAddressUpdate:   false,
		NeedsDefaultDeliveryAddressUpdate: true,
		DefaultPickupAddress: &v1.Address{
			AddressLine1: "1cc284",
			StreetName:   "123 Main St",
			City:         "San Francisco",
			State:        "CA",
			ZipCode:      "94105",
		},
		DefaultDeliveryAddress: &v1.Address{
			AddressLine1: "1cc284",
			StreetName:   "123 Main St",
			City:         "San Francisco",
			State:        "CA",
			ZipCode:      "94105",
		},
		BillingAddress: &v1.Address{
			AddressLine1: "1cc284",
			StreetName:   "123 Main St",
			City:         "San Francisco",
			State:        "CA",
			ZipCode:      "94105",
		},
		Address: &v1.Address{
			AddressLine1: "1cc284",
			StreetName:   "123 Main St",
			City:         "San Francisco",
			State:        "CA",
			ZipCode:      "94105",
		},

		CreatedAt: time.Now().Format(time.RFC3339),
	}
	business.Type = "business"
	itemMarshalled, err := attributevalue.MarshalMap(business)
	if err != nil {
		t.Error(err)
	}
	putItem := &dynamodb.PutItemInput{
		TableName: aws.String(conf.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":       &types.AttributeValueMemberS{Value: "pk#" + business.BusinessId},
			"sk":       &types.AttributeValueMemberS{Value: "business#" + business.BusinessId},
			"business": &types.AttributeValueMemberM{Value: itemMarshalled},
		},
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", "pk")),
	}
	_, err = db.Client.PutItem(context.Background(), putItem)
	if err != nil {
		t.Error(err)
	}
}
