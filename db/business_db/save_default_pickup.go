package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) SaveDefaultPickup(ctx context.Context, businessId string, address v1.Location) error {
	addressMarshalled, err := attributevalue.MarshalMap(address)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	input1 := dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		UpdateExpression: aws.String("SET #default_pickup_address = :default_pickup_address"),
		ExpressionAttributeNames: map[string]string{
			"#default_pickup_address": "default_pickup_address",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":default_pickup_address": &types.AttributeValueMemberM{Value: addressMarshalled},
		},
		ConditionExpression: aws.String("attribute_exists(sk)"),
	}
	_, err = businessDb.Client.UpdateItem(ctx, &input1)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	qryInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		UpdateExpression: aws.String("SET #business.#needsDefaultPickupAddressUpdate = :needsDefaultPickupAddressUpdate"),
		ExpressionAttributeNames: map[string]string{
			"#business":                        "business",
			"#needsDefaultPickupAddressUpdate": "needsDefaultPickupAddressUpdate",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":needsDefaultPickupAddressUpdate": &types.AttributeValueMemberBOOL{Value: false},
		},
		// items already not in the db by table sk which is same as "sk"
		ConditionExpression: aws.String("attribute_exists(sk)"),
		ReturnValues:        "ALL_NEW",
	}
	_, err = businessDb.Client.UpdateItem(ctx, qryInput)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return nil
}
