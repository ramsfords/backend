package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) AddLocationAddress(ctx context.Context, businessId string, address *v1.Address) (*v1.Address, error) {
	marshalledAddress, err := attributevalue.Marshal(address)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#business":           "business",
			"#address":            "address",
			"#needsAddressUpdate": "needsAddressUpdate",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":address":            marshalledAddress,
			":needsAddressUpdate": &types.AttributeValueMemberBOOL{Value: false},
		},
		UpdateExpression: aws.String("SET #business.#address = :address, #business.#needsAddressUpdate = :needsAddressUpdate"),
		ReturnValues:     types.ReturnValueAllNew,
	}
	_, err = locationdb.Client.UpdateItem(ctx, input)
	if err != nil {
		return nil, err
	}

	return address, nil
}
