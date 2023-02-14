package models

type SaveQuoteResponse struct {
	IsSucceeded  bool   `json:"isSucceeded" dynamodbav:"isSucceeded"`
	SavedQuoteID string `json:"savedQuoteId" dynamodbav:"savedQuoteId"`
}
