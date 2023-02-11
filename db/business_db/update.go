package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (businessDb BusinessDb) UpdateBusiness(ctx context.Context, businessId string, business v1.Business) error {
	business.Type = "business"
	businessMarshalled, err := attributevalue.MarshalMap(business)
	if err != nil {
		return err
	}
	qryInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + business.BusinessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + business.BusinessId},
		},
		UpdateExpression: aws.String("SET #business = :business"),
		ExpressionAttributeNames: map[string]string{
			"#business": "business",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":business": &types.AttributeValueMemberM{Value: businessMarshalled},
		},
		// items already not in the db by table sk which is same as "sk"
		ConditionExpression: aws.String("attribute_exists(sk)"),
		ReturnValues:        "ALL_NEW",
	}
	result, err := businessDb.Client.UpdateItem(ctx, qryInput)
	if err != nil {
		return err
	}
	var updatedBusinessData *v1.Business
	attributevalue.UnmarshalMap(result.Attributes, &updatedBusinessData)
	return nil
}
