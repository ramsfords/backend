package model

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// QuoteRequest is the request body for a quote request
type QuoteRequest struct {
	Bids         []*v1.Bid        `json:"bids" dynamodbav:"bids"`
	QuoteRequest v1.QuoteRequest  `json:"quoteRequest" dynamodbav:"quoteRequest"`
	SaveQuote    models.SaveQuote `json:"SaveQuote" dynamodbav:"SaveQuote"`
}

type BidsWithQuote struct {
	Bids         []*v1.Bid        `json:"bids" dynamodbav:"bids"`
	QuoteRequest *v1.QuoteRequest `json:"quoteRequest" dynamodbav:"quoteRequest"`
}
type BidWithQuote struct {
	Bid          *v1.Bid          `json:"bid" dynamodbav:"bid"`
	QuoteRequest *v1.QuoteRequest `json:"quoteRequest" dynamodbav:"quoteRequest"`
}
