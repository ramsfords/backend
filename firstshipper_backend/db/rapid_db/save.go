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
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(rapiddb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + quoteReq.BusinessId},
			"sk": &types.AttributeValueMemberS{Value: "quotes#" + quoteReq.QuoteId},
		},
		ExpressionAttributeNames: map[string]string{
			"#rapid_quote": "rapidQuoteRate",
			"#quotes":      "quotes",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":rapid_quote": &types.AttributeValueMemberM{Value: marshalledItem},
		},
		UpdateExpression: aws.String("SET #quotes.#rapid_quote = :rapid_quote"),
	}
	_, err = rapiddb.Client.UpdateItem(ctx, input)
	return err
}
