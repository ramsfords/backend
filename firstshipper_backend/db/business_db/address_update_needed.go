package business_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (businessDb BusinessDb) UpdateBusinessAddressUpdateNeeded(ctx context.Context, businessId string) error {
	qryInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		UpdateExpression: aws.String("SET #business.#needs_address_update = :needs_address_update"),
		ExpressionAttributeNames: map[string]string{
			"#business":             "business",
			"#needs_address_update": "needs_address_update",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":needs_address_update": &types.AttributeValueMemberBOOL{Value: false},
		},
		// items already not in the db by table sk which is same as "sk"
		ConditionExpression: aws.String("attribute_exists(sk)"),
		ReturnValues:        "ALL_NEW",
	}
	_, err := businessDb.Client.UpdateItem(ctx, qryInput)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return nil
}
