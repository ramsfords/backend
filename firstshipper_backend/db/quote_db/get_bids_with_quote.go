package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
)

func (quoteDb QuoteDb) GetBidsWithQuoteByQuoteId(ctx context.Context, quoteId string) (*model.QuoteRequest, error) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		IndexName:              aws.String("quote_index"),
		KeyConditionExpression: aws.String("#quote_pk = :quote_pk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quote_pk": &types.AttributeValueMemberS{Value: quoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#quote_pk": "quote_pk",
		},
	})
	if err != nil {
		return nil, err
	}
	bidsRequest := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Items[0], bidsRequest)
	if err != nil {
		return nil, err
	}
	quoteWithBidsRes := &model.QuoteRequest{
		QuoteRequest: bidsRequest.QuoteRequest,
		Bids:         bidsRequest.Bids,
	}
	return quoteWithBidsRes, nil
}
