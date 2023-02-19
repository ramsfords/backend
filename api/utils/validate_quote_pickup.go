package utils

import (
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuotePickup(req *v1.QuoteRequest) error {
	if req.Pickup.Address.ZipCode == "" || len(req.Pickup.Address.ZipCode) != 5 {
		return errs.ErrInvalidPickupZipCode
	}
	if !(req.LocationServices.PickupLocationWithDock || !req.LocationServices.LiftGatePickup) && (!req.LocationServices.PickupLocationWithDock || req.LocationServices.LiftGatePickup) {
		return errs.ErrInvalidPickupLocationServices.Err
	}
	return nil
}
