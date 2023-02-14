package models

type ShipmentItem struct {
	HandlingUnitType   string      `json:"handlingUnitType,omitempty" dynamodbav:"handlingUnitType"`
	HandlingUnitNumber int         `json:"handlingUnitNumber,omitempty" dynamodbav:"handlingUnitNumber"`
	Dimensions         Dimensions  `json:"dimensions,omitempty" dynamodbav:"dimensions"`
	Stackable          bool        `json:"stackable" dynamodbav:"stackable"`
	Commodities        []Commodity `json:"commodities,omitempty" dynamodbav:"commodities"`
}
