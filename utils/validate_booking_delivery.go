package utils

import (
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookingDelivery(bookRecLoc *v1.Location, quoteRecLocation *v1.Location) error {
	err := validateAddress(bookRecLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "delivery")
	}

	err = validateShippingParty(bookRecLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "delivery")
	}
	err = validateLocationServices(quoteRecLocation, bookRecLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "delivery")
	}
	return nil
}
