package rapid

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"
	"github.com/valyala/fasthttp"
)

func (rapid Rapid) GetQuoteHistory() ([]models.QuoteHistory, error) {
	addresses := struct {
		TotalCount   int              `json:"totalCount"`
		AddressList  []models.Address `json:"addressList"`
		HasAddresses bool             `json:"hasAddresses"`
	}{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	// per-request timeout
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(rapid.QuoteHistoryUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return []models.QuoteHistory{}, err
	}
	req.Header.Set(rapid.AuthTokenName, authToken)
	req.SetBody([]byte(`{"addressFilter":{"term":null,"isCanada":true,"isUSA":true,"stateCode":null},"gridFilter":{"sortOrder":null,"sortField":"companyName","take":10,"skip":null,"page":1},"isFromEditPage":false}"`))
	resp := fasthttp.AcquireResponse()
	err = rapid.Client.DoTimeout(req, resp, reqTimeout)
	var quoteHistory []models.QuoteHistory
	fasthttp.ReleaseRequest(req)
	if err == nil {
		statusCode := resp.StatusCode()
		respBody := resp.Body()
		fmt.Printf("DEBUG Response: %s\n", respBody)
		if statusCode == http.StatusOK {

			fmt.Println("\n", string(respBody))
			err = json.Unmarshal(respBody, &addresses)
			fmt.Println(string(respBody))
			if err == io.EOF || err == nil {
				fmt.Printf("DEBUG Parsed Response: %v\n", quoteHistory)
			} else {
				fmt.Fprintf(os.Stderr, "ERR failed to parse reponse: %s\n", err)
			}
		} else {
			fmt.Fprintf(os.Stderr, "ERR invalid HTTP response code: %d\n", statusCode)
		}
	} else {
		fmt.Fprintf(os.Stderr, "WARN conn error: %s\n", err)
	}
	fasthttp.ReleaseResponse(resp)
	return quoteHistory, nil
}
