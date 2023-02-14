package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) GetAllLocations(ctx context.Context) ([]*v1.Location, error) {
	qryInput := &dynamodb.QueryInput{
		TableName:              aws.String(locationdb.Config.GetFirstShipperTableName()),
		IndexName:              aws.String("location_index"),
		KeyConditionExpression: aws.String("#location_pk = :pkey"),

		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pkey": &types.AttributeValueMemberS{Value: "location"},
		},
		ExpressionAttributeNames: map[string]string{
			"#location_pk": "location_pk",
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
	return locationsData, err
}
