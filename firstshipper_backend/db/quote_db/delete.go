package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (quotedb QuoteDb) DeleteQuote(ctx context.Context, quoteId string) error {
	_, err := quotedb.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName:           aws.String(quotedb.Config.GetFirstShipperTableName()),
		ConditionExpression: aws.String(fmt.Sprintf("attribute_exists(%s)", "sk")),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "business#" + quoteId},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
		},
	})
	return err
}
