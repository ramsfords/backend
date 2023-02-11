package models

type DestinationShippingDetails struct {
	Address               Address `json:"address,omitempty" dynamodbav:"address"`
	RequestedDeliveryDate string  `json:"requestedDeliveryDate,omitempty" dynamodbav:"requestedDeliveryDate"`
}
