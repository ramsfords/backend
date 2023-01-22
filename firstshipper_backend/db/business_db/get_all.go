package business_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "first-shipper-dev",
// "IndexName": "business_index",
// "KeyConditionExpression": "#business_pk = :business_pk",
// "ExpressionAttributeNames": {"#business_pk":"business_pk"},
// "ExpressionAttributeValues": {":business_pk": {"S":"business"}}
func (businessDb BusinessDb) GetAllBusinesses(ctx context.Context, businessId string) ([]v1.Business, error) {
	urEmailInput := &dynamodb.QueryInput{
		TableName:              aws.String(businessDb.GetFirstShipperTableName()),
		IndexName:              aws.String("business_index"),
		KeyConditionExpression: aws.String("#business_pk = :pkey"),

		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: "business"},
		},
		ExpressionAttributeNames: map[string]string{
			"#business_pk": "business_pk",
		},
		ScanIndexForward: aws.Bool(true),
	}
	res, err := businessDb.Client.Query(ctx, urEmailInput)
	if err != nil {
		return []v1.Business{}, err
	}
	if len(res.Items) == 0 {
		return []v1.Business{}, nil
	}
	return []v1.Business{}, nil
}
