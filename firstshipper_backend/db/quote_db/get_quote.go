package quote_db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
)

// type QuoteRequest struct {
// 	Bids              map[string]v1.Bid         `json:"bids" dynamodbav:"bids"`
// 	QuoteRequest      *v1.QuoteRequest          `json:"quoteRequest" dynamodbav:"quoteRequest"`
// 	RapidSaveQuote    *models.SaveQuote         `json:"SaveQuote" dynamodbav:"rapidSaveQuote"`
// 	SaveQuoteResponse *models.SaveQuoteResponse `json:"saveQuoteResponse" dynamodbav:"saveQuoteResponse"`
// 	RapidBooking      *models.DispatchResponse  `json:"Booking" dynamodbav:"rapidBooking"`
// }

func (quoteDb QuoteDb) GetQuoteByQuoteId(ctx context.Context, quoteId string, businessId string) (*model.QuoteRequest, error) {
	res, err := quoteDb.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String(quoteDb.GetFirstShipperTableName()),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: "pk#" + businessId},
			"sk": &types.AttributeValueMemberS{Value: "quote#" + quoteId},
		},
	})
	if err != nil {
		return nil, err
	}

	quoteData := &model.QuoteRequest{}
	err = attributevalue.UnmarshalMap(res.Item, quoteData)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return quoteData, nil
}
