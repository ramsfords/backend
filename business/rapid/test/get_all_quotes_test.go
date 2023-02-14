package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
)

type GetAllQuotes struct {
	TotalCount int `json:"totalCount"`
	SourceList []struct {
		SavedQuoteID         string        `json:"savedQuoteId"`
		Step                 int           `json:"step"`
		ReferenceID          string        `json:"referenceId"`
		PickupDate           string        `json:"pickupDate"`
		ShipperCompanyName   string        `json:"shipperCompanyName"`
		OriginAddress        interface{}   `json:"originAddress"`
		ConsigneeCompanyName interface{}   `json:"consigneeCompanyName"`
		DestinationAddress   interface{}   `json:"destinationAddress"`
		CarrierName          interface{}   `json:"carrierName"`
		TotalCost            string        `json:"totalCost"`
		ServiceType          int           `json:"serviceType"`
		ServiceTypeName      string        `json:"serviceTypeName"`
		QuoteID              int           `json:"quoteId"`
		TotalWeight          interface{}   `json:"totalWeight"`
		TimeCreated          interface{}   `json:"timeCreated"`
		ShipmentID           interface{}   `json:"shipmentId"`
		IsFavorite           bool          `json:"isFavorite"`
		IsArchivedQuote      bool          `json:"isArchivedQuote"`
		IsShipmentCanceled   bool          `json:"isShipmentCanceled"`
		OrderID              interface{}   `json:"orderId"`
		ErrorSteps           []interface{} `json:"errorSteps"`
		OriginCode           interface{}   `json:"originCode"`
	} `json:"sourceList"`
}

