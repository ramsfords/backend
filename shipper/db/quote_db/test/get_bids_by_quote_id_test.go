package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/shipper/business/core/model"
)

func TestBidsByQuoteID(t *testing.T) {
	businessID := "1cc284"
	quoteID := "scoppVzQRL"
	res, err := quoteDb.Client.Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(quoteDb.GetFirstShipperTableName()),
		KeyConditionExpression: aws.String("#pk = :pk And begins_with(#sk, :sk)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "pk#" + businessID},
			":sk": &types.AttributeValueMemberS{Value: "quote#" + quoteID},
		},
		ExpressionAttributeNames: map[string]string{
			"#sk":     "sk",
			"#pk":     "pk",
			"#quotes": "quotes",
		},
		ProjectionExpression: aws.String("#quotes"),
	})
	if err != nil {
		t.Log(err)
	}
	quote, ok := res.Items[0]["quotes"]
	if ok {
		quoteRes := model.QuoteRequest{}
		err := attributevalue.Unmarshal(quote, &quoteRes)
		if err != nil {
			t.Log(err)
		}
		fmt.Println(quoteRes)
	}
	// for _, items := range bids {
	// 	newItem, ok := quote["bids"]
	// 	if ok {
	// 		bids := []*v1.Bids{}
	// 		err := attributevalue.Unmarshal(bids, &bids)
	// 		if err != nil {
	// 			t.Log(err)
	// 		}
	// 		fmt.Println(bids)
	// 	}
	// }
	// if ok {
	// 	bids, ok := quote(map[string]types.AttributeValue)["bids"]
	// 	if ok {
	// 		bids := []*v1.Bids{}
	// 		err := attributevalue.Unmarshal(bids, &bids)
	// 		if err != nil {
	// 			t.Log(err)
	// 		}
	// 		fmt.Println(bids)
	// 	}
	// 	fmt.Println(bids)
	// }

	fmt.Println(res, quote, ok)
}
