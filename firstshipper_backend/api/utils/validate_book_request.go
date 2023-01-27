package utils

import (
	"github.com/ramsfords/backend/menuloom_backend/api/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func ValidateBookRequest(qtReq *v1.QuoteRequest, bookReq *v1.QuoteRequest, bid v1.Bid) error {
	if qtReq.QuoteId != bid.QuoteId {
		return errs.ErrInvalidInputs
	}

	err := validateBookPickup(bookReq)
	if err != nil {
		return err
	}
	err = validateBookDelivery(bookReq)
	if err != nil {
		return err
	}

	err = validateBookRequestWithQuoteRequest(qtReq, bookReq)
	if err != nil {
		return err
	}
	return nil
}
