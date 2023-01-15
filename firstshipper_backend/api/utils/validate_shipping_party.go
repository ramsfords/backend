package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateShippingParty(loc *v1.Location) *errs.ApiErr {
	if loc.Contact.FirstName == "" || len(loc.Contact.FirstName) < 2 || loc.Contact.LastName == "" {
		return &errs.InvalidContactInfo
	}
	if (loc.CompanyName == "" || len(loc.CompanyName) < 5) && loc.Contact.FirstName == "" {
		return &errs.InvalidCompanyName
	}
	if loc.Contact.EmailAddress == "" || len(loc.Contact.EmailAddress) < 5 {
		return &errs.InvalidEmailAddress
	}
	if loc.Contact.PhoneNumber == "" || len(loc.Contact.PhoneNumber) < 10 || len(loc.Contact.PhoneNumber) > 12 {
		return &errs.InvalidPhoneNumber
	}
	return nil
}
