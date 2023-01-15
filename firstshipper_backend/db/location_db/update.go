package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (locationdb LocationDb) UpdateLocation(ctx context.Context, businessId string, data *v1.Location) error {
	marshalledLocation, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk":          &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk":          &types.AttributeValueMemberS{Value: "location#" + data.Id},
			"location_pk": &types.AttributeValueMemberS{Value: "location"},
			"location_sk": &types.AttributeValueMemberS{Value: data.Id},
			"location":    &types.AttributeValueMemberM{Value: marshalledLocation},
		},
		ExpressionAttributeNames: map[string]string{
			"#pk":          "pk",
			"#sk":          "sk",
			"#location_pk": "location_pk",
			"#location_sk": "location_sk",
			"#location":    "location",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk":          &types.AttributeValueMemberS{Value: "business#" + businessId},
			":sk":          &types.AttributeValueMemberS{Value: "location#" + data.Id},
			":location_pk": &types.AttributeValueMemberS{Value: "location"},
			":location_sk": &types.AttributeValueMemberS{Value: data.Id},
			":location":    &types.AttributeValueMemberM{Value: marshalledLocation},
		},
	}
	_, err = locationdb.Client.UpdateItem(ctx, input)
	return err
}
