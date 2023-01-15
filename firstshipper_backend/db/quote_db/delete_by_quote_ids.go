package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (quotedb QuoteDb) DeleteQuotesByQuoteIds(ctx context.Context, businessId string, quoteId []string) error {
	items := []types.TransactWriteItem{}
	for _, j := range quoteId {
		items = append(items, types.TransactWriteItem{
			Delete: &types.Delete{
				TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
				Key: map[string]types.AttributeValue{
					"pk": &types.AttributeValueMemberS{Value: "business#" + businessId},
					"sk": &types.AttributeValueMemberS{Value: "quote#" + j},
				},
			},
		})
	}
	_, err := quotedb.Client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: items,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
