package utils

import (
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuoteDelivery(req *v1.QuoteRequest) error {
	if req.Delivery.Address.ZipCode == "" || len(req.Delivery.Address.ZipCode) != 5 {
		return errs.ErrInvalidDeliveryZipCode
	}
	if !(req.LocationServices.DeliveryLocationWithDock && !req.LocationServices.LiftGateDelivery) || (!req.LocationServices.DeliveryLocationWithDock && req.LocationServices.LiftGateDelivery) {
		return errs.ErrInvalidPickupLocationServices.Err
	}
	return nil
}
