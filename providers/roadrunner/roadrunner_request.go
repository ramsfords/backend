package roadrunner

import (
	"encoding/xml"
	"fmt"
	"time"

	v1 "github.com/ramsfords/types_gen/v1"
)

type AuthenticationHeader struct {
	X        string `xml:"xmlns,attr,omitempty"`
	UserName string `xml:"UserName,omitempty"`
	Password string `xml:"Password,omitempty"`
	Site     string `xml:"Site,omitempty"`
}
type Header struct {
	XMLName              xml.Name             `xml:"soap12:Header,omitempty"`
	AuthenticationHeader AuthenticationHeader `xml:"AuthenticationHeader,omitempty"`
}
type ShipmentDetail struct {
	ActualClass float32 `xml:"ActualClass,omitempty"`
	Weight      int     `xml:"Weight,omitempty"`
}
type ShipmentDetails struct {
	ShipmentDetail []ShipmentDetail `xml:"ShipmentDetail,omitempty"`
}
type ServiceOption struct {
	ServiceCode string `xml:"ServiceCode,omitempty"`
}
type ServiceDeliveryOptions struct {
	ServiceOptions []ServiceOption `xml:"ServiceOptions,omitempty"`
}
type COD struct {
	Prepaid   string `xml:"Prepaid,omitempty"`
	CODAmount string `xml:"CODAmount,omitempty"`
}
type Request struct {
	OriginZip              string                 `xml:"OriginZip,omitempty"`
	DestinationZip         string                 `xml:"DestinationZip,omitempty"`
	ShipmentDetails        ShipmentDetails        `xml:"ShipmentDetails,omitempty"`
	OriginType             string                 `xml:"OriginType,omitempty"`
	PaymentType            string                 `xml:"PaymentType,omitempty"`
	PalletCount            string                 `xml:"PalletCount,omitempty"`
	LinearFeet             string                 `xml:"LinearFeet,omitempty"`
	CubicFeet              string                 `xml:"CubicFeet,omitempty"`
	Pieces                 int32                  `xml:"Pieces,omitempty"`
	ServiceDeliveryOptions ServiceDeliveryOptions `xml:"ServiceDeliveryOptions,omitempty"`
	COD                    COD                    `xml:"COD,omitempty"`
	Discount               string                 `xml:"Discount,omitempty"`
	ListedConsigneeCity    string                 `xml:"ListedConsigneeCity,omitempty"`
	InternalUse            string                 `xml:"InternalUse,omitempty"`
	PalletPosition         string                 `xml:"PalletPosition,omitempty"`
	ShipDate               string                 `xml:"ShipDate,omitempty"`
}
type RateQuote struct {
	X       string  `xml:"xmlns,attr,omitempty"`
	Request Request `xml:"request,omitempty"`
}
type Body struct {
	XMLName   xml.Name  `xml:"soap12:Body,omitempty"`
	RateQuote RateQuote `xml:"RateQuote,omitempty"`
}
type Envelope struct {
	XMLName xml.Name `xml:"soap12:Envelope,omitempty"`
	X       string   `xml:"xmlns:soap12,attr,omitempty"`
	Header  Header   `xml:"soap12:Header,omitempty"`
	Body    Body     `xml:"soap12:Body,omitempty"`
}

func NewRoadRunnerQuoteRequest(quoteReq *v1.QuoteRequest) (*Envelope, error) {
	envolope := &Envelope{
		Header: Header{
			AuthenticationHeader: AuthenticationHeader{
				X:        "https://webservices.rrts.com/ratequote/",
				UserName: "ramfords",
				Password: "Ferina@1234",
			},
		},
		Body: Body{
			RateQuote: RateQuote{
				X: "https://webservices.rrts.com/ratequote/",
				Request: Request{
					OriginZip:      quoteReq.Pickup.Address.ZipCode,
					DestinationZip: quoteReq.Delivery.Address.ZipCode,
					OriginType:     "O",
					PaymentType:    "C",
					PalletCount:    fmt.Sprintf("%d", quoteReq.TotalItems),
					Pieces:         quoteReq.TotalItems,
					ShipDate:       "0",
				},
			},
		},
	}
	envolope.X = "http://www.w3.org/2003/05/soap-envelope"
	envolope.Body.RateQuote.X = "https://webservices.rrts.com/ratequote/"
	pickupDate, err := time.Parse(time.RFC3339, quoteReq.PickupDate)
	if err != nil {
		return nil, err
	}
	pickupDateFormat := pickupDate.Format("2006-01-02")
	envolope.Body.RateQuote.Request.ShipDate = pickupDateFormat
	envolope.Body.RateQuote.Request.ShipmentDetails = ShipmentDetails{
		ShipmentDetail: []ShipmentDetail{},
	}
	for _, j := range quoteReq.Commodities {
		envolope.Body.RateQuote.Request.ShipmentDetails.ShipmentDetail = append(
			envolope.Body.RateQuote.Request.ShipmentDetails.ShipmentDetail, ShipmentDetail{
				ActualClass: GetActualFrightClass(j.FreightClass),
				Weight:      int(j.Weight),
			})
	}
	envolope.Body.RateQuote.Request.ServiceDeliveryOptions = ServiceDeliveryOptions{
		ServiceOptions: GetDeliveryServices(quoteReq.LocationServices),
	}
	return envolope, nil
}
