package models

type OriginShippingDetails struct {
	PickupDate            string  `json:"pickupDate,omitempty" dynamodbav:"pickupDate"`
	Address               Address `json:"address,omitempty" dynamodbav:"address"`
	RequestedDeliveryDate string  `json:"requestedDeliveryDate,omitempty" dynamodbav:"requestedDeliveryDate"`
}
