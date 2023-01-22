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

func (businessDb BusinessDb) AddAddress(ctx context.Context, businessId string, address v1.Address) error {
	addressMarshalled, err := attributevalue.MarshalMap(address)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String(businessDb.GetFirstShipperTableName()),
		UpdateExpression: aws.String("SET #location = :location"),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "location#" + businessId},
		},
		ExpressionAttributeNames: map[string]string{
			"#location": "location",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":location": &types.AttributeValueMemberM{Value: addressMarshalled},
		},
		ConditionExpression: aws.String("attribute_exists(pk)"),
	}
	_, err = businessDb.Client.UpdateItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
