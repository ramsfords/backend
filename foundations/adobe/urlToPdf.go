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
	"github.com/ramsfords/backend/menuloom_backend/api/errs"
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
}

func NewAdobe(s3Client S3.S3Client) *Adobe {
	adobe := &Adobe{
		S3Client: s3Client,
	}
	err := adobe.ExchangeToken()
	if err != nil {
		fmt.Println(err)
	} else {
		adobe.TokenRes.ApiKey = "1a0995378e964e85a260ce3e98f5e6b7"
	}

	return adobe
}
func (adobe *Adobe) UrlToPdf(bolId string, businessId string) (string, error) {
	tryAgain := true
	for tryAgain {
		url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf"
		firstShipperBolUrl := strings.ReplaceAll("https://firstshipper.com/bol/{BOLID}", "{BOLID}", "30009-1")
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
			fmt.Println(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 201 {
			adobe.ExchangeToken()
		} else {
			tryAgain = false
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(res)
			// err = ioutil.WriteFile("file.pdf", body, 0644)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			location := res.Header.Get("Location")
			pullObjectId := strings.Split(strings.Split(location, "htmltopdf/")[1], "/")[0]
			adobe.PullObjectResult(pullObjectId, bolId, businessId)
			fmt.Println(location)
			fmt.Println(string(body))
			return "", nil
		}

	}
	return "", nil
}
func (adobe *Adobe) PullObjectResult(pullObjectId string, bolId string, businessId string) (adobeResourceURL string, err error) {
	time.Sleep(5 * time.Second)
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{PULLOBJECTID}/status"
	url = strings.Replace(url, "{PULLOBJECTID}", pullObjectId, 1)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", adobe.TokenRes.AccessToken)
	req.Header.Add("x-api-key", adobe.TokenRes.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	resData := AdoblePullRes{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		fmt.Println(err)
	}
	adobe.UploadBOlTOS3(resData.Aseets.DownloadUri, bolId, businessId)
	return "", nil
}
func (adobe *Adobe) UploadBOlTOS3(adobeResourceURl string, bolId string, businessId string) (string, error) {
	// gets pdf from provided url
	reqs, err := http.Get(adobeResourceURl)
	if err != nil {
		return "", errs.ErrInternal
	}
	pdfBytes, err := io.ReadAll(reqs.Body)
	if err != nil {
		return "", errs.ErrInternal
	}
	reqs.Body.Close()
	quoteId := strings.Split(bolId, "-")[0]
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("BOL" + quoteId + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
		Metadata: map[string]string{
			"businessId": businessId,
		},
	}
	s3res, err := adobe.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		return "", errs.ErrInternal
	}
	fmt.Println(s3res)
	return "", nil
}
func (adobe *Adobe) ExchangeToken() error {
	urls := "https://ims-na1.adobelogin.com/ims/exchange/jwt/"
	form := url.Values{}
	form.Add("client_id", "1a0995378e964e85a260ce3e98f5e6b7")
	form.Add("client_secret", "p8e-PG3RgfPMq_b3t_KT9ZhLbPJJ1LASyPMA")
	form.Add("jwt_token", "eyJhbGciOiJSUzI1NiJ9.eyJleHAiOjE2NzU0NjE4NzMsImlzcyI6IjhGNzk2MUFFNjJGQTcyNEMwQTQ5NUM5M0BBZG9iZU9yZyIsInN1YiI6IkVFREEyNjQ5NjNEOTU5RDkwQTQ5NUVFQ0B0ZWNoYWNjdC5hZG9iZS5jb20iLCJodHRwczovL2ltcy1uYTEuYWRvYmVsb2dpbi5jb20vcy9lbnRfZG9jdW1lbnRjbG91ZF9zZGsiOnRydWUsImF1ZCI6Imh0dHBzOi8vaW1zLW5hMS5hZG9iZWxvZ2luLmNvbS9jLzFhMDk5NTM3OGU5NjRlODVhMjYwY2UzZTk4ZjVlNmI3In0.wKaTWPW7iZsNgzEsNdWDZuKaJBIfMPJrmqStxJM8CDQe1nZOEK8IWQycDrcTfd3KNAEuj2qvMUcCEY19L8cSUWBdzhLMTnH1wKoef_6SomgPzmUTgTw9_-0twPqaTzbkt_Ckzs-ajjEt-zHEEBQgI_sz1cNgvf9QawGuLEhKnV0l98fh3SHB5EirxgYbqZ4A_wBAxioDCR7mi86RN9kmGwGXdYeDWZFnXFPbDa6k5WmGQGPXL6cfmvrph4NkmYgWEmg3hwQgNVbpl9Ef4fgbUubV3CQ496-cHGCLGawbOP4qfWdV_rv5zN0VuKlrTPEK7UNIFOhe0NY6VDLuwe8LtA")

	res, err := http.PostForm(urls, form)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	token := &TokenRes{}
	err = json.Unmarshal(body, token)
	if err != nil || res.Status != "200 OK" {
		return err
	}
	token.AccessToken = "Bearer " + token.AccessToken
	adobe.TokenRes.AccessToken = token.AccessToken
	adobe.TokenRes.ExpiresIn = token.ExpiresIn
	adobe.TokenRes.TokenType = token.TokenType
	return nil
}
