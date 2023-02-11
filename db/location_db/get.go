package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "first-shipper-dev",
// "IndexName": "location_index",
// "KeyConditionExpression": "#location_pk = :location_pk And #location_sk = :location_sk",
// "ExpressionAttributeNames": {"#location_pk":"location_pk","#location_sk":"location_sk"},
// "ExpressionAttributeValues": {":location_pk": {"S":"location"},":location_sk": {"S":"asdf"}}
func (locationdb LocationDb) GetLocation(ctx context.Context, businessId string, locationId string) (*v1.Location, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "location#" + businessId},
		},
	}
	res, err := locationdb.Client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}
	loc := &v1.Location{}
	err = attributevalue.UnmarshalMap(res.Item, loc)
	if err != nil {
		return nil, err
	}
	return loc, nil
}
