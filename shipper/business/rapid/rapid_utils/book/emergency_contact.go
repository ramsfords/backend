package book

import (
	"github.com/ramsfords/backend/shipper/business/core/model"
	"github.com/ramsfords/backend/shipper/business/rapid/models"
)

func getEmmergencyContact(quoteReq *model.QuoteRequest) models.EmergencyContactPerson {
	return models.EmergencyContactPerson{
		Name:  quoteReq.QuoteRequest.Pickup.Contact.Name,
		Phone: quoteReq.QuoteRequest.Pickup.Contact.PhoneNumber,
	}
}
