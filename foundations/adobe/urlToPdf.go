package adobe

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/menuloom_backend/api/errs"
)

type Adobe struct {
	S3.S3Client
}

func NewAdobe(s3Client S3.S3Client) *Adobe {
	return &Adobe{
		S3Client: s3Client,
	}
}
func (adobe Adobe) UrlToPdf(bolId string) (string, error) {
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf"
	firstShipperBolUrl := strings.ReplaceAll("https://firstshipper.com/bol/{BOLID}", "{BOLID}", bolId)
	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsIng1dSI6Imltc19uYTEta2V5LWF0LTEuY2VyIiwia2lkIjoiaW1zX25hMS1rZXktYXQtMSIsIml0dCI6ImF0In0.eyJpZCI6IjE2NzUyODAwMDk3NjdfMDc0OWExYmItNWExNy00NzdhLWJkZTctNGFlMTcxY2IxYjljX3V3MiIsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJjbGllbnRfaWQiOiIxYTA5OTUzNzhlOTY0ZTg1YTI2MGNlM2U5OGY1ZTZiNyIsInVzZXJfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiYXMiOiJpbXMtbmExIiwiYWFfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiY3RwIjowLCJmZyI6IlhGRkpMMzdaRlBGNUk3NEtNTVFWWkhZQVg0PT09PT09IiwibW9pIjoiODU3YjM4NCIsImV4cGlyZXNfaW4iOiI4NjQwMDAwMCIsInNjb3BlIjoib3BlbmlkLERDQVBJLEFkb2JlSUQsYWRkaXRpb25hbF9pbmZvLm9wdGlvbmFsQWdyZWVtZW50cyIsImNyZWF0ZWRfYXQiOiIxNjc1MjgwMDA5NzY3In0.OSLwRgQoDYzqlHAl55uOJQbQvVBVZw-n7SHbrWoCYV5ynrC1qbz7e7FjErnLOYmhLDlxjHZDXlrmzFsiQM0lhuEmp81y6pD5tjTLvl8nT0ODP0i4XMMrN6i7flPZnM31pQkAR_mcHm-B0IQoLkHfWgP-Nba-n78y-XlOcIXRzqxG0QNOVAkNzhodSbWlewX_bdBmU3HvR8NM_dEBBVBGYQzP0UeCWYPCqT0GBr_Bn25_lzO3jBtkJpDvAoScKB28oRVis36pQyryn99YR7bvoa5xPghOlOH_UG6UoG4HnINGMIQAI2GodOPb4QpQ4PCZg5yWXW4pixrKv-8C2_lgUw")
	req.Header.Add("x-api-key", "1a0995378e964e85a260ce3e98f5e6b7")
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
	adobe.PullObjectResult(pullObjectId)
	fmt.Println(location)
	fmt.Println(string(body))
	return "", nil
}
func (Adobe Adobe) PullObjectResult(jobId string) (adobeResourceURL string, err error) {
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{JOBID}/status"
	url = strings.Replace(url, "{JOBID}", jobId, 1)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsIng1dSI6Imltc19uYTEta2V5LWF0LTEuY2VyIiwia2lkIjoiaW1zX25hMS1rZXktYXQtMSIsIml0dCI6ImF0In0.eyJpZCI6IjE2NzUyODAwMDk3NjdfMDc0OWExYmItNWExNy00NzdhLWJkZTctNGFlMTcxY2IxYjljX3V3MiIsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJjbGllbnRfaWQiOiIxYTA5OTUzNzhlOTY0ZTg1YTI2MGNlM2U5OGY1ZTZiNyIsInVzZXJfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiYXMiOiJpbXMtbmExIiwiYWFfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiY3RwIjowLCJmZyI6IlhGRkpMMzdaRlBGNUk3NEtNTVFWWkhZQVg0PT09PT09IiwibW9pIjoiODU3YjM4NCIsImV4cGlyZXNfaW4iOiI4NjQwMDAwMCIsInNjb3BlIjoib3BlbmlkLERDQVBJLEFkb2JlSUQsYWRkaXRpb25hbF9pbmZvLm9wdGlvbmFsQWdyZWVtZW50cyIsImNyZWF0ZWRfYXQiOiIxNjc1MjgwMDA5NzY3In0.OSLwRgQoDYzqlHAl55uOJQbQvVBVZw-n7SHbrWoCYV5ynrC1qbz7e7FjErnLOYmhLDlxjHZDXlrmzFsiQM0lhuEmp81y6pD5tjTLvl8nT0ODP0i4XMMrN6i7flPZnM31pQkAR_mcHm-B0IQoLkHfWgP-Nba-n78y-XlOcIXRzqxG0QNOVAkNzhodSbWlewX_bdBmU3HvR8NM_dEBBVBGYQzP0UeCWYPCqT0GBr_Bn25_lzO3jBtkJpDvAoScKB28oRVis36pQyryn99YR7bvoa5xPghOlOH_UG6UoG4HnINGMIQAI2GodOPb4QpQ4PCZg5yWXW4pixrKv-8C2_lgUw")
	req.Header.Add("x-api-key", "1a0995378e964e85a260ce3e98f5e6b7")
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

	fmt.Println(res)
	// err = ioutil.WriteFile("file.pdf", body, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(string(body))
	return "", nil
}
func (Adobe Adobe) UploadBOlTOS3(adobeResourceURl string, poNumber string, businessId string, s3Client S3.S3Client) (string, error) {
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
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("bol" + poNumber + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
		Metadata: map[string]string{
			"businessId": businessId,
		},
	}
	s3res, err := s3Client.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		return "", errs.ErrInternal
	}
	fmt.Println(s3res)
	return "", nil
}
