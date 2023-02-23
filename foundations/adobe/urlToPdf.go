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
		logger.Error(err, "Error in exchange token")
	} else {
		adobe.TokenRes.ApiKey = "1a0995378e964e85a260ce3e98f5e6b7"
	}

	return adobe
}
func (adobe *Adobe) UrlToPdf(bid *v1.Bid, fileName string) (string, error) {
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
			logger.Error(err, "I am in loop.Error in url to pdf")
			break
		}

		defer res.Body.Close()
		if res.StatusCode != 201 {
			adobe.ExchangeToken()
		} else {
			tryAgain = false
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logger.Error(err, "I am in loop.Error in url to pdf")
			}
			location := res.Header.Get("Location")
			pullObjectId := strings.Split(strings.Split(location, "htmltopdf/")[1], "/")[0]
			adobe.PullObjectResult(pullObjectId, bid, fileName)
			fmt.Println(location)
			fmt.Println(string(body))
			return "", nil
		}

	}
	return "", nil
}
func (adobe *Adobe) PullObjectResult(pullObjectId string, bid *v1.Bid, fileName string) (adobeResourceURL string, err error) {
	time.Sleep(10 * time.Second)
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{PULLOBJECTID}/status"
	url = strings.Replace(url, "{PULLOBJECTID}", pullObjectId, 1)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", adobe.TokenRes.AccessToken)
	req.Header.Add("x-api-key", adobe.TokenRes.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(err, "Error in pull object from adobe")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err, "Error in pull object from adobe")
	}
	resData := AdoblePullRes{}
	err = json.Unmarshal(body, &resData)
	if err != nil {
		logger.Error(err, "Error in pull object from adobe")
	}
	adobe.UploadBOlTOS3(resData.Aseets.DownloadUri, bid, fileName)
	logger.Error(err, "finised PullObjectResult from adobe")
	return "", nil
}
func (adobe *Adobe) UploadBOlTOS3(adobeResourceURl string, bid *v1.Bid, fileName string) (string, error) {
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
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("BOL" + fileName + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
	}
	s3res, err := adobe.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		logger.Error(err, "Error in putting object in s3")
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
		logger.Error(err, "Error in exchanging token from adobe")
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Error(err, "Error in exchanging token from adobe")
		return err
	}
	token := &TokenRes{}
	err = json.Unmarshal(body, token)
	if err != nil || res.Status != "200 OK" {
		logger.Error(err, "Error in exchanging token from adobe")
		return err
	}
	token.AccessToken = "Bearer " + token.AccessToken
	adobe.TokenRes.AccessToken = token.AccessToken
	adobe.TokenRes.ExpiresIn = token.ExpiresIn
	adobe.TokenRes.TokenType = token.TokenType
	return nil
}

