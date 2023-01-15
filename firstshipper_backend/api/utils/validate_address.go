package utils

import (
	valid "github.com/asaskevich/govalidator"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateAddress(add *v1.Location) *errs.ApiErr {
	if len(add.Address.ZipCode) != 5 && !valid.IsNumeric(add.Address.ZipCode) {
		return &errs.InvalidAddress
	}
	if len(add.Address.AddressLine_1) < 3 && add.Address.AddressLine_1 == "" {
		return &errs.InvalidAddress
	}
	if len(add.Address.City) < 3 && add.Address.City == "" {
		return &errs.InvalidAddress
	}
	if len(add.Address.Country) < 3 && add.Address.Country == "" {
		return &errs.InvalidAddress
	}
	return nil
}
