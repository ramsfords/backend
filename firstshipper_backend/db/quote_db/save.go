package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quotedb QuoteDb) SaveQuote(ctx context.Context, qtReq v1.QuoteRequest) error {
	marshalledItem, err := attributevalue.MarshalMap(qtReq)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(quotedb.Config.GetFirstShipperTableName()),
		Item: map[string]types.AttributeValue{
			"pk":       &types.AttributeValueMemberS{Value: "pk" + qtReq.BusinessId},
			"sk":       &types.AttributeValueMemberS{Value: "quote#" + qtReq.QuoteId},
			"quote_pk": &types.AttributeValueMemberS{Value: "quote"},
			"quote_sk": &types.AttributeValueMemberS{Value: qtReq.QuoteId},
			"qyote#":   &types.AttributeValueMemberM{Value: marshalledItem},
		},
	}
	_, err = quotedb.Client.PutItem(ctx, input)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
