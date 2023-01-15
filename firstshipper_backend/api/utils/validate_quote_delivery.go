package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuoteDelivery(req *v1.Location) error {
	if req.Address.ZipCode == "" || len(req.Address.ZipCode) != 5 {
		return errs.ErrInvalidDeliveryZipCode
	}
	withdock := false
	liftgate := false
	if IncludesService(toInt32ArrayFromPickupLocationServices(req.PickupLocationServices), int32(v1.PickupLocationServices_PICKUP_LOCATION_WITH_DOCK)) {
		withdock = true
	}
	if IncludesService(toInt32ArrayFromDeliveryLocationServices(req.DeliveryLocationServices), int32(v1.DeliveryLocationServices_DELIVERY_LOCATION_WITH_DOCK)) {
		liftgate = true
	}

	if !withdock && liftgate || withdock && !liftgate {
		return errs.ErrInvalidPickupLocationServices.Err
	}
	return nil
}
