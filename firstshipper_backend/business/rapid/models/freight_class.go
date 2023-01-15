package models

type FreightClass struct {
	Code            string `json:"code,omitempty" dynamodbav:"code"`
	Description     string `json:"description,omitempty" dynamodbav:"description"`
	PFortyFourClass string `json:"pFortyFourClass,omitempty" dynamodbav:"pFortyFourClass"`
}
