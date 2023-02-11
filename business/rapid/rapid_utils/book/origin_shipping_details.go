package book

import (
	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/business/rapid/models"
)

func originShippingDetails(quoteRequest *model.QuoteRequest) error {
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.AddressContacts = []*models.Contact{
		{
			FirstName: getNameArray(quoteRequest.QuoteRequest.Pickup.Contact.Name)[0],
			LastName:  getNameArray(quoteRequest.QuoteRequest.Pickup.Contact.Name)[1],
			Phone:     quoteRequest.QuoteRequest.Pickup.Contact.PhoneNumber,
			Email:     quoteRequest.QuoteRequest.Pickup.Contact.EmailAddress,
			IsPrimary: true,
			Name:      quoteRequest.QuoteRequest.Pickup.Contact.Name,
		},
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.City = quoteRequest.QuoteRequest.Pickup.Address.City
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.CommercialType = &models.CommercialType{
		Name:            "Business",
		AccessorialID:   72,
		IsOnlyForCanada: false,
		IsOnlyForUSA:    false,
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.CompanyName = quoteRequest.QuoteRequest.Pickup.CompanyName
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.Country = "USA"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.CountryCode = "US"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.DeliveryFromTime = "10:00:00 AM"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.DeliveryToTime = "5:00:00 PM"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.HasSaiaIntegration = false
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.IsCanada = false
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.Lat = float64(quoteRequest.QuoteRequest.Pickup.Address.Lat)
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.Long = float64(quoteRequest.QuoteRequest.Pickup.Address.Long)
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.PickUpInstructions = quoteRequest.QuoteRequest.ShipperInstructions
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.PostalCode = quoteRequest.QuoteRequest.Pickup.Address.ZipCode
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.PrimaryContactPerson = &models.Contact{
		FirstName: getNameArray(quoteRequest.QuoteRequest.Pickup.Contact.Name)[0],
		LastName:  getNameArray(quoteRequest.QuoteRequest.Pickup.Contact.Name)[1],
		Phone:     quoteRequest.QuoteRequest.Pickup.Contact.PhoneNumber,
		Email:     quoteRequest.QuoteRequest.Pickup.Contact.EmailAddress,
		IsPrimary: true,
		Name:      quoteRequest.QuoteRequest.Pickup.Contact.Name,
	}
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.ShippingFromTime = "10:00:00 AM"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.ShippingToTime = "5:00:00 PM"
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.State = quoteRequest.QuoteRequest.Pickup.Address.State
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.StateCode = quoteRequest.QuoteRequest.Pickup.Address.State
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.StreetLine1 = quoteRequest.QuoteRequest.Pickup.Address.AddressLine1
	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.StreetLine2 = quoteRequest.QuoteRequest.Pickup.Address.AddressLine2

	quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address = GetOriginAddress(quoteRequest)
	if quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials == nil {
		quoteRequest.RapidSaveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials = []models.AddressAccessorial{}
	}
	return nil
}
