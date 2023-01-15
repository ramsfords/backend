package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func UpdateQuoteDetails(quoteReq *v1.QuoteRequest, saveQuote *models.SaveQuote) *models.SaveQuote {
	if saveQuote.QuoteDetails == nil {
		saveQuote.QuoteDetails = &models.QuoteDetails{}
	}
	saveQuote.QuoteDetails.BillingAddress = GetBillingAddress()
	saveQuote.QuoteDetails.OriginShippingDetails.Address = models.Address{}
	saveQuote.QuoteDetails.OriginShippingDetails.Address = GetOriginAddress(quoteReq, saveQuote)
	if saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials == nil {
		saveQuote.QuoteDetails.OriginShippingDetails.Address.AddressAccessorials = []models.AddressAccessorial{}
	}
	saveQuote.QuoteDetails.DestinationShippingDetails.Address = models.Address{}
	saveQuote.QuoteDetails.DestinationShippingDetails.Address = GetDeliveryAddress(quoteReq, saveQuote)
	if saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials == nil {
		saveQuote.QuoteDetails.DestinationShippingDetails.Address.AddressAccessorials = []models.AddressAccessorial{}
	}
	emergencyContact := GetEmmergencyContact(quoteReq)
	saveQuote.QuoteDetails.EmergencyContactPerson = &emergencyContact
	return saveQuote
}
