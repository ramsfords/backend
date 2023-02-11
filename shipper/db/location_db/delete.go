package location_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (locationdb LocationDb) DeleteLocation(ctx context.Context, locationId string, businessId string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(locationdb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "location#" + locationId},
		},
	}
	_, err := locationdb.Client.DeleteItem(ctx, input)
	return err
}
