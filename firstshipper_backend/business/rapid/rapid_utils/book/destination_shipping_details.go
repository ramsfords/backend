package book

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
)

func destinationShippingDetails(quoteRequest *model.QuoteRequest) error {
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressContacts = []*models.Contact{
		{
			FirstName: getNameArray(quoteRequest.QuoteRequest.Delivery.Contact.Name)[0],
			LastName:  getNameArray(quoteRequest.QuoteRequest.Delivery.Contact.Name)[1],
			Phone:     quoteRequest.QuoteRequest.Delivery.Contact.PhoneNumber,
			Email:     quoteRequest.QuoteRequest.Delivery.Contact.EmailAddress,
			IsPrimary: true,
			Name:      quoteRequest.QuoteRequest.Delivery.Contact.Name,
		},
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.City = quoteRequest.QuoteRequest.Delivery.Address.City
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.CommercialType = &models.CommercialType{
		Name:            "Business",
		AccessorialID:   72,
		IsOnlyForCanada: false,
		IsOnlyForUSA:    false,
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.CompanyName = quoteRequest.QuoteRequest.Delivery.CompanyName
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.Country = "USA"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.CountryCode = "US"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.DeliveryFromTime = "10:00:00 AM"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.DeliveryToTime = "5:00:00 PM"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.HasSaiaIntegration = false
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.IsCanada = false
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.Lat = float64(quoteRequest.QuoteRequest.Delivery.Address.Lat)
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.Long = float64(quoteRequest.QuoteRequest.Delivery.Address.Long)
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.PickUpInstructions = quoteRequest.QuoteRequest.ShipperInstructions
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.PostalCode = quoteRequest.QuoteRequest.Delivery.Address.ZipCode
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.PrimaryContactPerson = &models.Contact{
		FirstName: getNameArray(quoteRequest.QuoteRequest.Delivery.Contact.Name)[0],
		LastName:  getNameArray(quoteRequest.QuoteRequest.Delivery.Contact.Name)[1],
		Phone:     quoteRequest.QuoteRequest.Delivery.Contact.PhoneNumber,
		Email:     quoteRequest.QuoteRequest.Delivery.Contact.EmailAddress,
		IsPrimary: true,
		Name:      quoteRequest.QuoteRequest.Delivery.Contact.Name,
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.ShippingFromTime = "10:00:00 AM"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.ShippingToTime = "5:00:00 PM"
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.State = quoteRequest.QuoteRequest.Delivery.Address.State
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.StateCode = quoteRequest.QuoteRequest.Delivery.Address.State
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.StreetLine1 = quoteRequest.QuoteRequest.Delivery.Address.AddressLine1
	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.StreetLine2 = quoteRequest.QuoteRequest.Delivery.Address.AddressLine2

	quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address = GetOriginAddress(quoteRequest)
	if quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials == nil {
		quoteRequest.RapidSaveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials = []models.AddressAccessorial{}
	}
	return nil
}
