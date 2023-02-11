package booking_db

import (
	"context"

	v1 "github.com/ramsfords/types_gen/v1"
)

func (bookingdb BookingDb) UpdateBooking(ctx context.Context, booking *v1.BookingResponse, businessId string) error {
	return nil
}
