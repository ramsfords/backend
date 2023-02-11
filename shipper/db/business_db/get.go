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

func (businessDb BusinessDb) GetBusiness(ctx context.Context, businessId string) (*v1.Business, error) {
	bisInput := &dynamodb.GetItemInput{
		TableName: aws.String(businessDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "business#" + businessId},
		},
		ProjectionExpression: aws.String("business"),
	}
	res, err := businessDb.Client.GetItem(ctx, bisInput)
	if err != nil {
		return nil, err
	}
	if len(res.Item) == 0 {
		return nil, err
	}
	business, ok := res.Item["business"]
	if !ok {
		return nil, err
	}
	businessData := &v1.Business{}
	err = attributevalue.Unmarshal(business, businessData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return businessData, nil
}
