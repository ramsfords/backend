package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuoteDelivery(req *v1.QuoteRequest) error {
	if req.Delivery.Address.ZipCode == "" || len(req.Delivery.Address.ZipCode) != 5 {
		return errs.ErrInvalidDeliveryZipCode
	}
	if !(req.DeliveryLocationServices.DeliverLocationWithDock && !req.DeliveryLocationServices.LiftGateDelivery) || (!req.DeliveryLocationServices.DeliverLocationWithDock && req.DeliveryLocationServices.LiftGateDelivery) {
		return errs.ErrInvalidPickupLocationServices.Err
	}
	return nil
}
