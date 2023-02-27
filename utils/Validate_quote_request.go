package utils

import v1 "github.com/ramsfords/types_gen/v1"

func ValidateQuoteRequest(qtReq *v1.QuoteRequest) error {
	err := validateQuotePickup(qtReq)
	if err != nil {
		return err
	}
	err = validateQuoteDelivery(qtReq)
	if err != nil {
		return err
	}
	err = validateQuoteCommodities(qtReq)
	if err != nil {
		return err
	}

	return nil
}
