package rapid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"
	"github.com/valyala/fasthttp"
)

func (rapid Rapid) GetQuoteById(rapidQuoteData model.QuoteRequest) (*models.SaveQuote, error) {
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	req := fasthttp.AcquireRequest()
	url := fmt.Sprintf("https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/SaveQuote/GetSavedQuote?savedQuoteId=%s&isReRunShipment=%v", rapidQuoteData.SaveQuoteResponse.SavedQuoteID, false)
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return nil, err
	}
	req.Header.Set(rapid.AuthTokenName, authToken)
	out, err := json.Marshal(rapidQuoteData)
	if err != nil {
		return nil, nil
	}
	req.SetBodyRaw(out)
	fmt.Println(string(out))
	resp := fasthttp.AcquireResponse()
	err = rapid.Client.DoTimeout(req, resp, reqTimeout)
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	saveRes := models.SaveQuote{}
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
