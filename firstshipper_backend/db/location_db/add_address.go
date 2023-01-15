package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) AddLocationAddress(ctx context.Context, businessId string, address *v1.Address) error {
	marshalledAddress, err := attributevalue.MarshalMap(address)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "location#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#address":  "address",
			"#location": "location",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":address": &types.AttributeValueMemberM{Value: marshalledAddress},
		},
		UpdateExpression: aws.String("SET #location.#address = :address"),
	}
	_, err = locationdb.Client.UpdateItem(ctx, input)
	return err
}
