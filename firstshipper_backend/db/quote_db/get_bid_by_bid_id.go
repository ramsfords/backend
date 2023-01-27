package quote_db

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "first-shipper-dev",
// "IndexName": "quote_index",
// "KeyConditionExpression": "#quote_pk = :quote_pk",
// "FilterExpression": "#pk = :pk",
// "ExpressionAttributeNames": {"#quote_pk":"quote_pk","#pk":"pk"},
// "ExpressionAttributeValues": {":quote_pk": {"S":"quote"},":pk": {"S":"business#1cc284"}}
func (quoteDb QuoteDb) GetBidByBidID(ctx context.Context, businessId string, quoteId string, bidId string) (*v1.Bid, error) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#bids": "bids",
		},
		ProjectionExpression: aws.String("#bids"),
	})
	if err != nil {
		return nil, err
	}
	bids, ok := res.Items[0]["bids"]
	bidRes := []*v1.Bid{}
	if ok {

		err := attributevalue.Unmarshal(bids, bidRes)
		if err != nil {
			return nil, err
		}
	}
	for _, bid := range bidRes {
		if bid.BidId == bidId {
			return bid, nil
		}
	}
	return nil, errors.New("bid not found")
}
