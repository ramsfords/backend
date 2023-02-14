package models

type CommercialType struct {
	AccessorialID   int         `json:"accessorialId,omitempty"  dynamodbav:"accessorialId"`
	Name            string      `json:"name,omitempty"  dynamodbav:"name"`
	Code            interface{} `json:"code,omitempty"  dynamodbav:"code"`
	DestinationCode interface{} `json:"destinationCode,omitempty"  dynamodbav:"destinationCode"`
	IsOnlyForCanada bool        `json:"isOnlyForCanada,omitempty"  dynamodbav:"isOnlyForCanada"`
	IsOnlyForUSA    bool        `json:"isOnlyForUSA,omitempty"  dynamodbav:"isOnlyForUSA"`
	DataType        interface{} `json:"dataType,omitempty"  dynamodbav:"dataType"`
	Value           interface{} `json:"value,omitempty"  dynamodbav:"value"`
}