func (adobe *Adobe) GetToken() (string, error) {
	jsonData := `{
		"variables": {
			"privateKey": "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDe8gfsF8VspKsucFKY91dWvoEu\nVpQNhBiKmM3fRe9ig6nrhBEfg3LcNDyVYfUo9ZR9WVg3gocL3bQnn7EyqeAYZhN+2meB1k7w4Nhx\nQwtvWX73Cwim+ICBoL1Aq8n0WMog8DIMuGwFKpucynCu9AF9xa51zXthY24BVEs+/eowe93HLvad\nU+38PL9mTU0x/z3rJlHDo9ROsRFN/Pdku6dgKqBRDMbMX5zqxErEXhnhZ4OuVLCFB74kMrqJy1+w\nzKMHfYDgzDWaP8rNvYzD0kmeN+LJvWvpuvBjkiMoeubr2Kh0gH7IgUdEncoyJGQnTAM7V3MSjaUs\nyqrPnWmgGh1JAgMBAAECggEAEQYJqVMyWcakJJZilDgUxPSkgBGP2g1bY1/iHnmkWxEjLS7nMNEU\nePCLLLvvYKqJ5V1oYUq3/aW27yygNvZmPG607+OE8lLXMcj1dgaQzbaXfY0r4rYId/16XgEQGXab\nLpMwuyxs4SMdAUaq/oz2vTAWT3v6fxf8yyCw4zU3x+5U2ZNLJMS4pCV5S9r5u7v0VOLE7DsawNgW\ndfG/e+t676KZfFTlRMNip9jAdHcl9qh/wi1KbtUkj/m5jG551EUZ2tmBm3Nngzn675sy/0uUNDtq\nP3BkLxAbaUIbBqoYigdO+0WH4j59uqgiQ0GPCnyXIOF8kllaoZ6MhG5C1w8oRQKBgQDiIa0+2Z9P\nb9snsA2qVGemzgyxCIJ9OiBmRMpEIl9JINkl3G6nOZwgQuwzxJnAXM/V0fVQ5twBPmtGpv7UIrPn\naL+s4I6vuu8KZjW7DGr3bnZS3kzf0DCf7gg06XLqboTleUAfGrCRgleBxo6yZdHuda0/Y8MMs2sg\nctzNI6mvbQKBgQD8ZJ7AMHihUU1myhif86TPlfpRK3fy7mkVPDdLkbTWgqn8WG++rLr/xU1dD0Mb\nO2OyEn5eZYYfwGG9UChYTyo1kSQXm0G7gNRoSGujs5AeT4s1nD4C2WNPC+MObDI35sgG8FQzpQxf\nZiybP4+XbvRnPOehaLyAw7HEgLOAWK5PzQKBgQCubh3OIl5SD0100sfnwI3nzH9bu920LTc6zAtd\n/UmVBkKaguvUEItPE3BSCnAzQySKohdiHYJNb1GewUhGaLJvaYyZFOrbwQ2M7wS5UT3duRaKm7Ge\n31/yrdEkx4L+NNxMingcxiC3TVyk3X3LPOFv2NQX1qNpU6jp8dBCaSUGDQKBgHqOJWPvZXP3tZz2\n/1QUC/BcxCrL77e/urj7/2Grg+MxmXcWPlSZLUhrNvC8K3q6sONUBZayt5kNYqh5ls2iyz0tmBf6\nZMW2fe2RVOstkwqU12UV1CqwAn/sprlnIk9wuapc4pYdS8+7HmfYSlJfJ0BGG7eN0xK3c8eWMxNc\nfG/1AoGAVzzIUOXgnM487aiBR91i60RnybiAbTJJiJ+PFHTuw0PvSux58LwLDKasieVJ/YQJlhqt\ni7xnX9MsRsaE3BY8J+Bfyz/JVT5xlGC9NdpkgReWXb7z+ppCh+UGryH+3BenAyrSP1K0FgDjI9Pv\n1p9dEDWZOlDOoPIHeOEe5Fxv1UE=\n-----END PRIVATE KEY-----",
			"orgId": "971322",
			"intId": "394004",
			"clientId": "1a0995378e964e85a260ce3e98f5e6b7",
			"clientSecret": "p8e-PG3RgfPMq_b3t_KT9ZhLbPJJ1LASyPMA"
		},
		"query": "query ($privateKey: String!, $orgId: String!, $intId: String!, $clientId: String, $clientSecret: String) {\n  getJwt(orgId: $orgId, intId: $intId) {\n    jwtToken(privateKey: $privateKey, clientId: $clientId, clientSecret: $clientSecret) {\n      token\n      command\n      __typename\n    }\n    __typename\n  }\n}"
		}`

	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", "https://developer.adobe.com/console/graphql", bytes.NewBuffer(jsonValue))
	if err != nil {
		logger.Error(err, "Error in exchanging token from adobe")
		return "", err
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil && response.StatusCode == 200 {
		logger.Error(err, "Error in exchanging token from adobe")
		return "", err
	}
	defer response.Body.Close()
	if err != nil {
		logger.Error(err, "Error in exchanging token from adobe")
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	return string(data), nil
}