func Test_GetAllQuotes(t *testing.T) {
	for {
		// qt := MakeQuoteWithJson()
		// // fmt.Println(qt)
		// rapidQuote := rapid.MakeQuote(qt)
		// fmt.Println(rapidQuote)
		// payloadBytess, err := json.Marshal(rapidQuote)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(payloadBytess))
		body := []byte(`{"isFavorite":false,"isSystemAdmin":false,"startDate":"Invalid date","endDate":"Invalid date","uploadId":null,"isLocationFilterDisabled":null,"gridFilter":{"sortOrder":true,"sortField":"ActivityDateTime","take":15,"skip":0,"page":null},"stepFilter":{"isQuoteDetails":true,"isCarrierSelection":true,"isConfirmAndDispatch":true,"isBol":true,"isCancelled":true,"isError":null},"saveQuoteSearch":{"shipperCompanyName":null,"originAddress":null,"consigneeCompanyName":null,"destinationAddress":null,"carrierName":null,"referenceId":null},"serviceTypeFilter":{"isLtl":true,"isVltl":true,"isTruckload":true}}`)
		bodys := bytes.NewReader(body)
		req, err := http.NewRequest("POST", "https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/SaveQuote/GetSaveQuotes", bodys)
		if err != nil {
			// handle err
		}
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Cache-Control", "no-cache")
		req.Header.Set("Cookie", "_gcl_au=1.1.1739013701.1652374336; _gid=GA1.2.251636717.1655151308; _gat_UA-114313627-1=1; _ga=GA1.1.1423264958.1654897822; __hstc=192659957.95ae10d589bf6987cd16bd998ade8017.1655260289021.1655260289021.1655260289021.1; hubspotutk=95ae10d589bf6987cd16bd998ade8017; __hssrc=1; __hssc=192659957.1.1655260289021; intercom-id-c9oc6fab=574a48b8-0cb5-4f39-9d33-55eb122ef365; intercom-session-c9oc6fab=KzhDU2V5WHFtY2NWNHJjOGpGZ25oanlvMHcxbVlVenpDMFJTUUZSSVJ1eXNhYWFzMWdidmdxMHR2WEt0Q0V3cS0tbldlQzBXT000NWJlQWNPMlV2Ym1CQT09--658aa0aa5b88fb224c8d574fd7643b3d23cd7efb; _ga_7EFE8PPXQV=GS1.1.1655259464.7.1.1655260298.0")
		req.Header.Set("Dnt", "1")
		req.Header.Set("Origin", "https://rapidshipltl.mycarriertms.com")
		req.Header.Set("Pragma", "no-cache")
		req.Header.Set("Referer", "https://rapidshipltl.mycarriertms.com/")
		req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
		req.Header.Set("Sec-Fetch-Dest", "empty")
		req.Header.Set("Sec-Fetch-Mode", "cors")
		req.Header.Set("Sec-Fetch-Site", "cross-site")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInJvbGUiOlsiTXlDYXJyaWVyX1NoaXAiLCJNeUNhcnJpZXJfTWFuYWdlQ29tcGFueSIsIkNhbkFjY2Vzc0hvbWVQYWdlIiwiQ2FuVXNlUXVvdGVQYWdlIiwiQ2FuVXNlU2hpcG1lbnRMaXN0UGFnZSIsIkNhblVzZUFkZHJlc3NCb29rUGFnZSIsIkNhblVzZUxvY2F0aW9uUGFnZSIsIkNhblVzZVByb2R1Y3RQYWdlIiwiQ2FuVXNlQmlsbGluZ0FkZHJlc3NQYWdlIiwiQ2FuVXNlQ2FsZW5kYXJIb21lUGFnZSIsIkNhblVzZVByb2ZpbGVDb25maWd1cmF0aW9uUGFnZSIsIkNhblVzZUN1c3RvbWVyVXNlcnNQYWdlIiwiQ2FuVXNlQ29udGFjdFN1cHBvcnQiLCJDYW5Vc2VTZWxlY3RMb2NhdGlvbkNhcnJpZXJQYWdlIiwiQ2FuVXNlUGFzdFF1b3RlIiwiQ2FuVXNlQnVsa1VwbG9hZEFkZHJlc3NQYWdlIiwiQ2FuVXNlUXVpY2tCb2xQYWdlIiwiQ2FuVXNlQnVsa0VkaXRBZGRyZXNzUGFnZSIsIkNhblVzZVRydWNrbG9hZEFwaSIsIkNhblVzZUN1c3RvbWVyQ2FycmllclJlc291cmNlc1BhZ2UiLCJDYW5Vc2VCdWxrVXBsb2FkT3JkZXJQYWdlIiwiQ2FuVXNlSW52b2ljZVBhZ2UiLCJDdXN0b21lckFkbWluaXN0cmF0b3IiXSwibmJmIjoxNjU2NDY1Mjg1LCJleHAiOjE2NTY0Njg4ODUsImlhdCI6MTY1NjQ2NTI4NX0.AsjCKRG4Yy8Tg4WiL4LAOpkhHZMr5p_X259xORchSbI")
		req.Header.Set("Environment", "PROD")
		req.Header.Set("Notificationresponseurl", "https://inf-prod-signalrhandler.azurewebsites.net/")
		req.Header.Set("Orderprocessingurl", "https://app-orderprocessing-prod-bulkapi.azurewebsites.net/")
		req.Header.Set("Request-Context", "appId=cid-v1:cf6939be-b751-46bb-b434-d14afeb50826")
		req.Header.Set("Request-Id", "|5aa3774e40db43f987cc846fe76d4c1b.0f3ce2ebcfa3435e")
		req.Header.Set("Timezone", "420")
		req.Header.Set("Traceparent", "00-5aa3774e40db43f987cc846fe76d4c1b-0f3ce2ebcfa3435e-01")
		req.Header.Set("Truckloadfunctionurl", "https://app-truckload-prod-api.azurewebsites.net/")
		req.Header.Set("Access-Control-Request-Headers", "content-type,sdk-context")
		req.Header.Set("Access-Control-Request-Method", "POST")
		req.Header.Set("Sdk-Context", "appId")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// handle err
			fmt.Println(err)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			// handle err
			fmt.Println(err)
		}
		quotes := &GetAllQuotes{}
		err = json.Unmarshal(bodyBytes, quotes)
		if err != nil {
			// handle err
			fmt.Println(err)
		}
		for _, i := range quotes.SourceList {
			id := i.SavedQuoteID
			body := []byte(`{}`)
			bodys := bytes.NewReader(body)
			req, err := http.NewRequest("POST", fmt.Sprintf("https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/SaveQuote/RemoveQuote?savedQuoteId=%s", id), bodys)
			if err != nil {
				// handle err
			}
			req.Header.Set("Accept", "*/*")
			req.Header.Set("Accept-Language", "en-US,en;q=0.9")
			req.Header.Set("Cache-Control", "no-cache")
			req.Header.Set("Cookie", "_gcl_au=1.1.1739013701.1652374336; _gid=GA1.2.251636717.1655151308; _gat_UA-114313627-1=1; _ga=GA1.1.1423264958.1654897822; __hstc=192659957.95ae10d589bf6987cd16bd998ade8017.1655260289021.1655260289021.1655260289021.1; hubspotutk=95ae10d589bf6987cd16bd998ade8017; __hssrc=1; __hssc=192659957.1.1655260289021; intercom-id-c9oc6fab=574a48b8-0cb5-4f39-9d33-55eb122ef365; intercom-session-c9oc6fab=KzhDU2V5WHFtY2NWNHJjOGpGZ25oanlvMHcxbVlVenpDMFJTUUZSSVJ1eXNhYWFzMWdidmdxMHR2WEt0Q0V3cS0tbldlQzBXT000NWJlQWNPMlV2Ym1CQT09--658aa0aa5b88fb224c8d574fd7643b3d23cd7efb; _ga_7EFE8PPXQV=GS1.1.1655259464.7.1.1655260298.0")
			req.Header.Set("Dnt", "1")
			req.Header.Set("Origin", "https://rapidshipltl.mycarriertms.com")
			req.Header.Set("Pragma", "no-cache")
			req.Header.Set("Referer", "https://rapidshipltl.mycarriertms.com/")
			req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
			req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
			req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
			req.Header.Set("Sec-Fetch-Dest", "empty")
			req.Header.Set("Sec-Fetch-Mode", "cors")
			req.Header.Set("Sec-Fetch-Site", "cross-site")
			req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.61 Safari/537.36")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInJvbGUiOlsiTXlDYXJyaWVyX1NoaXAiLCJNeUNhcnJpZXJfTWFuYWdlQ29tcGFueSIsIkNhbkFjY2Vzc0hvbWVQYWdlIiwiQ2FuVXNlUXVvdGVQYWdlIiwiQ2FuVXNlU2hpcG1lbnRMaXN0UGFnZSIsIkNhblVzZUFkZHJlc3NCb29rUGFnZSIsIkNhblVzZUxvY2F0aW9uUGFnZSIsIkNhblVzZVByb2R1Y3RQYWdlIiwiQ2FuVXNlQmlsbGluZ0FkZHJlc3NQYWdlIiwiQ2FuVXNlQ2FsZW5kYXJIb21lUGFnZSIsIkNhblVzZVByb2ZpbGVDb25maWd1cmF0aW9uUGFnZSIsIkNhblVzZUN1c3RvbWVyVXNlcnNQYWdlIiwiQ2FuVXNlQ29udGFjdFN1cHBvcnQiLCJDYW5Vc2VTZWxlY3RMb2NhdGlvbkNhcnJpZXJQYWdlIiwiQ2FuVXNlUGFzdFF1b3RlIiwiQ2FuVXNlQnVsa1VwbG9hZEFkZHJlc3NQYWdlIiwiQ2FuVXNlUXVpY2tCb2xQYWdlIiwiQ2FuVXNlQnVsa0VkaXRBZGRyZXNzUGFnZSIsIkNhblVzZVRydWNrbG9hZEFwaSIsIkNhblVzZUN1c3RvbWVyQ2FycmllclJlc291cmNlc1BhZ2UiLCJDYW5Vc2VCdWxrVXBsb2FkT3JkZXJQYWdlIiwiQ2FuVXNlSW52b2ljZVBhZ2UiLCJDdXN0b21lckFkbWluaXN0cmF0b3IiXSwibmJmIjoxNjU2NDY1Mjg1LCJleHAiOjE2NTY0Njg4ODUsImlhdCI6MTY1NjQ2NTI4NX0.AsjCKRG4Yy8Tg4WiL4LAOpkhHZMr5p_X259xORchSbI")
			req.Header.Set("Environment", "PROD")
			req.Header.Set("Notificationresponseurl", "https://inf-prod-signalrhandler.azurewebsites.net/")
			req.Header.Set("Orderprocessingurl", "https://app-orderprocessing-prod-bulkapi.azurewebsites.net/")
			req.Header.Set("Request-Context", "appId=cid-v1:cf6939be-b751-46bb-b434-d14afeb50826")
			req.Header.Set("Request-Id", "|5aa3774e40db43f987cc846fe76d4c1b.0f3ce2ebcfa3435e")
			req.Header.Set("Timezone", "420")
			req.Header.Set("Traceparent", "00-5aa3774e40db43f987cc846fe76d4c1b-0f3ce2ebcfa3435e-01")
			req.Header.Set("Truckloadfunctionurl", "https://app-truckload-prod-api.azurewebsites.net/")
			req.Header.Set("Access-Control-Request-Headers", "content-type,sdk-context")
			req.Header.Set("Access-Control-Request-Method", "POST")
			req.Header.Set("Sdk-Context", "appId")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				// handle err
				fmt.Println(err)
			}

			bodyBytes, err := io.ReadAll(resp.Body)
			fmt.Println(string(bodyBytes))
			if err != nil {
				// handle err
				fmt.Println(err)
			}
			resp.Body.Close()
		}

	}

}
