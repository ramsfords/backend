package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/business/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetShippperPartyDetails(quoteRes *v1.QuoteResponse, saveQuote *models.SaveQuote) models.PartyDetails {
	shippingAddress := quoteRes.QuoteRequest.Pickup.Address
	shipperParty := models.PartyDetails{}
	shipperParty.CompanyName = &quoteRes.QuoteRequest.Pickup.CompanyName
	shipperParty.StreetLine1 = &shippingAddress.AddressLine1
	shipperParty.StreetLine2 = nil
	shipperParty.City = &shippingAddress.City
	shipperParty.State = &shippingAddress.State
	shipperParty.StateCode = &shippingAddress.StateCode
	if quoteRes.QuoteRequest.Pickup.Address.Country == "United States" {
		shipperParty.Country = utils.StrPtr("USA")
	}
	shipperParty.CountryCode = &shippingAddress.CountryCode
	shipperParty.PostalCode = &shippingAddress.ZipCode
	shipperParty.Lat = float64(shippingAddress.Lat)
	shipperParty.Long = float64(shippingAddress.Long)
	shipperParty.PrimaryContactPerson = GetShipperPrimaryContact(quoteRes.QuoteRequest)
	shipperParty.DateTimeWindow = NewPickupDateTimeWindow(quoteRes.QuoteRequest)
	shipperParty.Accessorials = saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials
	shipperParty.InstructionNote = &quoteRes.QuoteRequest.ShipperInstructions
	shipperParty.SaveToAddressBook = true
	shipperParty.AddressID = 0
	shipperParty.IsFromAddressBook = false
	return shipperParty

}
func GetConsigneePartyDetails(quoteRes *v1.QuoteResponse, saveQuote *models.SaveQuote) models.PartyDetails {
	receiverAddress := quoteRes.QuoteRequest.Delivery.Address
	receiverParty := models.PartyDetails{}
	receiverParty.CompanyName = &quoteRes.QuoteRequest.Delivery.CompanyName
	receiverParty.StreetLine1 = &receiverAddress.AddressLine1
	receiverParty.StreetLine2 = nil
	receiverParty.City = &receiverAddress.City
	receiverParty.State = &receiverAddress.State
	receiverParty.StateCode = &receiverAddress.StateCode
	if quoteRes.QuoteRequest.Delivery.Address.Country == "United States" {
		receiverParty.Country = utils.StrPtr("USA")
	}
	receiverParty.Country = &receiverAddress.Country
	receiverParty.CountryCode = &receiverAddress.CountryCode
	receiverParty.PostalCode = &receiverAddress.ZipCode
	receiverParty.Lat = float64(receiverAddress.Lat)
	receiverParty.Long = float64(receiverAddress.Long)
	receiverParty.PrimaryContactPerson = GetConsigneePrimaryContact(quoteRes.QuoteRequest)
	receiverParty.DateTimeWindow = NewPickupDateTimeWindow(quoteRes.QuoteRequest)
	receiverParty.Accessorials = saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials
	receiverParty.InstructionNote = &quoteRes.QuoteRequest.ReceiverInstructions
	receiverParty.SaveToAddressBook = true
	receiverParty.AddressID = 0
	receiverParty.IsFromAddressBook = false
	return receiverParty
}
