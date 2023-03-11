package ai_parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

// {\"response_as_dict\":true,\"attributes_as_list\":false,\"show_original_response\":false,\"providers\":\"amazon\",\"file\":\"data:image/jpeg;name=testinvoice.jpeg;base64,/
type Payload struct {
	ResponseAsDict       bool   `json:"response_as_dict"`
	AttributesAsList     bool   `json:"attributes_as_list"`
	ShowOriginalResponse bool   `json:"show_original_response"`
	Providers            string `json:"providers"`
	FileURL              string `json:"file_url"`
	Language             string `json:"language"`
}

func (aiParser AIParser) ParseInvoice(ctx echo.Context) error {
	ctx.Request().ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, err := ctx.FormFile("invoice")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	reqData := &ParseRequest{}
	for key, values := range ctx.Request().Form { // range over map
		for _, value := range values { // range over []string
			if key == "authContext" {
				err = json.Unmarshal([]byte(value), &reqData.AuthContext)
				if err != nil {
					fmt.Println(err)
				}

			}
			if key == "bookingResponse" {
				err = json.Unmarshal([]byte(value), &reqData.BookingResponse)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}
			if key == "customerId" {
				reqData.CustomerID = value
				continue
			}
			if key == "paymentMethod" {
				reqData.PaymentMethod = value
				continue
			}
			if key == "paymentTerm" {
				reqData.PaymentTerm = value
				continue
			}
			if key == "accountingPlatform" {
				reqData.AccountingPlatform = value
				continue
			}
			if key == "includeShipping" {
				if value == "true" {
					reqData.IncludeShipping = true
				} else {
					reqData.IncludeShipping = false
				}
				continue
			}
		}
	}
	reqData.FileName = file.Filename
	datas, err := file.Open()
	if err != nil {
		fmt.Println(err, datas)
	}
	// return ctx.JSON(http.StatusCreated, reqData)

	s3FileName := strings.Split(reqData.AuthContext.User.Email, "@")[0] + "/" + file.Filename
	err = aiParser.services.S3Client.UploadCustomerInvoices(datas, s3FileName, "image/jpeg")
	if err != nil {
		fmt.Println(err)
	}
	payLoadObj := Payload{
		ResponseAsDict:       true,
		AttributesAsList:     false,
		ShowOriginalResponse: false,
		Providers:            "amazon,google,microsoft",
		FileURL:              "https://customer-invoices-vault.s3.us-west-1.amazonaws.com/" + s3FileName,
		Language:             "en",
	}
	// convert struct to json
	payLoadObjJson, err := json.Marshal(payLoadObj)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(imgBase64Str)
	url := "https://api.edenai.run/v2/ocr/invoice_parser"
	payload := strings.NewReader(string(payLoadObjJson))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("Origin", "https://docs.edenai.co")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://docs.edenai.co/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36")
	req.Header.Set("accept", "application/json")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYjcxYjhkOTAtMjA1Yi00YWJmLThkMDAtNmVjNjc3NGM3NDRkIiwidHlwZSI6ImFwaV90b2tlbiJ9.NSf4b1zUp5jItIjTiyVlnVu13GvTrMRFOCmIzMgk7hU")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("sec-ch-ua", `"Chromium";v="110", "Not A(Brand";v="24", "Google Chrome";v="110"`)
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set("sec-ch-ua-platform", `"Android"`)
	req.Header.Set("x-readme-api-explorer", "4.349.1")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bodyText))
	respData := &ParserResponse{}
	err = json.Unmarshal(bodyText, respData)
	if err != nil {
		fmt.Println(err)
	}
	return ctx.JSON(http.StatusCreated, respData)
}
