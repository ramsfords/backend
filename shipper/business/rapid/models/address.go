package models

type Address struct {
	AddressID             int                  `json:"addressId"  dynamodbav:"addressId"`
	StreetLine1           string               `json:"streetLine1"  dynamodbav:"streetLine1"`
	StreetLine2           string               `json:"streetLine2"  dynamodbav:"streetLine2"`
	State                 string               `json:"state"  dynamodbav:"state"`
	StateCode             string               `json:"stateCode"  dynamodbav:"stateCode"`
	IsCanada              bool                 `json:"isCanada"  dynamodbav:"isCanada"`
	City                  string               `json:"city"  dynamodbav:"city"`
	Country               string               `json:"country"  dynamodbav:"country"`
	CountryCode           string               `json:"countryCode"  dynamodbav:"countryCode"`
	PostalCode            string               `json:"postalCode"  dynamodbav:"postalCode"`
	Lat                   float64              `json:"lat"  dynamodbav:"lat"`
	Long                  float64              `json:"long"  dynamodbav:"long"`
	Location              *Location            `json:"location"  dynamodbav:"location"`
	PickUpInstructions    string               `json:"pickUpInstructions"  dynamodbav:"pickUpInstructions"`
	DeliveryInstructions  string               `json:"deliveryInstructions"  dynamodbav:"deliveryInstructions"`
	CompanyName           string               `json:"companyName"  dynamodbav:"companyName"`
	PrimaryContactPerson  *Contact             `json:"primaryContactPerson"  dynamodbav:"primaryContactPerson"`
	AddressContacts       []*Contact           `json:"addressContacts"  dynamodbav:"addressContacts"`
	DeliveryFromTime      string               `json:"deliveryFromTime"  dynamodbav:"deliveryFromTime"`
	DeliveryToTime        string               `json:"deliveryToTime"  dynamodbav:"deliveryToTime"`
	ShippingFromTime      string               `json:"shippingFromTime"  dynamodbav:"shippingFromTime"`
	ShippingToTime        string               `json:"shippingToTime"  dynamodbav:"shippingToTime"`
	CommercialType        *CommercialType      `json:"commercialType"  dynamodbav:"commercialType"`
	AddressAccessorials   []AddressAccessorial `json:"addressAccessorials"  dynamodbav:"addressAccessorials"`
	BulkAddressID         string               `json:"bulkAddressId"  dynamodbav:"bulkAddressId"`
	BulkEditGUID          string               `json:"bulkEditGuid"  dynamodbav:"bulkEditGuid"`
	HasSaiaIntegration    bool                 `json:"hasSaiaIntegration"  dynamodbav:"hasSaiaIntegration"`
	PickupDate            string               `json:"pickupDate,omitempty"`
	RequestedDeliveryDate string               `json:"requestedDeliveryDate,omitempty"`
}
