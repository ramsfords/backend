package email

import (
	"fmt"
	"strings"

	v1 "github.com/ramsfords/types_gen/v1"
)

func (email Email) getBookingConfirmationTempalte(data EmailData, bookingRes *v1.BookingResponse) string {
	template := strings.ReplaceAll(bookingConfirmationEmailTemplate, "{DOWNLOADBOL_URL}", bookingRes.BookingInfo.BolUrl)
	template = strings.ReplaceAll(template, "{SHIPMENT_PONUMBER}", bookingRes.BookingInfo.CarrierProNumber)
	template = strings.ReplaceAll(template, "{FIRSTSHIPPER_BOL_NUMBER}", bookingRes.BookingInfo.FirstShipperBolNumber)
	template = strings.ReplaceAll(template, "{SHIPMENT_WEIGHT}", floatToString(bookingRes.QuoteRequest.TotalWeight))
	template = strings.ReplaceAll(template, " {SHIPMENT_COUNT}", floatToString(float32(bookingRes.QuoteRequest.TotalItems)))
	template = strings.ReplaceAll(template, "{CARRIER_NAME}", bookingRes.BookingInfo.CarrierName)
	template = strings.ReplaceAll(template, "{CARRIER_EMAIL}", bookingRes.BookingInfo.CarrierEmail)
	template = strings.ReplaceAll(template, "{CARRIER_CONTACT_NUMBER}", bookingRes.BookingInfo.CarrierPhone)
	template = strings.ReplaceAll(template, "{SHIPPER_COMPANY_NAME}", bookingRes.QuoteRequest.Pickup.CompanyName)
	template = strings.ReplaceAll(template, "{SHIPPER_NAME}", bookingRes.QuoteRequest.Pickup.Contact.Name)
	template = strings.ReplaceAll(template, "{SHIPPER_ADDRESS1}", bookingRes.QuoteRequest.Pickup.Address.AddressLine1)
	template = strings.ReplaceAll(template, "{SHIPPER_CITY}", bookingRes.QuoteRequest.Pickup.Address.City)
	template = strings.ReplaceAll(template, "{SHIPPER_ZIPCODE}", bookingRes.QuoteRequest.Pickup.Address.ZipCode)
	template = strings.ReplaceAll(template, "{SHIPPER_STATE}", bookingRes.QuoteRequest.Pickup.Address.State)
	template = strings.ReplaceAll(template, "{RECEIVER_COMPANY_NAME}", bookingRes.QuoteRequest.Delivery.CompanyName)
	template = strings.ReplaceAll(template, "{RECEIVER_NAME}", bookingRes.QuoteRequest.Delivery.Contact.Name)
	template = strings.ReplaceAll(template, "{RECEIVER_ADDRESS1}", bookingRes.QuoteRequest.Delivery.Address.AddressLine1)
	template = strings.ReplaceAll(template, "{RECEIVER_CITY}", bookingRes.QuoteRequest.Delivery.Address.City)
	template = strings.ReplaceAll(template, "{RECEIVER_ZIPCODE}", bookingRes.QuoteRequest.Delivery.Address.ZipCode)
	template = strings.ReplaceAll(template, "{RECEIVER_STATE}", bookingRes.QuoteRequest.Delivery.Address.State)
	fmt.Println(template)
	return template
}
func intToString(i int) string {
	return strings.Split(fmt.Sprintf("%d", i), ".")[0]
}
func floatToString(f float32) string {
	return strings.Split(fmt.Sprintf("%f", f), ".")[0]
}
