package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (businessDb BusinessDb) UpdateBusinessName(ctx context.Context, businessId string, businessName string) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#business":     "business",
			"#businessName": "businessName",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":businessName": &types.AttributeValueMemberS{Value: businessName},
		},
		UpdateExpression: aws.String("SET #business.#businessName = :businessName"),
		ReturnValues:     types.ReturnValueAllNew,
	}
	_, err := businessDb.Client.UpdateItem(ctx, input)
	if err != nil {
		return err
	}

	return err
}
