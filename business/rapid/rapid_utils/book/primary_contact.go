package book

import (
	"strings"

	"github.com/ramsfords/backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func getShipperPrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		Name:      quoteReq.Pickup.Contact.Name,
		Phone:     quoteReq.Pickup.Contact.PhoneNumber,
		Email:     "kandelsuren@gmail.com",
		FirstName: getNameArray(quoteReq.Pickup.Contact.Name)[0],
		LastName:  getNameArray(quoteReq.Pickup.Contact.Name)[1],
		IsPrimary: true,
	}
}

func getConsigneePrimaryContact(quoteReq *v1.QuoteRequest) *models.Contact {
	return &models.Contact{
		Name:      quoteReq.Delivery.Contact.Name,
		Phone:     quoteReq.Delivery.Contact.PhoneNumber,
		Email:     quoteReq.Delivery.Contact.EmailAddress,
		FirstName: getNameArray(quoteReq.Pickup.Contact.Name)[0],
		LastName:  getNameArray(quoteReq.Pickup.Contact.Name)[1],
		IsPrimary: false,
	}
}
func getNameArray(name string) []string {
	return strings.Split(name, " ")
}
