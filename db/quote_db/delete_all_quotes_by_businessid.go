package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quoteDb QuoteDb) DeleteAllQuoteByBusinessId(ctx context.Context, buisnessId string) error {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + buisnessId},
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
		return err
	}
	for _, item := range res.Items {
		quote := v1.QuoteRequest{}
		quoteData, ok := item["quotes"]
		if !ok {
			return err
		}
		err := attributevalue.Unmarshal(quoteData, &quote)
		if err != nil {
			return err
		}
		_, err = quoteDb.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
			TableName: aws.String(quoteDb.GetFirstShipperTableName()),
			Key: map[string]types.AttributeValue{
				"pk": &types.AttributeValueMemberS{Value: "pk#" + quote.BusinessId},
				"sk": &types.AttributeValueMemberS{Value: "quote#" + quote.QuoteId},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
