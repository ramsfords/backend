package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuotePickup(req *v1.QuoteRequest) error {
	if req.Pickup.Address.ZipCode == "" || len(req.Pickup.Address.ZipCode) != 5 {
		return errs.ErrInvalidPickupZipCode
	}
	if !(req.PickupLocationServices.PickupLocationWithDock || !req.PickupLocationServices.LiftGatePickup) && (!req.PickupLocationServices.PickupLocationWithDock || req.PickupLocationServices.LiftGatePickup) {
		return errs.ErrInvalidPickupLocationServices.Err
	}
	return nil
}
