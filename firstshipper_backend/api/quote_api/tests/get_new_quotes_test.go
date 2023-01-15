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
		Pickup: &v1.Location{
			Address: &v1.Address{
				ZipCode: "90650",
			},
			PickupLocationServices: []v1.PickupLocationServices{
				v1.PickupLocationServices_LIFTGATE_PICKUP,
			},
		},
		Delivery: &v1.Location{
			Address: &v1.Address{
				ZipCode: "77081",
			},
			DeliveryLocationServices: []v1.DeliveryLocationServices{
				v1.DeliveryLocationServices_DELIVERY_LOCATION_WITH_DOCK,
			},
		},
		Commodities: []*v1.Commodity{
			{
				Length:              48,
				Height:              75,
				Width:               40,
				DimensionUom:        0,
				WeightUom:           0,
				Quantity:            1,
				PackageType:         8,
				FreightClass:        4,
				Stackable:           false,
				CommodityServices:   []v1.CommodityServices{},
				ShipmentDescription: "novelties",
			},
		},
		ShipmentDetails: &v1.ShipmentDetails{
			TotalItems:  1,
			DisplayDate: "2022-08-12T17:00:00.000Z",
		},
	}
	json, err := json.Marshal(quote_req_data)
	body := bytes.NewReader(json)

	req, err := http.NewRequest("POST", "http://127.0.0.1:8090/carrierservice/quote/carrier/new_quote", body)
	if err != nil {
		// handle err
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
	fmt.Println("hell")
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	defer resp.Body.Close()
}
