package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/business/core/model"
)

// "TableName": "first-shipper-dev",
// "IndexName": "quote_index",
// "KeyConditionExpression": "#quote_pk = :quote_pk",
// "FilterExpression": "#pk = :pk",
// "ExpressionAttributeNames": {"#quote_pk":"quote_pk","#pk":"pk"},
// "ExpressionAttributeValues": {":quote_pk": {"S":"quote"},":pk": {"S":"business#1cc284"}}
func (quoteDb QuoteDb) GetAllQuotesByBusinessId(ctx context.Context, businessId string) ([]*model.QuoteRequest, error) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#1cc284"},
			":sk": &types.AttributeValueMemberS{Value: "quote"},
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
	quotes := []*model.QuoteRequest{}
	for _, item := range res.Items {
		quote := model.QuoteRequest{}
		quoteData, ok := item["quotes"]
		if !ok {
			return nil, err
		}
		err := attributevalue.Unmarshal(quoteData, &quote)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, &quote)
	}
	if err != nil {
		return nil, err
	}
	return quotes, nil
}
