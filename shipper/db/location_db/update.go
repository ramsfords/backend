package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) UpdateLocation(ctx context.Context, businessId string, data *v1.Address) error {
	marshalledAddress, err := attributevalue.Marshal(data)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#business":                        "business",
			"#defaultPickupAddress":            "defaultPickupAddress",
			"#needsDefaultPickupAddressUpdate": "needsDefaultPickupAddressUpdate",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":defaultPickupAddress":            marshalledAddress,
			":needsDefaultPickupAddressUpdate": &types.AttributeValueMemberBOOL{Value: false},
		},
		UpdateExpression: aws.String("SET #business.#defaultPickupAddress = :defaultPickupAddress, #business.#needsDefaultPickupAddressUpdate = :needsDefaultPickupAddressUpdate"),
	}
	_, err = locationdb.Client.UpdateItem(ctx, input)
	return err
}
