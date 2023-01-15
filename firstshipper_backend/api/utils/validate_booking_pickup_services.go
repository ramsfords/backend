package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateLocationServices(pickLoc *v1.Location, bookPickup *v1.Location) *errs.ApiErr {
	valid := true
	for _, j := range bookPickup.PickupLocationServices {
		if !IncludesService(toInt32ArrayFromPickupLocationServices(pickLoc.PickupLocationServices), int32(*j.Enum())) {
			valid = false
		}
	}
	if !valid {
		return &errs.ErrInvalidPickupLocationServices
	}
	return nil
}
