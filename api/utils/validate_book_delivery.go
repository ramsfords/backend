package utils

import (
	"errors"

	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookDelivery(req *v1.QuoteRequest) error {
	if req.Delivery.Address.ZipCode == "" || len(req.Delivery.Address.ZipCode) != 5 {
		return errors.New("invalid delivery zip code")
	}
	if !(req.LocationServices.DeliveryLocationWithDock || !req.LocationServices.LiftGateDelivery) && (!req.LocationServices.DeliveryLocationWithDock || req.LocationServices.LiftGateDelivery) {
		return errors.New("invalid delivery location services")
	}
	if req.Delivery.Address.AddressLine1 == "" || len(req.Delivery.Address.AddressLine1) < 5 {
		return errors.New("invalid delivery address line 1")
	}
	if req.Delivery.Address.City == "" || len(req.Delivery.Address.City) < 2 {
		return errors.New("invalid delivery city")
	}
	if req.Delivery.Address.State == "" || len(req.Delivery.Address.State) < 2 {
		return errors.New("invalid delivery state")
	}
	if req.Delivery.CompanyName == "" || len(req.Delivery.CompanyName) < 2 {
		return errors.New("invalid delivery company name")
	}
	if req.Delivery.Contact.Name == "" || len(req.Delivery.Contact.Name) < 2 {
		return errors.New("invalid delivery contact name")
	}
	if req.Delivery.Contact.EmailAddress == "" || len(req.Delivery.Contact.EmailAddress) < 2 {
		return errors.New("invalid delivery contact email address")
	}
	if req.Delivery.Contact.PhoneNumber == "" || len(req.Delivery.Contact.PhoneNumber) < 2 {
		return errors.New("invalid delivery contact phone number")
	}
	return nil
}
