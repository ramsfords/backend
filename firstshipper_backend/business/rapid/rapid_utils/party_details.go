package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/business/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetShippperPartyDetails(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote) models.PartyDetails {
	shippingAddress := quoteReq.Pickup.Address
	shipperParty := models.PartyDetails{}
	shipperParty.CompanyName = &quoteReq.Pickup.CompanyName
	shipperParty.StreetLine1 = &shippingAddress.AddressLine_1
	shipperParty.StreetLine2 = nil
	shipperParty.City = &shippingAddress.City
	shipperParty.State = &shippingAddress.State
	shipperParty.StateCode = &shippingAddress.StateCode
	if quoteReq.Pickup.Address.Country == "United States" {
		shipperParty.Country = utils.StrPtr("USA")
	}
	shipperParty.CountryCode = &shippingAddress.CountryCode
	shipperParty.PostalCode = &shippingAddress.ZipCode
	shipperParty.Lat = float64(shippingAddress.Lat)
	shipperParty.Long = float64(shippingAddress.Long)
	shipperParty.PrimaryContactPerson = GetShipperPrimaryContact(quoteReq)
	shipperParty.DateTimeWindow = NewPickupDateTimeWindow(quoteReq)
	shipperParty.Accessorials = saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials
	shipperParty.InstructionNote = &quoteReq.Pickup.ShipperInstructions
	shipperParty.SaveToAddressBook = true
	shipperParty.AddressID = 0
	shipperParty.IsFromAddressBook = false
	return shipperParty

}
func GetConsigneePartyDetails(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote) models.PartyDetails {
	receiverAddress := quoteReq.Delivery.Address
	receiverParty := models.PartyDetails{}
	receiverParty.CompanyName = &quoteReq.Delivery.CompanyName
	receiverParty.StreetLine1 = &receiverAddress.AddressLine_1
	receiverParty.StreetLine2 = nil
	receiverParty.City = &receiverAddress.City
	receiverParty.State = &receiverAddress.State
	receiverParty.StateCode = &receiverAddress.StateCode
	if quoteReq.Delivery.Address.Country == "United States" {
		receiverParty.Country = utils.StrPtr("USA")
	}
	receiverParty.Country = &receiverAddress.Country
	receiverParty.CountryCode = &receiverAddress.CountryCode
	receiverParty.PostalCode = &receiverAddress.ZipCode
	receiverParty.Lat = float64(receiverAddress.Lat)
	receiverParty.Long = float64(receiverAddress.Long)
	receiverParty.PrimaryContactPerson = GetConsigneePrimaryContact(quoteReq)
	receiverParty.DateTimeWindow = NewPickupDateTimeWindow(quoteReq)
	receiverParty.Accessorials = saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials
	receiverParty.InstructionNote = &quoteReq.Delivery.ReceiverInstructions
	receiverParty.SaveToAddressBook = true
	receiverParty.AddressID = 0
	receiverParty.IsFromAddressBook = false
	return receiverParty
}
