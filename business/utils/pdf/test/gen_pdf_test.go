package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestPdfGen(t *testing.T) {

	// {
	//     "options": {
	//       "requestConfig": {
	//         "method": "POST",
	//         "uriTemplate": "https://pdf-services.adobe.io/operation/{operationName}",
	//         "connectTimeout": 10000,
	//         "readTimeout": 10000,
	//         "uriParams": {
	//           "operationName": "htmltopdf"
	//         }
	//       },
	//       "headers": {
	//         "x-api-app-info": "@adobe/pdfservices-node-sdk-3.2.0",
	//         "x-api-key": "1a0995378e964e85a260ce3e98f5e6b7",
	//         "x-request-id": "38f21d20-1a39-404c-bca5-35b65c9e3092",
	//         "x-dcsdk-ops-info": "Create PDF Operation",
	//         "Content-Type": "application/json"
	//       },
	//       "authenticate": true
	//     },
	//     "multipartData": {},
	//     "identityAccess": {
	//       "token": null,
	//       "audience": "https://ims-na1.adobelogin.com/c/1a0995378e964e85a260ce3e98f5e6b7",
	//       "endpoint": "https://ims-na1.adobelogin.com/ims/exchange/jwt/"
	//     },
	//     "content": {
	//       "pageLayout": {
	//         "pageWidth": 8.3,
	//         "pageHeight": 11.7
	//       },
	//       "includeHeaderFooter": false,
	//       "json": "{}",
	//       "inputUrl": "https://firstshipper.com/pdfdemo"
	//     }
	//   }
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsIng1dSI6Imltc19uYTEta2V5LWF0LTEuY2VyIiwia2lkIjoiaW1zX25hMS1rZXktYXQtMSIsIml0dCI6ImF0In0.eyJpZCI6IjE2NzUyODAwMDk3NjdfMDc0OWExYmItNWExNy00NzdhLWJkZTctNGFlMTcxY2IxYjljX3V3MiIsInR5cGUiOiJhY2Nlc3NfdG9rZW4iLCJjbGllbnRfaWQiOiIxYTA5OTUzNzhlOTY0ZTg1YTI2MGNlM2U5OGY1ZTZiNyIsInVzZXJfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiYXMiOiJpbXMtbmExIiwiYWFfaWQiOiJFRURBMjY0OTYzRDk1OUQ5MEE0OTVFRUNAdGVjaGFjY3QuYWRvYmUuY29tIiwiY3RwIjowLCJmZyI6IlhGRkpMMzdaRlBGNUk3NEtNTVFWWkhZQVg0PT09PT09IiwibW9pIjoiODU3YjM4NCIsImV4cGlyZXNfaW4iOiI4NjQwMDAwMCIsInNjb3BlIjoib3BlbmlkLERDQVBJLEFkb2JlSUQsYWRkaXRpb25hbF9pbmZvLm9wdGlvbmFsQWdyZWVtZW50cyIsImNyZWF0ZWRfYXQiOiIxNjc1MjgwMDA5NzY3In0.OSLwRgQoDYzqlHAl55uOJQbQvVBVZw-n7SHbrWoCYV5ynrC1qbz7e7FjErnLOYmhLDlxjHZDXlrmzFsiQM0lhuEmp81y6pD5tjTLvl8nT0ODP0i4XMMrN6i7flPZnM31pQkAR_mcHm-B0IQoLkHfWgP-Nba-n78y-XlOcIXRzqxG0QNOVAkNzhodSbWlewX_bdBmU3HvR8NM_dEBBVBGYQzP0UeCWYPCqT0GBr_Bn25_lzO3jBtkJpDvAoScKB28oRVis36pQyryn99YR7bvoa5xPghOlOH_UG6UoG4HnINGMIQAI2GodOPb4QpQ4PCZg5yWXW4pixrKv-8C2_lgUw")
	req.Header.Add("x-api-key", "1a0995378e964e85a260ce3e98f5e6b7")
	req.Header.Add("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`{
        "pageLayout": {
            "pageWidth": 9,
            "pageHeight": 12.5
        },
        "includeHeaderFooter": false,
        "json": "{}",
        "inputUrl": "https://firstshipper.com/pdfdemo"
    }`))

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
	PullObjectResult(pullObjectId)
	fmt.Println(location)
	fmt.Println(string(body))
}
func PullObjectResult(jobId string) {
	url := "https://pdf-services-ue1.adobe.io/operation/htmltopdf/{jobID}/status"
	url = strings.Replace(url, "{jobID}", jobId, 1)
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
}

func TestSVG(t *testing.T) {
	url := "https://bwipjs-api.metafloor.com/?bcid=code128&text=1234567890"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(string(body))
}
