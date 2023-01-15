package rapid

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"
	"github.com/valyala/fasthttp"
)

func (rapid Rapid) GetQuote(rapidQuoteData *models.QuoteDetails) (*models.QuoteRate, error) {
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(rapid.RateUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return nil, err
	}
	req.Header.Set(rapid.AuthTokenName, authToken)
	out, err := json.Marshal(rapidQuoteData)
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
	quoteRes := models.QuoteRate{}
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Println(string(respBody))
		if statusCode == http.StatusOK {
			err = json.Unmarshal(respBody, &quoteRes)
			if err != io.EOF || err == nil {
				fmt.Println(quoteRes)
				return &quoteRes, nil
			} else {
				return nil, fmt.Errorf("ERR failed to parse reponse: %s", err)
			}
		} else {
			return nil, fmt.Errorf("ERR invalid HTTP response code from rapid quote response: %d", statusCode)
		}
	} else {
		return nil, fmt.Errorf("WARN conn error: %s", err)
	}

}
