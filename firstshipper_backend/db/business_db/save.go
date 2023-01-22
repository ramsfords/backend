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

func (businessDb BusinessDb) SaveBusiness(ctx context.Context, business v1.Business) error {
	business.Type = "business"
	itemMarshalled, err := attributevalue.MarshalMap(business)
	if err != nil {
		return err
	}
	putItem := &dynamodb.PutItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":       &types.AttributeValueMemberS{Value: "pk#" + business.BusinessId},
			"sk":       &types.AttributeValueMemberS{Value: "business#" + business.BusinessId},
			"business": &types.AttributeValueMemberM{Value: itemMarshalled},
		},
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", "pk")),
	}
	_, err = businessDb.Client.PutItem(ctx, putItem)
	if err != nil {
		return err
	}
	return nil
}
