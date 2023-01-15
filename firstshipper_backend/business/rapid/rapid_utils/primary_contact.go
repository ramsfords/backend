package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetShipperPrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		LastName:  quoteReq.Pickup.Contact.LastName,
		FirstName: quoteReq.Pickup.Contact.FirstName,
		Name:      quoteReq.Pickup.Contact.FirstName + " " + quoteReq.Pickup.Contact.LastName,
		Phone:     quoteReq.Pickup.Contact.PhoneNumber,
		Email:     quoteReq.Pickup.Contact.EmailAddress,
		IsPrimary: false,
	}
}

func GetConsigneePrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		LastName:  quoteReq.Delivery.Contact.LastName,
		FirstName: quoteReq.Delivery.Contact.FirstName,
		Name:      quoteReq.Delivery.Contact.FirstName + " " + quoteReq.Delivery.Contact.LastName,
		Phone:     quoteReq.Delivery.Contact.PhoneNumber,
		Email:     quoteReq.Delivery.Contact.EmailAddress,
		IsPrimary: false,
	}
}
