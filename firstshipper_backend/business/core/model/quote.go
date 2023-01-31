package model

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// QuoteRequest is the request body for a quote request
type QuoteRequest struct {
	Bids              []*v1.Bid                 `json:"bids" dynamodbav:"bids"`
	Bid               *v1.Bid                   `json:"bid" dynamodbav:"bid"`
	QuoteRequest      *v1.QuoteRequest          `json:"quoteRequest" dynamodbav:"quoteRequest"`
	RapidSaveQuote    *models.SaveQuote         `json:"SaveQuote" dynamodbav:"rapidSaveQuote"`
	SaveQuoteResponse *models.SaveQuoteResponse `json:"saveQuoteResponse" dynamodbav:"saveQuoteResponse"`
	RapidBooking      *models.DispatchResponse  `json:"Booking" dynamodbav:"rapidBooking"`
	BookingInfo       *v1.BookingInfo           `json:"bookingInfo" dynamodbav:"bookingInfo"`
}

type BidsWithQuote struct {
	Bids         []*v1.Bid        `json:"bids" dynamodbav:"bids"`
	QuoteRequest *v1.QuoteRequest `json:"quoteRequest" dynamodbav:"quoteRequest"`
}
type BidWithQuote struct {
	Bid          *v1.Bid          `json:"bid" dynamodbav:"bid"`
	QuoteRequest *v1.QuoteRequest `json:"quoteRequest" dynamodbav:"quoteRequest"`
}
