package model

import (
	"github.com/ramsfords/backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Bid interface {
	GetBid() interface{}
}
type BidWrapper struct {
	SaveQuote models.SaveQuote `json:"quote,omitempty"`
}

func (bidWrapper BidWrapper) GetBid() interface{} {
	return bidWrapper.SaveQuote
}

type Booking struct {
	Booking models.DispatchResponse `dynamodbav:"booking,omitempty" json:"booking,omitempty"`
	Quote   v1.QuoteRequest         `dynamodbav:"quote,omitempty" json:"quote,omitempty"`
}
