package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
)

type Script struct {
	CustomScript string `json:"CustomScript,omitempty" dynamodbav:"CustomScript,omitempty"`
}
type BOlGenerateReq struct {
	Url             string `json:"url,omitempty" dynamodbav:"url,omityempty"`
	Margins         string `json:"margins,omitempty" dynamodbav:"margins,omityempty"`
	PaperSize       string `json:"paperSize,omitempty" dynamodbav:"paperSize,omityempty"`
	Orientation     string `json:"orientation,omitempty" dynamodbav:"orientation,omityempty"`
	PrintBackground bool   `json:"printBackground,omitempty" dynamodbav:"printBackground,omityempty"`
	Header          string `json:"header,omitempty" dynamodbav:"header,omityempty"`
	Footer          string `json:"footer,omitempty" dynamodbav:"footer,omityempty"`
	MediaType       string `json:"mediaType,omitempty" dynamodbav:"mediaType,omityempty"`
	Async           bool   `json:"async,omitempty" dynamodbav:"async,omityempty"`
	Encrypt         bool   `json:"encrypt,omitempty" dynamodbav:"encrypt,omityempty"`
	Profiles        Script `json:"profiles,omitempty" dynamodbav:"profiles,omityempty"`
}
type BolGenerateResponse struct {
	Url              string `json:"url,omitempty" dynamodbav:"url,omitempty"`
	PageCount        int    `json:"pageCount,omitempty" dynamodbav:"pageCount,omitempty"`
	Error            bool   `json:"error,omitempty" dynamodbav:"error,omitempty"`
	Status           int    `json:"status,omitempty" dynamodbav:"status,omitempty"`
	Name             string `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Credits          int32  `json:"credits,omitempty" dynamodbav:"credits,omitempty"`
	Duration         int32  `json:"duration,omitempty" dynamodbav:"duration,omitempty"`
	RemainingCredits int32  `json:"remainingCredits,omitempty" dynamodbav:"remainingCredits,omitempty"`
}

func GenerateBOL(conf configs.Config, bookingId string, cli *S3.S3Client) error {
	pdfRenderData := strings.NewReader(fmt.Sprintf(`{
		"url": "https://www.firstshipper.com/bol/%s",
		"margins": "5mm",
		"paperSize": "Letter",
		"orientation": "Portrait",
		"printBackground": true,
		"header": "",
		"footer": "",
		"mediaType": "print",
		"async": false,
		"encrypt": false,
		"profiles": "{ \"CustomScript\": \";; // put some custom js script here \"}"
	}`, bookingId))

	pdfRenderReq, err := http.NewRequest("POST", conf.PdfRenderer.Url, pdfRenderData)
	if err != nil {
		return err
	}
	pdfRenderReq.Header.Add("content-type", "application/json")
	pdfRenderReq.Header.Add("x-api-key", conf.PdfRenderer.ApiKey)
	pdfRenderRes, err := http.DefaultClient.Do(pdfRenderReq)
	if err != nil {
		return err
	}
	defer pdfRenderRes.Body.Close()
	resObj := &BolGenerateResponse{}
	err = json.NewDecoder(pdfRenderRes.Body).Decode(resObj)
	if err != nil {
		return err
	}
	fmt.Println(err)

	err = json.NewDecoder(pdfRenderReq.Body).Decode(resObj)
	if err != nil {
		return err
	}
	reqs, err := http.Get(resObj.Url)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(reqs.Body)
	if err != nil {
		return err
	}
	inputs := strings.NewReader(string(data))
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("bol" + bookingId + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               inputs,
	}
	_, err = cli.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		return err
	}
	return nil
}
