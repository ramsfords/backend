package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	v1 "github.com/ramsfords/types_gen/v1"
)

func Test_New_Get_Quotes(t *testing.T) {
	quote_req_data := v1.QuoteRequest{
		LocationServices: &v1.LocationServices{
			PickupLocationWithDock:      true,
			DeliveryLocationWithDock:    true,
			LiftGatePickup:              true,
			LiftGateDelivery:            true,
			InsidePickup:                true,
			InsideDelivery:              true,
			PickupAppointment:           true,
			DeliveryAppointment:         true,
			PickupNotification:          true,
			DeliveryNotification:        true,
			ShipperDeliveryNotification: true,
			ReceiverPickupNotification:  true,
		},
		Pickup: &v1.Location{
			Address: &v1.Address{
				ZipCode: "90650",
			},
		},
		Delivery: &v1.Location{
			Address: &v1.Address{
				ZipCode: "77081",
			},
		},
		Commodities: []*v1.Commodity{
			{
				Length:              48,
				Height:              75,
				Width:               40,
				Quantity:            1,
				PackageType:         8,
				FreightClass:        4,
				ShipmentDescription: "novelties",
			},
		},
	}
	json, err := json.Marshal(quote_req_data)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	fmt.Println(string(json))
	body := bytes.NewReader(json)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8090/carrierservice/quote/carrier/new_quote", body)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  "login_guard",
		Value: "ysoGbIJvszeVLGvkjdqFyqiGiAUJDCupAWtjbbNkURjvJtkqFH",
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	fmt.Println(string(bodyBytes))
	defer resp.Body.Close()
}
