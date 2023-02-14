package models

type QuoteRate struct {
	IsValid                    bool                `json:"isValid" dynamodbav:"isValid"`
	TitleHeadData              TitleHeadData       `json:"titleHeadData" dynamodbav:"titleHeadData"`
	QuoteID                    int                 `json:"quoteId" dynamodbav:"quoteId"`
	ErrorMessages              []interface{}       `json:"errorMessages" dynamodbav:"errorMessages"`
	CarriersNotInQuote         []CarrierNotInQuote `json:"carriersNotInQuote" dynamodbav:"carriersNotInQuote"`
	DayDeliveries              []DayDelivery       `json:"dayDeliveries" dynamodbav:"dayDeliveries"`
	SelectedCarrier            interface{}         `json:"selectedCarrier" dynamodbav:"selectedCarrier"`
	SelectedCarrierTransitTime interface{}         `json:"selectedCarrierTransitTime" dynamodbav:"selectedCarrierTransitTime"`
}
