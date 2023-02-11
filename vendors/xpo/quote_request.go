package xpo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	v1 "github.com/ramsfords/types_gen/v1"
)

type Dimensions struct {
	Length        int    `json:"length"`
	Width         int    `json:"width"`
	Height        int    `json:"height"`
	DimensionsUom string `json:"dimensionsUom"`
}
type GrossWeight struct {
	Weight    int    `json:"weight"`
	WeightUom string `json:"weightUom"`
}
type Address struct {
	Name         string `json:"name"`
	AddressLine1 string `json:"addressLine1"`
	AddressLine2 string `json:"addressLine2"`
	CityName     string `json:"cityName"`
	StateCd      string `json:"stateCd"`
	CountryCd    string `json:"countryCd"`
	PostalCd     string `json:"postalCd"`
}
type Shipper struct {
	AcctInstID string  `json:"acctInstId"`
	AcctMadCd  string  `json:"acctMadCd"`
	Address    Address `json:"address"`
}
type Consignee struct {
	Address Address `json:"address"`
}
type Accessorial struct {
	AccessorialCd string `json:"accessorialCd"`
}
type Bill2Party struct {
	AcctInstID string  `json:"acctInstId"`
	Address    Address `json:"address"`
}
type Commodity struct {
	PieceCnt     int         `json:"pieceCnt"`
	PackageCode  string      `json:"packageCode"`
	GrossWeight  GrossWeight `json:"grossWeight"`
	Desc         interface{} `json:"desc"`
	NmfcClass    string      `json:"nmfcClass"`
	HazmatInd    interface{} `json:"hazmatInd"`
	Dimensions   Dimensions  `json:"dimensions"`
	StackableInd interface{} `json:"stackableInd"`
	Charge       Charge      `json:"charge"`
}
type QuoteShipmentInfo struct {
	Bill2Party         Bill2Party         `json:"bill2Party"`
	Accessorials       []Accessorial      `json:"accessorials"`
	Commodity          []Commodity        `json:"commodity"`
	Shipper            Shipper            `json:"shipper"`
	Consignee          Consignee          `json:"consignee"`
	PaymentTermCd      string             `json:"paymentTermCd"`
	GarntInd           bool               `json:"garntInd"`
	G12Ind             bool               `json:"g12Ind"`
	RrsInd             bool               `json:"rrsInd"`
	LinealFt           int                `json:"linealFt"`
	ShipmentDate       string             `json:"shipmentDate"`
	Comment            string             `json:"comment"`
	PalletCnt          string             `json:"palletCnt"`
	TotCommodityCharge TotCommodityCharge `json:"totCommodityCharge"`
	TotCommodityWeight TotCommodityWeight `json:"totCommodityWeight"`
}
type XPOQuoteRequest struct {
	QuoteShipmentInfo         QuoteShipmentInfo `json:"shipmentInfo"`
	ManualQuoteInd            bool              `json:"manualQuoteInd"`
	ShowSpecialServiceCharges bool              `json:"showSpecialServiceCharges"`
}

