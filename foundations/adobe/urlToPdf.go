package adobe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
	"golang.org/x/crypto/bcrypt"
)

type TokenRes struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	ApiKey      string `json:"api_key"`
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
	TokenRes TokenRes `json:"token_res"`
	logger   *logger.Logger
}

func NewAdobe(s3Client S3.S3Client, logger *logger.Logger) *Adobe {
	adobe := &Adobe{
		S3Client: s3Client,
		logger:   logger,
	}
	err := adobe.ExchangeToken()
	if err != nil {
		logger.Errorf("Error in exchange token %v", err)
	} else {
		adobe.TokenRes.ApiKey = "1a0995378e964e85a260ce3e98f5e6b7"
	}

	return adobe
}
func (adobe *Adobe) UrlToPdf(bid *v1.Bid) (string, error) {
	adobe.logger.Infof("starting url to pdf %v", time.Now().Local())
	tryAgain := true
	for tryAgain {
		url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf"
		firstShipperBolUrl := strings.ReplaceAll("https://firstshipper.com/bol/{BOLID}", "{BOLID}", bid.BidId)
		req, _ := http.NewRequest("POST", url, nil)

		req.Header.Add("Authorization", adobe.TokenRes.AccessToken)
		req.Header.Add("x-api-key", adobe.TokenRes.ApiKey)
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
			adobe.logger.Errorf("I am in loop.Error in url to pdf %v", err)
			break
		}

		defer res.Body.Close()
		if res.StatusCode != 201 {
			adobe.ExchangeToken()
		} else {
			tryAgain = false
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				adobe.logger.Errorf("I am in loop.Error in url to pdf %v", err)
			}
			location := res.Header.Get("Location")
			pullObjectId := strings.Split(strings.Split(location, "htmltopdf/")[1], "/")[0]
			adobe.PullObjectResult(pullObjectId, bid)
			fmt.Println(location)
			fmt.Println(string(body))
			return "", nil
		}

	}
	adobe.logger.Infof("finished url to pdf %v", time.Now().Local())
	return "", nil
}
func (adobe *Adobe) PullObjectResult(pullObjectId string, bid *v1.Bid) (adobeResourceURL string, err error) {
	adobe.logger.Infof("starint pull object from adobe %v", time.Now().Local())
	time.Sleep(10 * time.Second)
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{PULLOBJECTID}/status"
	url = strings.Replace(url, "{PULLOBJECTID}", pullObjectId, 1)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", adobe.TokenRes.AccessToken)
	req.Header.Add("x-api-key", adobe.TokenRes.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		adobe.logger.Errorf("Error in pull object from adobe %v", err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		adobe.logger.Errorf("Error in pull object from adobe %v", err)
	}
	resData := AdoblePullRes{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		adobe.logger.Errorf("Error in pull object from adobe %v", err)
	}
	adobe.UploadBOlTOS3(resData.Aseets.DownloadUri, bid)
	adobe.logger.Info("finised PullObjectResult from adobe")
	return "", nil
}
func (adobe *Adobe) UploadBOlTOS3(adobeResourceURl string, bid *v1.Bid) (string, error) {
	adobe.logger.Infof("starting upload bol to s3", time.Now().Local())
	// gets pdf from provided url
	reqs, err := http.Get(adobeResourceURl)
	if err != nil {
		adobe.logger.Errorf("Error in pull object from adobe %v", err)
	}
	pdfBytes, err := io.ReadAll(reqs.Body)
	if err != nil {
		adobe.logger.Errorf("Error in pull object from adobe %v", err)
	}
	reqs.Body.Close()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bid.QuoteId), bcrypt.DefaultCost)
	if err != nil {
		adobe.logger.Errorf("Error in created hashed bol %v", err)
	}
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("BOL" + string(hashedPassword) + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
	}
	s3res, err := adobe.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		adobe.logger.Errorf("Error in putting object in s3 %v", err)
	}
	fmt.Println(s3res)
	adobe.logger.Info("finished upload bol to s3")
	return "", nil
}
func (adobe *Adobe) ExchangeToken() error {
	adobe.logger.Infof("starting exchange token form adobe %v", time.Now().Local())
	urls := "https://ims-na1.adobelogin.com/ims/exchange/jwt/"
	form := url.Values{}
	form.Add("client_id", "1a0995378e964e85a260ce3e98f5e6b7")
	form.Add("client_secret", "p8e-PG3RgfPMq_b3t_KT9ZhLbPJJ1LASyPMA")
	form.Add("jwt_token", "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE2NzU0NjE4NzMsImlzcyI6IjhGNzk2MUFFNjJGQTcyNEMwQTQ5NUM5M0BBZG9iZU9yZyIsInN1YiI6IkVFREEyNjQ5NjNEOTU5RDkwQTQ5NUVFQ0B0ZWNoYWNjdC5hZG9iZS5jb20iLCJodHRwczovL2ltcy1uYTEuYWRvYmVsb2dpbi5jb20vcy9lbnRfZG9jdW1lbnRjbG91ZF9zZGsiOnRydWUsImF1ZCI6Imh0dHBzOi8vaW1zLW5hMS5hZG9iZWxvZ2luLmNvbS9jLzFhMDk5NTM3OGU5NjRlODVhMjYwY2UzZTk4ZjVlNmI3In0.wKaTWPW7iZsNgzEsNdWDZuKaJBIfMPJrmqStxJM8CDQe1nZOEK8IWQycDrcTfd3KNAEuj2qvMUcCEY19L8cSUWBdzhLMTnH1wKoef_6SomgPzmUTgTw9_-0twPqaTzbkt_Ckzs-ajjEt-zHEEBQgI_sz1cNgvf9QawGuLEhKnV0l98fh3SHB5EirxgYbqZ4A_wBAxioDCR7mi86RN9kmGwGXdYeDWZFnXFPbDa6k5WmGQGPXL6cfmvrph4NkmYgWEmg3hwQgNVbpl9Ef4fgbUubV3CQ496-cHGCLGawbOP4qfWdV_rv5zN0VuKlrTPEK7UNIFOhe0NY6VDLuwe8LtA")

	res, err := http.PostForm(urls, form)
	if err != nil {
		adobe.logger.Errorf("Error in exchanging token from adobe %v", err)
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		adobe.logger.Errorf("Error in exchanging token from adobe %v", err)
		return err
	}
	token := &TokenRes{}
	err = json.Unmarshal(body, token)
	if err != nil || res.Status != "200 OK" {
		adobe.logger.Errorf("Error in exchanging token from adobe %v", err)
		return err
	}
	token.AccessToken = "Bearer " + token.AccessToken
	adobe.TokenRes.AccessToken = token.AccessToken
	adobe.TokenRes.ExpiresIn = token.ExpiresIn
	adobe.TokenRes.TokenType = token.TokenType
	adobe.logger.Infof("done exchange token form adobe %v", time.Now().Local())
	return nil
}
