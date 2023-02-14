package rapid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"
	"github.com/valyala/fasthttp"
)

func (rapid Rapid) Dispatch(rapidQuoteData *models.ConfirmAndDispatch) (*models.DispatchResponse, error) {
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(rapid.DispatchUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return nil, err
	}
	req.Header.Set(rapid.AuthTokenName, authToken)
	out, err := json.Marshal(rapidQuoteData)
	fmt.Println("\n \n \n dispatch request")
	fmt.Println(string(out))
	if err != nil {
		return nil, nil
	}
	req.SetBodyRaw(out)
	fmt.Println(string(out))
	resp := fasthttp.AcquireResponse()
	err = rapid.HttpClient.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	dispatchRes := models.DispatchResponse{}
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Println(string(respBody))
		if statusCode == http.StatusOK {
			err = json.Unmarshal(respBody, &dispatchRes)
			if err != io.EOF || err == nil {
				fmt.Println(dispatchRes)
				return &dispatchRes, nil
			} else {
				return nil, fmt.Errorf("ERR failed to parse reponse: %s", err)
			}
		} else {
			fmt.Println(string(respBody))
			return nil, fmt.Errorf("ERR invalid HTTP response code from rapid quote response: %d", statusCode)
		}
	} else {
		return nil, fmt.Errorf("WARN conn error: %s", err)
	}

}
