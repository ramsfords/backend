package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quotedb QuoteDb) UpdateQuote(ctx context.Context, quoteReq *v1.QuoteRequest) error {
	marshalledItem, err := attributevalue.MarshalMap(quoteReq)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk":       &types.AttributeValueMemberS{Value: "business#" + quoteReq.BusinessId},
			"sk":       &types.AttributeValueMemberS{Value: "quote#" + quoteReq.QuoteId},
			"quote_pk": &types.AttributeValueMemberS{Value: "quote"},
			"quote_sk": &types.AttributeValueMemberS{Value: quoteReq.QuoteId},
			"quote":    &types.AttributeValueMemberM{Value: marshalledItem},
		},
		UpdateExpression: aws.String("SET #quote = :quote"),
		ExpressionAttributeNames: map[string]string{
			"#quote": "quote",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":quote": &types.AttributeValueMemberM{Value: marshalledItem},
		},
	}
	_, err = quotedb.Client.UpdateItem(ctx, input)
	return err
}
