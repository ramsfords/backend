package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetShipperPrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		Name:      quoteReq.Pickup.Contact.Name,
		Phone:     quoteReq.Pickup.Contact.PhoneNumber,
		Email:     quoteReq.Pickup.Contact.EmailAddress,
		IsPrimary: false,
	}
}

func GetConsigneePrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		Name:      quoteReq.Delivery.Contact.Name,
		Phone:     quoteReq.Delivery.Contact.PhoneNumber,
		Email:     quoteReq.Delivery.Contact.EmailAddress,
		IsPrimary: false,
	}
}
