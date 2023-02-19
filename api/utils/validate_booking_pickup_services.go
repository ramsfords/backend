package utils

import (
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateLocationServices(pickLoc *v1.Location, bookPickup *v1.Location) *errs.ApiErr {
	valid := true
	if !valid {
		return &errs.ErrInvalidPickupLocationServices
	}
	return nil
}
