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

func (quoteDb QuoteDb) GetQuoteByQuoteId(ctx context.Context, quoteId string, businessId string) (*v1.QuoteRequest, error) {
	res, err := quoteDb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + "1cc284"},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + "2"},
		},
		ProjectionExpression: aws.String("quotes"),
	})
	if err != nil {
		return nil, err
	}

	quoteData := &v1.QuoteRequest{}
	data, ok := res.Item["quotes"]
	if !ok {
		return nil, errors.New("No quotes found for quoteId: " + quoteId + " and businessId: " + businessId)
	}
	err = attributevalue.Unmarshal(data, &quoteData)
	if err != nil {
		return nil, err
	}
	return quoteData, nil
}
