package adobe

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/dgrijalva/jwt-go"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

//go:embed refresh_token.json
var refreshToken string

type TokenRes struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
type MetaData struct {
	Type string `json:"type"`
	Size int64  `json:"size"`
}
type Asset struct {
	MetaData    MetaData `json:"metadata"`
	DownloadUri string   `json:"downloadUri"`
	AssetId     string   `json:"assetId"`
}
type AdoblePullRes struct {
	Status string `json:"status"`
	Aseets Asset  `json:"asset"`
}
type Adobe struct {
	S3.S3Client
	Conf       *configs.Config
	localToken LocalToken
	ExpiresAt  time.Time
	Claims     jwt.MapClaims
}
type LocalToken struct {
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresAt string `json:"refresh_expires_at"`
	AccessToken      string `json:"access_token"`
	AccessExpiresAt  string `json:"access_expires_at"`
	ExpiresIn        int64  `json:"expires_in"`
	TokenType        string `json:"token_type"`
}

func NewAdobe(s3Client S3.S3Client, conf *configs.Config) (*Adobe, error) {
	// parse embbeded token
	localToken := &LocalToken{}
	err := json.Unmarshal([]byte(refreshToken), localToken)
	if err != nil {
		fmt.Println("Failed to parse JWT token:", err)
		return nil, err
	}
	accessExp, err := strconv.ParseInt(localToken.AccessExpiresAt, 10, 64)
	if err != nil {
		fmt.Println("Failed to get expiry time for access token", err)
	}

	adobe := &Adobe{
		S3Client:   s3Client,
		localToken: *localToken,
		Conf:       conf,
	}
	// parse unix time to time.Time
	expirationTime := time.Unix(accessExp, 0)
	if time.Now().After(expirationTime) || adobe.localToken.AccessToken == "" {
		adobe.exchangeToken()
	}
	return adobe, nil
}
func (adobe *Adobe) UrlToPdf(bookingResponse *v1.BookingResponse, fileName string) (string, error) {
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf"
	firstShipperBolUrl := ""
	if adobe.Conf.Env == "dev" {
		firstShipperBolUrl = strings.ReplaceAll("http://127.0.0.1:3000/bol/{BOLID}", "{BOLID}", bookingResponse.Bid.BidId)
		logger.Error(nil, "I am in dev mode. Url to pdf is "+firstShipperBolUrl)
		return firstShipperBolUrl, nil
	} else {
		firstShipperBolUrl = strings.ReplaceAll("https://www.firstshipper.com/bol/{BOLID}", "{BOLID}", bookingResponse.Bid.BidId)
	}
	count := 0
	for count < 10 {
		req, _ := http.NewRequest("POST", url, nil)

		req.Header.Add("Authorization", adobe.localToken.AccessToken)
		req.Header.Add("x-api-key", adobe.Conf.Adobe.ClientId)
		req.Header.Add("Content-Type", "application/json")
		reqBody := `{
				"pageLayout": {
					"pageWidth": 9,
					"pageHeight": 12.5
				},
				"includeHeaderFooter": false,
				"json": "{}",
				"inputUrl": "{INPUTURL}"
			}`
		reqBody = strings.ReplaceAll(reqBody, "{INPUTURL}", firstShipperBolUrl)
		req.Body = ioutil.NopCloser(bytes.NewBufferString(reqBody))
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			logger.Error(err, "I am in loop.Error in url to pdf")
			count++
			continue
		}

		defer res.Body.Close()
		if res.StatusCode != 201 || time.Now().After(adobe.ExpiresAt) {
			adobe.exchangeToken()
			count++
			continue
		} else {
			location := res.Header.Get("Location")
			pullObjectId := strings.Split(strings.Split(location, "htmltopdf/")[1], "/")[0]
			return adobe.PullObjectResult(pullObjectId, bookingResponse, fileName)
		}
	}
	return "", nil
}
func (adobe *Adobe) PullObjectResult(pullObjectId string, bookingResponse *v1.BookingResponse, fileName string) (adobeResourceURL string, err error) {
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{PULLOBJECTID}/status"
	url = strings.Replace(url, "{PULLOBJECTID}", pullObjectId, 1)
	resData := AdoblePullRes{}
	count := 0
	for count < 10 {
		resData = AdoblePullRes{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", adobe.localToken.AccessToken)
		req.Header.Add("x-api-key", adobe.Conf.Adobe.ClientId)
		req.Header.Add("Content-Type", "application/json")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			count++
			continue
		}
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			count++
			continue
		}
		err = json.Unmarshal(body, &resData)
		if err != nil {
			count++
			continue
		}
		if res.StatusCode == 200 && resData.Status != "done" {
			res.Body.Close()
			time.Sleep(2 * time.Second)
			count++
			// logger.Error(err, "Error in pull object from adobe")
		} else if resData.Status == "failed" {
			adobe.UrlToPdf(bookingResponse, fileName)

		} else if resData.Status == "done" {
			return adobe.UploadBOlTOS3(resData.Aseets.DownloadUri, bookingResponse, fileName)
		}
	}
	return "", nil
}
func (adobe *Adobe) UploadBOlTOS3(adobeResourceURl string, bookingResponse *v1.BookingResponse, fileName string) (string, error) {
	// gets pdf from provided url
	reqs, err := http.Get(adobeResourceURl)
	if err != nil {
		logger.Error(err, "Error in pull object from adobe")
	}
	pdfBytes, err := io.ReadAll(reqs.Body)
	if err != nil {
		logger.Error(err, "Error in pull object from adobe")
	}
	reqs.Body.Close()
	//TotalWeight string
	totalWeightStr := fmt.Sprintf("%f", bookingResponse.QuoteRequest.TotalWeight)
	// TotalItems string
	totalItemsStr := fmt.Sprintf("%d", bookingResponse.QuoteRequest.TotalItems)
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String(fileName),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
		Metadata: map[string]string{
			"bolUrl":                bookingResponse.BookingInfo.BolUrl,
			"carrierProNumber":      bookingResponse.BookingInfo.CarrierProNumber,
			"firstShipperBolNumber": bookingResponse.BookingInfo.FirstShipperBolNumber,
			"totalWeight":           totalWeightStr,
			"totalItems":            totalItemsStr,
			"carrierName":           bookingResponse.BookingInfo.CarrierName,
			"carrierPhone":          bookingResponse.BookingInfo.CarrierPhone,
			"shipperCompanyName":    bookingResponse.QuoteRequest.Pickup.CompanyName,
			"shipperName":           bookingResponse.QuoteRequest.Pickup.Contact.Name,
			"shipperAddressLine1":   bookingResponse.QuoteRequest.Pickup.Address.AddressLine1,
			"shipperCity":           bookingResponse.QuoteRequest.Pickup.Address.City,
			"shipperZipCode":        bookingResponse.QuoteRequest.Pickup.Address.ZipCode,
			"shipperState":          bookingResponse.QuoteRequest.Pickup.Address.State,
			"receiverCompanyName":   bookingResponse.QuoteRequest.Delivery.CompanyName,
			"receiverName":          bookingResponse.QuoteRequest.Delivery.Contact.Name,
			"receiverAddressLine1":  bookingResponse.QuoteRequest.Delivery.Address.AddressLine1,
			"receiverCity":          bookingResponse.QuoteRequest.Delivery.Address.City,
			"receiverZipCode":       bookingResponse.QuoteRequest.Delivery.Address.ZipCode,
			"receiverState":         bookingResponse.QuoteRequest.Delivery.Address.State,
			"bookingRequesterEmail": bookingResponse.QuoteRequest.Pickup.Contact.EmailAddress,
		},
	}
	res, err := adobe.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		logger.Error(err, "Error in putting object in s3")
	}
	return *res.Expiration, nil
}
func (adobe *Adobe) exchangeToken() (*TokenRes, error) {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("client_id")
	if err != nil {
		return nil, err
	}
	_, err = formField.Write([]byte(adobe.Conf.Adobe.ClientId))
	if err != nil {
		return nil, err
	}
	formField, err = writer.CreateFormField("client_secret")
	if err != nil {
		return nil, err
	}
	_, err = formField.Write([]byte(adobe.Conf.Adobe.ClientSecret))
	if err != nil {
		return nil, err
	}
	formField, err = writer.CreateFormField("jwt_token")
	if err != nil {
		return nil, err
	}
	_, err = formField.Write([]byte(adobe.localToken.RefreshToken))
	if err != nil {
		return nil, err
	}
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://ims-na1.adobelogin.com/ims/exchange/jwt/", form)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tokenRes := &TokenRes{}
	err = json.Unmarshal(bodyText, tokenRes)
	if err != nil {
		return nil, err
	}
	adobe.localToken.AccessToken = tokenRes.AccessToken
	adobe.localToken.AccessExpiresAt = fmt.Sprint(time.Now().Add(24 * time.Hour).Unix())
	adobe.ExpiresAt = time.Now().Add(24 * time.Hour)
	adobe.localToken.RefreshExpiresAt = fmt.Sprint(time.Now().Add(time.Hour * 24 * 60).Unix())
	//convert to string bytes
	bytes, err := json.Marshal(adobe.localToken)
	if err != nil {
		return nil, err
	}
	// write to file
	err = ioutil.WriteFile("./foundations/adobe/refresh_token.json", bytes, 0644)
	if err != nil {
		return nil, err
	}
	return tokenRes, nil
}
