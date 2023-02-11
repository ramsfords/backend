package book

import (
	"github.com/ramsfords/backend/shipper/business/core/model"
	"github.com/ramsfords/backend/shipper/business/rapid/models"
)

func GetOriginAddress(saveQuote *model.QuoteRequest) models.Address {
	address := models.Address{}
	if saveQuote.QuoteRequest.ShipperInstructions != "" {
		address.PickUpInstructions = saveQuote.QuoteRequest.ShipperInstructions
	}
	if saveQuote.QuoteRequest.Pickup.Address.AddressLine1 != "" {
		address.StreetLine1 = saveQuote.QuoteRequest.Pickup.Address.AddressLine1
	}

	if saveQuote.QuoteRequest.Pickup.Address.AddressLine2 != "" {
		address.StreetLine2 = saveQuote.QuoteRequest.Pickup.Address.AddressLine2
	}

	if saveQuote.QuoteRequest.Pickup.Address.City != "" {
		address.City = saveQuote.QuoteRequest.Pickup.Address.City
	}

	if saveQuote.QuoteRequest.Pickup.Address.State != "" {
		address.State = saveQuote.QuoteRequest.Pickup.Address.State
	}
	if saveQuote.QuoteRequest.Pickup.Address.StateCode != "" {
		address.StateCode = saveQuote.QuoteRequest.Pickup.Address.StateCode
	}
	if saveQuote.QuoteRequest.Pickup.Address.Country != "" {
		address.Country = saveQuote.QuoteRequest.Pickup.Address.Country
	}
	if saveQuote.QuoteRequest.Pickup.Address.CountryCode != "" {
		address.CountryCode = saveQuote.QuoteRequest.Pickup.Address.CountryCode
	}
	if saveQuote.QuoteRequest.Pickup.Address.ZipCode != "" {
		address.PostalCode = saveQuote.QuoteRequest.Pickup.Address.ZipCode
	}
	if saveQuote.QuoteRequest.Pickup.Address.ZipCode != "" {
		address.Lat = float64(saveQuote.QuoteRequest.Pickup.Address.Lat)
	}
	if saveQuote.QuoteRequest.Pickup.Address.ZipCode != "" {
		address.Long = float64(saveQuote.QuoteRequest.Pickup.Address.Long)
	}
	address.PrimaryContactPerson = getShipperPrimaryContact(saveQuote.QuoteRequest)
	pickupTimeWindow := pickupDateTimeWindow(saveQuote)
	address.ShippingFromTime = pickupTimeWindow.StartTime
	address.ShippingToTime = pickupTimeWindow.EndTime
	if saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.CommercialType != nil {
		address.CommercialType = saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.CommercialType
	}
	if saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.Location != nil {
		address.Location = saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.Location
	}

	if saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials != nil {
		address.AddressAccessorials = saveQuote.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials
	}
	return address
}
func GetDeliveryAddress(saveQuote *model.QuoteRequest) models.Address {
	address := models.Address{}
	if saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.DeliveryInstructions != "" {
		address.DeliveryInstructions = saveQuote.QuoteRequest.ReceiverInstructions
	}
	if saveQuote.QuoteRequest.Delivery.Address.AddressLine1 != "" {
		address.StreetLine1 = saveQuote.QuoteRequest.Delivery.Address.AddressLine1
	}

	if saveQuote.QuoteRequest.Delivery.Address.AddressLine2 != "" {
		address.StreetLine2 = saveQuote.QuoteRequest.Delivery.Address.AddressLine2
	}

	if saveQuote.QuoteRequest.Delivery.Address.City != "" {
		address.City = saveQuote.QuoteRequest.Delivery.Address.City
	}

	if saveQuote.QuoteRequest.Delivery.Address.State != "" {
		address.State = saveQuote.QuoteRequest.Delivery.Address.State
	}
	if saveQuote.QuoteRequest.Delivery.Address.StateCode != "" {
		address.StateCode = saveQuote.QuoteRequest.Delivery.Address.StateCode
	}
	if saveQuote.QuoteRequest.Delivery.Address.Country != "" {
		address.Country = saveQuote.QuoteRequest.Delivery.Address.Country
	}
	if saveQuote.QuoteRequest.Delivery.Address.CountryCode != "" {
		address.CountryCode = saveQuote.QuoteRequest.Delivery.Address.CountryCode
	}
	if saveQuote.QuoteRequest.Delivery.Address.ZipCode != "" {
		address.PostalCode = saveQuote.QuoteRequest.Delivery.Address.ZipCode
	}
	if saveQuote.QuoteRequest.Delivery.Address.ZipCode != "" {
		address.Lat = float64(saveQuote.QuoteRequest.Delivery.Address.Lat)
	}
	if saveQuote.QuoteRequest.Delivery.Address.ZipCode != "" {
		address.Long = float64(saveQuote.QuoteRequest.Delivery.Address.Long)
	}
	address.PrimaryContactPerson = getConsigneePrimaryContact(saveQuote.QuoteRequest)
	deliveryTimeWindow := deliveryDateTimeWindow()
	address.ShippingFromTime = deliveryTimeWindow.StartTime
	address.ShippingToTime = deliveryTimeWindow.EndTime
	if saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.CommercialType != nil {
		address.CommercialType = saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.CommercialType
	}
	if saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.Location != nil {
		address.Location = saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.Location
	}

	if saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials != nil {
		address.AddressAccessorials = saveQuote.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials
	}
	return address
}
