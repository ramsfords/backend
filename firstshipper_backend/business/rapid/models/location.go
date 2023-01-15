package models

type Location struct {
	AddressID int    `json:"addressId,omitempty" dynamodbav:"addressId"`
	Name      string `json:"name,omitempty" dynamodbav:"name"`
	IsDefault bool   `json:"isDefault" dynamodbav:"isDefault"`
}
