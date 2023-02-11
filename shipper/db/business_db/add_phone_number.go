package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) AddPhoneNumber(ctx context.Context, businessId string, phoneNumber *v1.PhoneNumber) (*v1.PhoneNumber, error) {
	marshalledPhoneNumber, err := attributevalue.Marshal(phoneNumber)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#business":    "business",
			"#phoneNumber": "phoneNumber",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":phoneNumber": marshalledPhoneNumber,
		},
		UpdateExpression: aws.String("SET #business.#phoneNumber = :phoneNumber"),
		ReturnValues:     types.ReturnValueAllNew,
	}
	_, err = businessDb.Client.UpdateItem(ctx, input)
	if err != nil {
		return nil, err
	}

	return phoneNumber, nil
}
