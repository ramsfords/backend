package utils

import v1 "github.com/ramsfords/types_gen/v1"

func ValidateBookingRequest(quoteReq *v1.QuoteRequest, oldQuote *v1.QuoteRequest) error {
	err := validateBookingPickup(quoteReq.Pickup, oldQuote.Pickup)
	if err != nil {
		return err
	}
	err = validateBookingDelivery(quoteReq.Delivery, oldQuote.Delivery)
	if err != nil {
		return err
	}
	err = validateBookingCommodities(quoteReq, oldQuote)
	if err != nil {
		return err
	}
	return nil
}
