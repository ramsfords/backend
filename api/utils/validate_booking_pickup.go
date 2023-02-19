package utils

import (
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookingPickup(bookingshipperLoc *v1.Location, quoteShipperlocation *v1.Location) error {
	err := validateAddress(bookingshipperLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "pickup")
	}

	err = validateShippingParty(bookingshipperLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "pickup")
	}
	err = validateLocationServices(quoteShipperlocation, bookingshipperLoc)
	if err != nil {
		return errs.NewCustomErr(*err, "pickup")
	}
	return nil
}
