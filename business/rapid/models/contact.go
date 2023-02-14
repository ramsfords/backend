package models

type Contact struct {
	AddressContactID int         `json:"addressContactId" dynamodbav:"addressContactId"`
	Name             string      `json:"name" dynamodbav:"name"`
	Email            string      `json:"email" dynamodbav:"email"`
	FirstName        string      `json:"firstName" dynamodbav:"firstName"`
	IsPrimary        bool        `json:"isPrimary" dynamodbav:"isPrimary"`
	LastName         string      `json:"lastName" dynamodbav:"lastName"`
	Phone            string      `json:"phone" dynamodbav:"phone"`
	Position         interface{} `json:"position" dynamodbav:"position"`
	Ext              interface{} `json:"ext" dynamodbav:"ext"`
}
