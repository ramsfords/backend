package quote_db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (quotedb QuoteDb) DeleteAllQuoteByBusinessId(ctx context.Context, buisnessId string) error {
	deleteItem := &dynamodb.DeleteItemInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "quote_count"},
			"sk": &types.AttributeValueMemberS{Value: "quote_count"},
		},
	}
	_, err := quotedb.Client.DeleteItem(ctx, deleteItem)
	if err != nil {
		return err
	}
	return nil
}
