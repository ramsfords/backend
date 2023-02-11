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

func (rapid Rapid) Login(auth *models.AuthRequestPayload) error {
	var ErrorChan chan error
	for {
		// per-request timeout
		reqTimeout := time.Duration(10) * time.Second
		payload, err := json.Marshal(auth)
		if err != nil {
			ErrorChan <- err
			return err
		}
		req := fasthttp.AcquireRequest()
		req.SetRequestURI(rapid.AuthUrl)
		req.Header.SetMethod(fasthttp.MethodPost)
		req.Header.SetContentTypeBytes([]byte("application/json"))
		req.SetBodyRaw(payload)
		resp := fasthttp.AcquireResponse()
		err = rapid.Client.DoTimeout(req, resp, reqTimeout)
		fasthttp.ReleaseRequest(req)
		if err == nil {
			statusCode := resp.StatusCode()
			respBody := resp.Body()
			fmt.Printf("DEBUG Response: %s\n", respBody)
			if statusCode == http.StatusOK {
				authRes := models.AuthResponsePayload{}
				err = json.Unmarshal(respBody, &authRes)
				if err == io.EOF || err == nil {
					rapid.Vault.AddAuthToken(vault.RapidShipLTL, "Bearer "+authRes.Token)
				} else {
					return fmt.Errorf("ERR failed to parse reponse: %s\n", err)
				}

			} else {
				return fmt.Errorf("ERR invalid HTTP response code: %d\n", statusCode)
			}
		} else {
			return fmt.Errorf("WARN conn error: %s\n", err)
		}
		fasthttp.ReleaseResponse(resp)
		if ErrorChan != nil {
			close(ErrorChan)
			break
		}
		time.Sleep(58 * time.Minute)
	}
	return nil

}
