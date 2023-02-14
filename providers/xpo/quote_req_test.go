package xpo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	v1 "github.com/ramsfords/types_gen/v1"
)

func TestXpoQuote(t *testing.T) {
	reqData := &v1.QuoteRequest{
		Pickup: &v1.Location{
			Address: &v1.Address{
				ZipCode: "90010",
			},
		},
		Delivery: &v1.Location{
			Address: &v1.Address{
				ZipCode: "77081",
			},
		},
		Commodities: []*v1.Commodity{
			{
				Weight:              1000,
				Height:              75,
				Length:              48,
				Width:               48,
				Quantity:            1,
				FreightClass:        5,
				ShipmentDescription: "novelties",
				WeightUOM: &v1.WeightUOM{
					LB: true,
					KG: false,
				},
				DimensionUOM: &v1.DimensionUOM{
					INCH: true,
					CM:   false,
				},
			},
		},
		LocationServices: &v1.LocationServices{
			InsidePickup:                true,
			InsideDelivery:              true,
			LiftGatePickup:              true,
			LiftGateDelivery:            true,
			PickupNotification:          true,
			DeliveryNotification:        true,
			ReceiverPickupNotification:  false,
			ShipperDeliveryNotification: false,
			PickupAppointment:           true,
			DeliveryAppointment:         false,
		},
		TotalItems:  1,
		TotalWeight: 1000,
		PickupDate:  "2023-01-30T00:00:00Z",
	}
	xpoQuoteReq := NEWXPOQuoteRequest(reqData)
	xpoReqData, err := json.Marshal(xpoQuoteReq)
	if err != nil {
		fmt.Println(err)
	}

	strData := string(xpoReqData)
	fmt.Println(strData)
	req, err := http.NewRequest(http.MethodPost, "https://api.ltl.xpo.com/rating/1.0/ratequotes", bytes.NewBufferString(strData))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header = make(http.Header)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer c2c2a24a-3142-3384-9154-538e2a37d7e6")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Body : %s", body)
	quoteResData := &XPOQuoteResponse{}
	err = json.Unmarshal(body, quoteResData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(quoteResData)
}
