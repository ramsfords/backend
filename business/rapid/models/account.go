package models

type Account struct {
	Code string `json:"code" dynamodbav:"code"`
}
