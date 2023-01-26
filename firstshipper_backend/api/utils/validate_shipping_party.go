package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateShippingParty(loc *v1.Location) *errs.ApiErr {
	if len(loc.Contact.Name) < 2 {
		return &errs.InvalidContactInfo
	}
	if len(loc.CompanyName) < 5 {
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
