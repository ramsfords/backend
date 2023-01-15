package rapid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"

	"github.com/valyala/fasthttp"
)

func (rapid Rapid) AddAddress(conf configs.RapidShipLTL) error {
	quoteRes := models.QuoteRate{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	// per-request timeout
	facility := models.Address{
		City:                "Norwalk",
		CompanyName:         "Test2",
		CountryCode:         "US",
		DeliveryFromTime:    "9:00:00 AM",
		DeliveryToTime:      "5:00:00 PM",
		PostalCode:          "90650",
		ShippingFromTime:    "9:00:00 AM",
		ShippingToTime:      "5:00:00 PM",
		State:               "California",
		StateCode:           "CA",
		StreetLine1:         "14812 Gridley Road",
		StreetLine2:         "apt 41",
		Lat:                 33.8960202,
		Long:                -118.0907096,
		AddressAccessorials: []models.AddressAccessorial{},
		AddressContacts:     []*models.Contact{},
	}
	payload, err := json.Marshal(facility)
	if err != nil {
		return err
	}
	fmt.Println(string(payload))
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(rapid.AddAddressUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return err
	}
	req.Header.Set(conf.AuthTokenName, authToken)
	req.Header.Cookie("")
	req.SetBodyRaw(payload)
	resp := fasthttp.AcquireResponse()
	err = rapid.Client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Printf("DEBUG Response: %s\n", respBody)
		if statusCode == http.StatusOK {
			fmt.Println(string(respBody))

			err = json.Unmarshal(respBody, &quoteRes)
			if err == io.EOF || err == nil {
				fmt.Printf("DEBUG Parsed Response: %v\n successes")
			} else {
				fmt.Fprintf(os.Stderr, "ERR failed to parse reponse: %s", err)
			}

		} else {
			fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)
		}
	} else {
		fmt.Fprintf(os.Stderr, "WARN conn error: %s\n", err)
	}
	fasthttp.ReleaseResponse(resp)
	return nil

}
