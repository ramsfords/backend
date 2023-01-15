package models

import v1 "github.com/ramsfords/types_gen/v1"

type Wrapper struct {
	Booking *v1.BookingResponse `json:"booking"`
	Quote   *v1.QuoteRequest    `json:"quote"`
	Bid     *v1.Bid             `json:"bid"`
}
