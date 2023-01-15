package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetEmmergencyContact(quoteReq *v1.QuoteRequest) models.EmergencyContactPerson {
	contact := models.EmergencyContactPerson{}
	if quoteReq.Pickup.Contact.FirstName != "" {
		contact.Name += quoteReq.Pickup.Contact.FirstName
	}
	if quoteReq.Pickup.Contact.LastName != "" {
		contact.Name += " " + quoteReq.Pickup.Contact.LastName
	}
	if quoteReq.Pickup.Contact.PhoneNumber != "" {
		contact.Phone = quoteReq.Pickup.Contact.PhoneNumber
	}
	return contact
}
