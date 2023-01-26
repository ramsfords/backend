package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetEmmergencyContact(quoteRes *v1.QuoteResponse) models.EmergencyContactPerson {
	contact := models.EmergencyContactPerson{}
	if quoteRes.QuoteRequest.Pickup.Contact.Name != "" {
		contact.Name += quoteRes.QuoteRequest.Pickup.Contact.Name
	}
	if quoteRes.QuoteRequest.Pickup.Contact.Name != "" {
		contact.Name += " " + quoteRes.QuoteRequest.Pickup.Contact.Name
	}
	if quoteRes.QuoteRequest.Pickup.Contact.PhoneNumber != "" {
		contact.Phone = quoteRes.QuoteRequest.Pickup.Contact.PhoneNumber
	}
	return contact
}
