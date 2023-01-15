package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetBillingAddress() models.Address {
	return models.Address{
		AddressID:            1039341,
		PickUpInstructions:   "",
		DeliveryInstructions: "",
		City:                 "Ontario",
		CompanyName:          "RamsFord inc",
		PrimaryContactPerson: &models.Contact{
			AddressContactID: 1082586,
			Name:             "Surendra Kandel",
			Email:            "kandelsuren@gmail.com",
			FirstName:        "Surendra",
			IsPrimary:        true,
			LastName:         "Kandel",
			Phone:            "7135162836",
			Position:         nil,
			Ext:              nil,
		},
		AddressContacts: []*models.Contact{
			{
				AddressContactID: 1082586,
				Name:             "Surendra Kandel",
				FirstName:        "Surendra",
				LastName:         "Kandel",
				Phone:            "7135162836",
				Ext:              nil,
				Email:            "kandelsuren@gmail.com",
				Position:         nil,
				IsPrimary:        true,
			},
		},
		Country:          "USA",
		CountryCode:      "US",
		DeliveryFromTime: "8:00:00 AM",
		DeliveryToTime:   "5:30:00 PM",
		IsCanada:         false,
		PostalCode:       "91762",
		ShippingFromTime: "8:00:00 AM",
		ShippingToTime:   "5:30:00 PM",
		State:            "California",
		StateCode:        "CA",
		StreetLine1:      "1131 West 6th Street",
		StreetLine2:      "",
		Location: &models.Location{
			AddressID: 1039341,
			Name:      "RamsFord inc",
			IsDefault: false,
		},
		CommercialType:      nil,
		AddressAccessorials: []models.AddressAccessorial{},
		Lat:                 34.0845839,
		Long:                -117.6719935,
		HasSaiaIntegration:  false,
	}
}
func GetOriginAddress(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote) models.Address {
	address := models.Address{}
	if quoteReq.Pickup.ShipperInstructions != "" {
		address.PickUpInstructions = quoteReq.Pickup.ShipperInstructions
	}
	if quoteReq.Pickup.Address.AddressLine_1 != "" {
		address.StreetLine1 = quoteReq.Pickup.Address.AddressLine_1
	}

	if quoteReq.Pickup.Address.AddressLine_2 != "" {
		address.StreetLine2 = quoteReq.Pickup.Address.AddressLine_2
	}

	if quoteReq.Pickup.Address.City != "" {
		address.City = quoteReq.Pickup.Address.City
	}

	if quoteReq.Pickup.Address.State != "" {
		address.State = quoteReq.Pickup.Address.State
	}
	if quoteReq.Pickup.Address.StateCode != "" {
		address.StateCode = quoteReq.Pickup.Address.StateCode
	}
	if quoteReq.Pickup.Address.Country != "" {
		address.Country = quoteReq.Pickup.Address.Country
	}
	if quoteReq.Pickup.Address.CountryCode != "" {
		address.CountryCode = quoteReq.Pickup.Address.CountryCode
	}
	if quoteReq.Pickup.Address.ZipCode != "" {
		address.PostalCode = quoteReq.Pickup.Address.ZipCode
	}
	if quoteReq.Pickup.Address.ZipCode != "" {
		address.Lat = float64(quoteReq.Pickup.Address.Lat)
	}
	if quoteReq.Pickup.Address.ZipCode != "" {
		address.Long = float64(quoteReq.Pickup.Address.Long)
	}
	address.PrimaryContactPerson = GetShipperPrimaryContact(quoteReq)
	pickupTimeWindow := NewPickupDateTimeWindow(quoteReq)
	address.ShippingFromTime = pickupTimeWindow.StartTime
	address.ShippingToTime = pickupTimeWindow.EndTime
	if saveQuote.QuoteDetails.OriginShippingDetails.Address.CommercialType != nil {
		address.CommercialType = saveQuote.QuoteDetails.OriginShippingDetails.Address.CommercialType
	}
	if saveQuote.QuoteDetails.OriginShippingDetails.Address.Location != nil {
		address.Location = saveQuote.QuoteDetails.OriginShippingDetails.Address.Location
	}

	if saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials != nil {
		address.AddressAccessorials = saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials
	}
	return address
}
func GetDeliveryAddress(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote) models.Address {
	address := models.Address{}
	if quoteReq.Delivery.ReceiverInstructions != "" {
		address.DeliveryInstructions = quoteReq.Delivery.ReceiverInstructions
	}
	if quoteReq.Delivery.Address.AddressLine_1 != "" {
		address.StreetLine1 = quoteReq.Delivery.Address.AddressLine_1
	}

	if quoteReq.Delivery.Address.AddressLine_2 != "" {
		address.StreetLine2 = quoteReq.Delivery.Address.AddressLine_2
	}

	if quoteReq.Delivery.Address.City != "" {
		address.City = quoteReq.Delivery.Address.City
	}

	if quoteReq.Delivery.Address.State != "" {
		address.State = quoteReq.Delivery.Address.State
	}
	if quoteReq.Delivery.Address.StateCode != "" {
		address.StateCode = quoteReq.Delivery.Address.StateCode
	}
	if quoteReq.Delivery.Address.Country != "" {
		address.Country = quoteReq.Delivery.Address.Country
	}
	if quoteReq.Delivery.Address.CountryCode != "" {
		address.CountryCode = quoteReq.Delivery.Address.CountryCode
	}
	if quoteReq.Delivery.Address.ZipCode != "" {
		address.PostalCode = quoteReq.Delivery.Address.ZipCode
	}
	if quoteReq.Delivery.Address.ZipCode != "" {
		address.Lat = float64(quoteReq.Delivery.Address.Lat)
	}
	if quoteReq.Delivery.Address.ZipCode != "" {
		address.Long = float64(quoteReq.Delivery.Address.Long)
	}
	address.PrimaryContactPerson = GetConsigneePrimaryContact(quoteReq)
	deliveryTimeWindow := NewDeliveryDateTimeWindow(quoteReq)
	address.ShippingFromTime = deliveryTimeWindow.StartTime
	address.ShippingToTime = deliveryTimeWindow.EndTime
	if saveQuote.QuoteDetails.DestinationShippingDetails.Address.CommercialType != nil {
		address.CommercialType = saveQuote.QuoteDetails.DestinationShippingDetails.Address.CommercialType
	}
	if saveQuote.QuoteDetails.DestinationShippingDetails.Address.Location != nil {
		address.Location = saveQuote.QuoteDetails.DestinationShippingDetails.Address.Location
	}

	if saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials != nil {
		address.AddressAccessorials = saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials
	}
	return address
}
