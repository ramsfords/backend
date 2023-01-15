package quote_db

import (
	"context"

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
func (quotedb QuoteDb) GetAllQuotesByBusinessId(ctx context.Context, businessId string) ([]*v1.QuoteRequest, error) {
	input := &dynamodb.QueryInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		IndexName: aws.String("quote_index"),
		ExpressionAttributeNames: map[string]string{
			"#quote_pk": "quote_pk",
			"#pk":       "pk",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quote_pk": &types.AttributeValueMemberS{Value: "quote"},
			":pk":       &types.AttributeValueMemberS{Value: "business#1cc284"},
		},
		KeyConditionExpression: aws.String("#quote_pk = :quote_pk"),
		FilterExpression:       aws.String("#pk = :pk"),
	}
	out, err := quotedb.Client.Query(ctx, input)
	if err != nil {
		return nil, err
	}
	quotes := []*v1.QuoteRequest{}
	err = attributevalue.UnmarshalListOfMaps(out.Items, &quotes)
	if err != nil {
		return nil, err
	}
	return quotes, nil
}
