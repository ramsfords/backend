package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quoteDb QuoteDb) GetBidsByQuoteId(ctx context.Context, businessId string, quoteId string) ([]*v1.Bid, error) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			":sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#sk":   "sk",
			"#pk":   "pk",
			"#bids": "bids",
		},
		ProjectionExpression: aws.String("#bids"),
	})
	if err != nil {
		return nil, err
	}
	quote, ok := res.Items[0]["bids"]
	if !ok {
		return nil, nil
	}
	unMarshlled := map[string]v1.Bid{}
	if ok {

		err := attributevalue.Unmarshal(quote, &unMarshlled)
		if err != nil {
			return nil, err
		}
	}
	bids := []*v1.Bid{}
	for _, bid := range unMarshlled {
		bids = append(bids, &bid)
	}
	return bids, nil
}
