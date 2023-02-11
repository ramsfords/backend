package models

type HandelingUnitType struct {
	Code               string `json:"code,omitempty" dynamodbav:"code"`
	Description        string `json:"description,omitempty" dynamodbav:"description"`
	IsHazmatOnly       bool   `json:"isHazmatOnly,omitempty" dynamodbav:"isHazmatOnly"`
	IsHandlingUnitOnly bool   `json:"isHandlingUnitOnly,omitempty" dynamodbav:"isHandlingUnitOnly"`
}
