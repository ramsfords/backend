package roadrunner

import "encoding/xml"

type Customer struct {
	Text          string `xml:",chardata"`
	AccountNumber string `xml:"AccountNumber"`
	Name          string `xml:"Name"`
	Address1      string `xml:"Address1"`
	Address2      string `xml:"Address2"`
	City          string `xml:"City"`
	State         string `xml:"State"`
	ZipCode       string `xml:"ZipCode"`
}
type RoutingInfo struct {
	Text                 string `xml:",chardata"`
	DestinationState     string `xml:"DestinationState"`
	DestinationZip       string `xml:"DestinationZip"`
	OriginState          string `xml:"OriginState"`
	OriginZip            string `xml:"OriginZip"`
	EstimatedTransitDays string `xml:"EstimatedTransitDays"`
	OriginTerminal       string `xml:"OriginTerminal"`
}
type QuoteDetail []struct {
	Text          string `xml:",chardata"`
	ActualClass   string `xml:"ActualClass"`
	RatedClass    string `xml:"RatedClass"`
	Charge        string `xml:"Charge"`
	Code          string `xml:"Code"`
	Description   string `xml:"Description"`
	Rate          string `xml:"Rate"`
	Weight        string `xml:"Weight"`
	ExtraMessages string `xml:"ExtraMessages"`
}
type RateDetails struct {
	Text        string      `xml:",chardata"`
	QuoteDetail QuoteDetail `xml:"QuoteDetail"`
}
type RateQuoteResult struct {
	Text                   string      `xml:",chardata"`
	QuoteNumber            string      `xml:"QuoteNumber"`
	NetCharge              string      `xml:"NetCharge"`
	Customer               Customer    `xml:"Customer"`
	RoutingInfo            RoutingInfo `xml:"RoutingInfo"`
	RateDetails            RateDetails `xml:"RateDetails"`
	OriginType             string      `xml:"OriginType"`
	PaymentType            string      `xml:"PaymentType"`
	CODAmount              string      `xml:"CODAmount"`
	ShipmentDate           string      `xml:"ShipmentDate"`
	CustomerCubicFoot      string      `xml:"CustomerCubicFoot"`
	HawaiianRatedCubicFoot string      `xml:"HawaiianRatedCubicFoot"`
}
type ResponseBody struct {
	Text              string            `xml:",chardata"`
	RateQuoteResponse RateQuoteResponse `xml:"RateQuoteResponse"`
}
type RateQuoteResponse struct {
	Text            string          `xml:",chardata"`
	Xmlns           string          `xml:"xmlns,attr"`
	RateQuoteResult RateQuoteResult `xml:"RateQuoteResult"`
}
type RoadRunneQuoteResponse struct {
	XMLName      xml.Name     `xml:"Envelope"`
	Text         string       `xml:",chardata"`
	Soap         string       `xml:"soap,attr"`
	Xsi          string       `xml:"xsi,attr"`
	Xsd          string       `xml:"xsd,attr"`
	ResponseBody ResponseBody `xml:"Body"`
}
