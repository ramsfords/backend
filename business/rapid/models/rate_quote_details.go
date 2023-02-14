package models

type RateQuoteDetails struct {
	Amount           float64     `json:"amount" dynamodbav:"amount"`
	Rate             float64     `json:"rate" dynamodbav:"rate"`
	ItemFreightClass interface{} `json:"itemFreightClass" dynamodbav:"itemFreightClass"`
	Code             string      `json:"code" dynamodbav:"code"`
	Description      string      `json:"description" dynamodbav:"description"`
}
