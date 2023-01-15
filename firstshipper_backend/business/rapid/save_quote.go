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

type SaveResponse struct {
	IsSucceeded  bool   `json:"isSucceeded"`
	SavedQuoteId string `json:"savedQuoteId"`
}

func (rapid Rapid) SaveQuote(rapidQuoteData *models.SaveQuote) (*SaveResponse, error) {
	// per-request timeout
	reqTimeout := time.Duration(20) * time.Second
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(rapid.SaveQuoteUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentTypeBytes([]byte("application/json"))
	authToken, err := rapid.Vault.GetAuthToken(vault.RapidShipLTL)
	if err != nil {
		return nil, err
	}
	fmt.Println(rapid.AuthTokenName)
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
	saveRes := SaveResponse{}
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
