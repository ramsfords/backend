package roadrunner

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	v1 "github.com/ramsfords/types_gen/v1"
)

func Test_QuoteRequest(t *testing.T) {
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
	reqs, err := NewRoadRunnerQuoteRequest(reqData)
	if err != nil {
		fmt.Println(err)
	}
	out, err := xml.MarshalIndent(&reqs, " ", "  ")
	fmt.Println(err)
	fmt.Println(string(out))
	body := string(out)
	client := &http.Client{}

	response, err := client.Post("https://webservices.rrts.com/rating/ratequote.asmx", "application/soap+xml", bytes.NewBufferString(body))
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	bodyRes := &RoadRunneQuoteResponse{}
	err = xml.Unmarshal(content, bodyRes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.StatusCode)
	s := strings.TrimSpace(string(content))
	fmt.Println(s)
}
