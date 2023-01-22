package test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	v1 "github.com/ramsfords/types_gen/v1"
)

// "TableName": "firstshipper-dev",
// "KeyConditionExpression": "#fcc61 = :fcc61 And begins_with(#fcc62, :fcc62)",
// "ProjectionExpression": "#fcc60",
// "ExpressionAttributeNames": {"#fcc60":"quotes","#fcc61":"pk","#fcc62":"sk"},
// "ExpressionAttributeValues": {":fcc61": {"S":"pk#1cc284"},":fcc62": {"S":"quote"}}
func TestGetQuotesByBusinessId(t *testing.T) {
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#1cc284"},
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
		t.Fatal(err)
	}
	quotes := []v1.QuoteRequest{}
	for _, item := range res.Items {
		quote := v1.QuoteRequest{}
		quoteData, ok := item["quotes"]
		if !ok {
			t.Fatal("no quotes")
		}
		err := attributevalue.Unmarshal(quoteData, &quote)
		if err != nil {
			t.Fatal(err)
		}
		quotes = append(quotes, quote)
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
