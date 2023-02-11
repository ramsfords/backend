package rapid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ramsfords/backend/foundations/vault"
	"github.com/ramsfords/backend/shipper/business/rapid/models"
	"github.com/valyala/fasthttp"
)

func (rapid Rapid) GetShipmentById(id string) (*models.ShipmentByIdResponse, error) {
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	req := fasthttp.AcquireRequest()
	url := fmt.Sprintf("https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/Quote/GetShipmentById?shipmentId=%s", id)
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return nil, err
	}
	req.Header.Set(rapid.AuthTokenName, authToken)
	resp := fasthttp.AcquireResponse()
	err = rapid.Client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	saveRes := models.ShipmentByIdResponse{}
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		if statusCode == http.StatusOK {
			err = json.Unmarshal(respBody, &saveRes)
			if err != io.EOF || err == nil {
				fmt.Println(saveRes)
				return &saveRes, nil
			} else {
				return nil, fmt.Errorf("ERR failed to parse reponse: %s", err)
			}
		} else {
			return nil, fmt.Errorf("ERR invalid HTTP response code: %d", statusCode)
		}
	} else {
		return nil, fmt.Errorf("WARN conn error: %s", err)
	}

}
