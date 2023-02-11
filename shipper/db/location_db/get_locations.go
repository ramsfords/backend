package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) GetLocations(ctx context.Context, businessId string) ([]*v1.Location, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(locationdb.Config.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#business_pk = :pkey and begins_with(#business_sk, :skey)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: businessId},
			":skey": &types.AttributeValueMemberS{Value: "location"},
		},
		ExpressionAttributeNames: map[string]string{
			"#business_pk": "pk",
			"#business_sk": "sk",
		},
		ScanIndexForward: aws.Bool(true),
	}
	res, err := locationdb.Client.Query(ctx, qryInput)
	if err != nil {
		return nil, err
	}
	if len(res.Items) == 0 {
		return nil, nil
	}
	locationsData := []*v1.Location{}
	err = attributevalue.UnmarshalListOfMaps(res.Items, &locationsData)
	if err != nil {
		return nil, err
	}
	return locationsData, nil
}
