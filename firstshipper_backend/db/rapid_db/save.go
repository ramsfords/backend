package rapid_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (rapiddb RapidDb) SaveRapidQuote(ctx context.Context, quote models.QuoteRate, quoteReq v1.QuoteRequest) error {
	marshalledItem, err := attributevalue.MarshalMap(quote)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(rapiddb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":             &types.AttributeValueMemberS{Value: "pk#" + quoteReq.BusinessId},
			"sk":             &types.AttributeValueMemberS{Value: "rapid_quote_sk#" + quoteReq.BusinessId},
			"rapid_quote_pk": &types.AttributeValueMemberS{Value: "rapid_quote"},
			"rapid_quote_sk": &types.AttributeValueMemberS{Value: quoteReq.BusinessId},
			"rapid_quote":    &types.AttributeValueMemberM{Value: marshalledItem},
		},
	}
	_, err = rapiddb.Client.PutItem(ctx, input)
	return err
}
