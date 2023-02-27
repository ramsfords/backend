package utils

import (
	"time"

	v1 "github.com/ramsfords/types_gen/v1"
)

func SanitizeBusiness(req *v1.User) (businessDb *v1.Business, err error) {
	// set new businessId to request object
	dbBusiness := v1.Business{
		Type:                 "business",
		BusinessEmail:        req.Email,
		AccountingEmail:      req.Email,
		HighPriorityEmail:    req.Email,
		CustomerServiceEmail: req.Email,
		CreatedAt:            time.Now().String(),
		AdminEmail:           req.Email,
		NeedsAddressUpdate:   true,
	}
	return &dbBusiness, nil
}
