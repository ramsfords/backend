package quote_db

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
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
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#sk":     "sk",
			"#pk":     "pk",
			"#quotes": "quotes",
		},
		ProjectionExpression: aws.String("#quotes"),
	})
	if err != nil {
		return nil, err
	}
	quote, ok := res.Items[0]["quotes"]
	quoteRes := model.QuoteRequest{}
	if ok {

		err := attributevalue.Unmarshal(quote, &quoteRes)
		if err != nil {
			return nil, err
		}
		fmt.Println(quoteRes)
	}
	for _, bid := range quoteRes.Bids {
		if bid.BidId == bidId {
			return bid, nil
		}
	}
	return nil, errors.New("bid not found")
}
