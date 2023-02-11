package models

type PartyDetails struct {
	CompanyName           *string               `json:"companyName" dynamodbav:"companyName"`
	StreetLine1           *string               `json:"streetLine1" dynamodbav:"streetLine1"`
	StreetLine2           *string               `json:"streetLine2" dynamodbav:"streetLine2"`
	City                  *string               `json:"city" dynamodbav:"city"`
	State                 *string               `json:"state" dynamodbav:"state"`
	StateCode             *string               `json:"stateCode" dynamodbav:"stateCode"`
	PostalCode            *string               `json:"postalCode" dynamodbav:"postalCode"`
	Country               *string               `json:"country" dynamodbav:"country"`
	CountryCode           *string               `json:"countryCode" dynamodbav:"countryCode"`
	PrimaryContactPerson  *Contact              `json:"primaryContactPerson" dynamodbav:"primaryContactPerson"`
	DateTimeWindow        DateTimeWindow        `json:"dateTimeWindow" dynamodbav:"dateTimeWindow"`
	IsFromAddressBook     bool                  `json:"isFromAddressBook" dynamodbav:"isFromAddressBook"`
	IsCanada              bool                  `json:"isCanada" dynamodbav:"isCanada"`
	Accessorials          []AddressAccessorial  `json:"accessorials" dynamodbav:"accessorials"`
	InstructionNote       *string               `json:"instructionNote" dynamodbav:"instructionNote"`
	SaveToAddressBook     bool                  `json:"saveToAddressBook" dynamodbav:"saveToAddressBook"`
	Lat                   float64               `json:"lat" dynamodbav:"lat"`
	Long                  float64               `json:"long" dynamodbav:"long"`
	UpdateDefaultsAddress UpdateDefaultsAddress `json:"updateDefaultsAddress" dynamodbav:"updateDefaultsAddress"`
	AddressID             int                   `json:"addressId" dynamodbav:"addressId"`
}
