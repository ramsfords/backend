package emails

import v1 "github.com/ramsfords/types_gen/v1"

func (email Email) SendBookingConfirmationEmail(data EmailData, bookingRes *v1.BookingResponse) error {
	return email.SendEmail(email.getBookingConfirmationTempalte(data, bookingRes), data)
}