func NEWXPOQuoteRequest(quoteReq *v1.QuoteRequest) *XPOQuoteRequest {
	xpoQuote := &XPOQuoteRequest{
		QuoteShipmentInfo: QuoteShipmentInfo{},
	}
	xpoQuote.QuoteShipmentInfo.ShipmentDate = quoteReq.PickupDate
	xpoQuote.QuoteShipmentInfo.Shipper.Address.PostalCd = quoteReq.Pickup.Address.ZipCode
	xpoQuote.QuoteShipmentInfo.Consignee.Address.PostalCd = quoteReq.Delivery.Address.ZipCode
	xpoQuote.QuoteShipmentInfo.PalletCnt = fmt.Sprintf("%d", quoteReq.TotalItems)
	xpoQuote.QuoteShipmentInfo.PaymentTermCd = "C"
	xpoQuote.QuoteShipmentInfo.Accessorials = getAssorials(quoteReq)
	xpoQuote.QuoteShipmentInfo.Commodity = []Commodity{}
	for _, item := range quoteReq.Commodities {
		commodity := Commodity{
			Desc:       item.ShipmentDescription,
			HazmatInd:  false,
			Dimensions: Dimensions{},
		}
		commodity.PieceCnt = int(item.Quantity)
		commodity.PackageCode = getPackageType(item)
		commodity.GrossWeight.Weight = int(item.Weight)
		commodity.GrossWeight.WeightUom = getWeightUOM(item)
		commodity.Desc = item.ShipmentDescription
		commodity.HazmatInd = false
		commodity.NmfcClass = getClass(item)
		commodity.Dimensions.Length = int(item.Length)
		commodity.Dimensions.Width = int(item.Width)
		commodity.Dimensions.Height = int(item.Height)
		commodity.Dimensions.DimensionsUom = getDimensionUOM(item)
		commodity.StackableInd = false
		xpoQuote.QuoteShipmentInfo.Commodity = append(xpoQuote.QuoteShipmentInfo.Commodity, commodity)
	}

	return xpoQuote

}
func getAssorials(quoteReq *v1.QuoteRequest) []Accessorial {
	accessorials := []Accessorial{}
	if quoteReq.LocationServices.LiftGatePickup {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "OLG"})
	}
	if quoteReq.LocationServices.InsidePickup {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "OIP"})
	}
	if quoteReq.LocationServices.PickupAppointment {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "APT"})
	}
	if quoteReq.LocationServices.LiftGateDelivery {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "DLG"})
	}
	if quoteReq.LocationServices.InsideDelivery {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "DID"})
	}
	if quoteReq.LocationServices.DeliveryNotification {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "DNC"})
	}
	if quoteReq.LocationServices.DeliveryAppointment {
		accessorials = append(accessorials, Accessorial{AccessorialCd: "APT"})
	}

	return accessorials
}
func getPackageType(commodity *v1.Commodity) string {
	if commodity.PackageType.String() == "PALLET" {
		return "PLT"
	}
	if commodity.PackageType.String() == "BOX" {
		return "BOX"
	}
	if commodity.PackageType.String() == "CRATE" {
		return "CRATE"
	}
	if commodity.PackageType.String() == "DRUM" {
		return "DRUM"
	}
	if commodity.PackageType.String() == "CARTON" {
		return "CTN"
	}
	if commodity.PackageType.String() == "CASE" {
		return "CASE"
	}
	if commodity.PackageType.String() == "BUNDLE" {
		return "BDL"
	}

	return ""
}
func getClass(commodity *v1.Commodity) string {
	if commodity.FreightClass == v1.FreightClass_CLASS50 {
		return "50"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS55 {
		return "55"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS60 {
		return "60"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS65 {
		return "65"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS70 {
		return "70"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS775 {
		return "77.5"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS85 {
		return "85"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS925 {
		return "92.5"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS100 {
		return "100"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS110 {
		return "110"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS125 {
		return "125"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS150 {
		return "150"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS175 {
		return "175"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS200 {
		return "200"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS250 {
		return "250"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS300 {
		return "300"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS400 {
		return "400"
	}
	if commodity.FreightClass == v1.FreightClass_CLASS500 {
		return "500"
	}
	return ""
}
func getDimensionUOM(commodity *v1.Commodity) string {
	if commodity.DimensionUOM.CM {
		return "CM"
	}
	if commodity.DimensionUOM.INCH {
		return "INCH"
	}
	return "INCH"
}
func getWeightUOM(commodity *v1.Commodity) string {
	if commodity.WeightUOM.KG {
		return "KG"
	}
	if commodity.WeightUOM.LB {
		return "LBS"
	}
	return "LBS"
}
func GetXPOQuote(quoteReq *v1.QuoteRequest) (*XPOQuoteResponse, error) {
	xpoQuote := NEWXPOQuoteRequest(quoteReq)
	c := http.Client{Timeout: time.Duration(20) * time.Second}
	req := http.Request{
		Method: "POST",
		URL:    &url.URL{Scheme: "https", Host: "api.ltl.xpo.com", Path: "rating/v1/ratequotes"},
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer c2c2a24a-3142-3384-9154-538e2a37d7e6")
	reqData, err := json.Marshal(xpoQuote)
	if err != nil {
		return nil, err
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(reqData))
	resp, err := c.Do(&req)
	if err != nil {
		fmt.Printf("Error %s", err)

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Body : %s", body)
	quoteResData := &XPOQuoteResponse{}
	err = json.Unmarshal(body, quoteResData)
	if err != nil {
		return nil, err
	}
	return quoteResData, nil
}
