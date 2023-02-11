package utils

import (
	"errors"

	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookPickup(req *v1.QuoteRequest) error {
	if req.Pickup.Address.ZipCode == "" || len(req.Pickup.Address.ZipCode) != 5 {
		return errors.New("invalid pickup zip code")
	}
	if !(req.LocationServices.PickupLocationWithDock || !req.LocationServices.LiftGatePickup) && (!req.LocationServices.PickupLocationWithDock || req.LocationServices.LiftGatePickup) {
		return errors.New("invalid pickup location services")
	}
	if req.Pickup.Address.AddressLine1 == "" || len(req.Pickup.Address.AddressLine1) < 5 {
		return errors.New("invalid pickup address line 1")
	}
	if req.Pickup.Address.City == "" || len(req.Pickup.Address.City) < 2 {
		return errors.New("invalid pickup city")
	}
	if req.Pickup.Address.State == "" || len(req.Pickup.Address.State) < 2 {
		return errors.New("invalid pickup state")
	}
	if req.Pickup.CompanyName == "" || len(req.Pickup.CompanyName) < 2 {
		return errors.New("invalid shipper company name")
	}
	if req.PickupDate == "" || len(req.PickupDate) < 2 {
		return errors.New("invalid pickup date")
	}
	return nil
}
