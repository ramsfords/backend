package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeQuoteWithJson() v1.QuoteRequest {
	req := &v1.QuoteRequest{}
	err := json.Unmarshal([]byte(`{"origin":{"address":{"zip_code":"90013"},"contact":{},"location_services":[4,2,3]},"delivery":{"address":{"zip_code":"60126"},"contact":{},"location_services":[5,7,6]},"commodities":[{"length":48,"width":40,"height":75,"weight":1000,"dimension_uom":1,"quantity":1,"freight_class":4,"commodity_services":[3,0,1,4,2]}],"pickup_date":"2022/05/12"}`), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(req)
	return *req
}
func Test_Payload(t *testing.T) {
	// qt := MakeQuoteWithJson()
	// // fmt.Println(qt)
	// rapidQuote := rapid.MakeQuote(qt)
	// fmt.Println(rapidQuote)
	// payloadBytess, err := json.Marshal(rapidQuote)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(payloadBytess))
	payloadBytes := []byte(`{
		"billingAddress": {
		  "addressId": 1039341,
		  "addressAccessorials": []
		},
		"originShippingDetails": {
		  "pickupDate": "06/15/2022",
		  "address": {
			"addressId": 1039341,
			"postalCode": "91762",
			"commercialType": {
			  "accessorialId": 72,
			  "name": "Business"
			},
			"addressAccessorials": []
		  }
		},
		"destinationShippingDetails": {
		  "address": {
			"postalCode": "60126",
			"commercialType": {
			  "accessorialId": 72,
			  "name": "Business"
			},
			"addressAccessorials": []
		  }
		},
		"shipmentItems": [
		  {	
			"handlingUnitType": "PLT",
			"handlingUnitNumber": 1,
			"dimensions": {
			  "length": "48",
			  "width": "40",
			  "height": "75",
			  "currentLength": "48",
			  "currentWidth": "40",
			  "currentHeight": "75"
			},
			"stackable": false,
			"commodities": [
			  {
				"description": "novelties",
				"commodityClass": "70",
				"pieces": 1,
				"totalWeight": "800",
				"currentTotalWeight": "800",
				"hazmatDetailInfo": {
				  "unCode": "UN"
				}
			  }
			]
		  }
		],
		"freightCharge": 1,
		"referenceNumberInfo": {},
		"emergencyContactPerson": {
		  "name": "Surendra Kandel",
		  "phone": "7135162836"
		},
		"serviceType": 1,
		"isShowVLTLToggle": true,
		"isAutoEmailTrackingEnabled": true,
		"isManualDispatchSettingEnabled": true,
		"isExistConnectedCarriers": true,
		"quoteErrors": [],
		"unusedCommodities": [],
		"customerCarrierId": 1070,
		"truckloadCapacityProviders": [
		  "EMRG"
		],
		"totalWeight": "800.00",
		"quoteDate": "06/15/2022"
	  }`)
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/Quote/GetQuoteRates", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "dc.services.visualstudio.com")
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
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInJvbGUiOlsiTXlDYXJyaWVyX1NoaXAiLCJNeUNhcnJpZXJfTWFuYWdlQ29tcGFueSIsIkNhbkFjY2Vzc0hvbWVQYWdlIiwiQ2FuVXNlUXVvdGVQYWdlIiwiQ2FuVXNlU2hpcG1lbnRMaXN0UGFnZSIsIkNhblVzZUFkZHJlc3NCb29rUGFnZSIsIkNhblVzZUxvY2F0aW9uUGFnZSIsIkNhblVzZVByb2R1Y3RQYWdlIiwiQ2FuVXNlQmlsbGluZ0FkZHJlc3NQYWdlIiwiQ2FuVXNlQ2FsZW5kYXJIb21lUGFnZSIsIkNhblVzZVByb2ZpbGVDb25maWd1cmF0aW9uUGFnZSIsIkNhblVzZUN1c3RvbWVyVXNlcnNQYWdlIiwiQ2FuVXNlQ29udGFjdFN1cHBvcnQiLCJDYW5Vc2VTZWxlY3RMb2NhdGlvbkNhcnJpZXJQYWdlIiwiQ2FuVXNlUGFzdFF1b3RlIiwiQ2FuVXNlQnVsa1VwbG9hZEFkZHJlc3NQYWdlIiwiQ2FuVXNlUXVpY2tCb2xQYWdlIiwiQ2FuVXNlQnVsa0VkaXRBZGRyZXNzUGFnZSIsIkNhblVzZVRydWNrbG9hZEFwaSIsIkNhblVzZUN1c3RvbWVyQ2FycmllclJlc291cmNlc1BhZ2UiLCJDYW5Vc2VCdWxrVXBsb2FkT3JkZXJQYWdlIiwiQ2FuVXNlSW52b2ljZVBhZ2UiLCJDdXN0b21lckFkbWluaXN0cmF0b3IiXSwibmJmIjoxNjU5NDgwNjE4LCJleHAiOjE2NTk0ODQyMTgsImlhdCI6MTY1OTQ4MDYxOH0.uyMit8AAB6lGt1jDNGY7gmukMYRIXjyZnZk5h2er6Dk")
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
	fmt.Println("hell")
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	defer resp.Body.Close()
}

func TestTime(t *testing.T) {
	now, err := time.Parse(time.RFC3339, "2022-06-15T17:00:00.000Z")
	fmt.Println(now, err)
	fmt.Println(now.Format("01/02/2006"))
}

func Test_Save_Payload(t *testing.T) {
	// qt := MakeQuoteWithJson()
	// // fmt.Println(qt)
	// rapidQuote := rapid.MakeQuote(qt)
	// fmt.Println(rapidQuote)
	// payloadBytess, err := json.Marshal(rapidQuote)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(payloadBytess))
	payloadBytes := []byte(`{
		"quoteDetails": {
			"billingAddress": {
				"addressId": 1039341,
				"addressAccessorials": [],
			},
			"originShippingDetails": {
				"pickupDate": "06/16/2022",
				"address": {
					"addressId": 1039341,
					"postalCode": "91762",
					"commercialType": {
						"accessorialId": 72,
						"name": "Business",
					},
					"addressAccessorials": [],
				}
			},
			"destinationShippingDetails": {
				"address": {
					"addressId": 0,
					"city": "Elmhurst",
					"postalCode": "60126",
					"commercialType": {
						"accessorialId": 72,
						"name": "Business"
					},
					"addressAccessorials": []
				}
			},
			"shipmentItems": [
				{
					"handlingUnitType": "PLT",
					"handlingUnitNumber": 1,
					"dimensions": {
						"length": "48",
						"width": "40",
						"height": "75",
						"currentLength": "48",
						"currentWidth": "40",
						"currentHeight": "75"
					},
					"stackable": false,
					"commodities": [
						{
							"description": "novelties",
							"commodityClass": "70",
							"pieces": 1,
							"totalWeight": "800",
							"currentTotalWeight": "800",
							"isHazardous": false,
							"hazmatDetailInfo": {
								"unCode": "UN"
							}
						}
					]
				}
			],
			"freightCharge": 1,
			"shipmentId": null,
			"isInsuredShipment": false,
			"insuranceEditShipmentData": null,
			"carrierCode": null,
			"referenceNumberInfo": {},
			"emergencyContactPerson": {
				"name": "Surendra Kandel",
				"phone": "7135162836"
			},
			"serviceLevelCode": null,
			"isAdminQuote": false,
			"referenceId": null,
			"isFromSavedQuote": false,
			"isFromShipment": false,
			"specialInstructionDefaultText": null,
			"serviceType": 1,
			"isShowVLTLToggle": true,
			"linearFeet": null,
			"currentLinearFeet": null,
			"isVLTLResponse": false,
			"isAutoEmailTrackingEnabled": true,
			"isManualDispatchSettingEnabled": true,
			"isMetricMeasurementEnabled": false,
			"equipmentTypeId": null,
			"equipmentTypeName": null,
			"equipmentSize": null,
			"targetRate": null,
			"shipmentNote": null,
			"isHotLoad": false,
			"opportunityId": null,
			"optionId": null,
			"truckloadAccessorials": null,
			"quoteId": 7412196,
			"quoteOpportunityId": null,
			"shipmentTruckloadId": null,
			"isReRunShipment": false,
			"specialInstruction": null,
			"carriersNotInQuote": null,
			"quoteKey": null,
			"isExistConnectedCarriers": true,
			"isCargoInsuranceToogleOn": false,
			"cargoInsuranceQuoteInfo": {
				"isAllowedPromptForCargoValue": true,
				"maxCargoValue": 240000,
				"adminCargoInsuranceInfo": {
					"integrationFee": 30,
					"ltlMinimumPremium": 7,
					"premiumRate": 0.003
				},
				"customerIntegrationFee": 0,
				"isInsureShipmentsEnabledByCarrier": true,
				"isSaiaCustomer": false,
				"carrierIntegrationFee": 0,
				"carrierInsuranceCoeffecientCoverage": 0
			},
			"quoteErrors": [],
			"unusedCommodities": [],
			"customerCarrierId": 1070,
			"integratedLocationAccountNumber1": null,
			"timeCreated": null,
			"originCode": null,
			"truckloadCapacityProviders": [
				"EMRG"
			],
			"isVicsBolShipment": false,
			"totalWeight": "800.00",
			"currentEquipmentSize": null,
			"quoteDate": "06/15/2022"
		},
		"confirmAndDispatch": null,
		"quoteRate": {
			"isValid": true,
			"titleHeadData": {
				"accessorialServices": [],
				"originZip": "91762",
				"destinationZip": "60126",
				"weight": 800,
				"quoteKey": "5f57d277-e624-40d1-a98a-cffdd3452331",
				"classes": [
					"70"
				],
				"pickUpDate": "2022-06-16T00:00:00",
				"formatedPickUpDate": {
					"formatedDay": "16",
					"formatedWeekDay": "Thu",
					"formatedMonth": "06"
				}
			},
			"quoteId": 7412196,
			"errorMessages": [],
			"carriersNotInQuote": [],
			"dayDeliveries": [
				{
					"formatedGuaranteedDate": "2022-06-22T00:00:00",
					"guarantedDate": "2022-06-22T00:00:00",
					"isHoliday": false,
					"isWeekend": false,
					"holidayName": null,
					"isUnspecifiedTransitCarrier": false,
					"dayDeliveryFormated": {
						"formatedDay": "22",
						"formatedWeekDay": "Wed",
						"formatedMonth": "06"
					},
					"transitDays": 4,
					"standart": [
						{
							"carrierCode": "RDFS",
							"carrierCodeAdditional": "RDFS",
							"carrierName": "Roadrunner Freight",
							"networkPartnerId": null,
							"isGuaranty": false,
							"carrierDeliveryDate": "06/22",
							"carrierDeliveryTime": null,
							"estimateDeliveryDate": "2022-06-22T00:00:00",
							"deliveryTime": null,
							"capacityProviderQuoteNumber": "4348598",
							"truckloadIconLogo": null,
							"capacityProviderAccountGroup": {
								"code": "RamsFordinc_PROD_91762_1039341_202202111959",
								"accounts": [
									{
										"code": "RDFS"
									}
								]
							},
							"carrierRequiredField": {
								"carrierRequiredFieldId": 1070,
								"shipperEmailAddress": false,
								"consigneeEmailAddress": false,
								"shipperLastName": false,
								"consigneeLastName": false,
								"nmfcCode": false,
								"consigneeTime": false
							},
							"carrierAPIs": {
								"carrierAPIId": 1070,
								"quotingAPIEnabled": true,
								"dispatchAPIEnabled": true,
								"trackingAPIEnabled": true,
								"documentImagesAPIEnabled": true
							},
							"onTimeRisk": 91.2,
							"logo": "https://content.mycarriertms.com/carriers/d5700f5d-ae85-4119-b4c0-b0d14ad08da4.png",
							"largeLogo": "https://content.mycarriertms.com/carriers/651af2b5-638a-463e-918d-9dedb42da42b.png",
							"quoteNumber": "4348598",
							"infoMessage": null,
							"currencyCode": "0",
							"total": 293.79,
							"serviceName": null,
							"bolText": null,
							"carrierId": 1070,
							"isTermsAndConditionsEnabled": false,
							"termAndCondition": null,
							"laneType": 0,
							"laneTypeName": "DIRECT",
							"rateQuoteDetails": [
								{
									"amount": 2399.52,
									"rate": 299.94,
									"itemFreightClass": null,
									"code": "ITEM",
									"description": "Item Charge"
								},
								{
									"amount": 60,
									"rate": 0,
									"itemFreightClass": null,
									"code": "CARBFEE",
									"description": "California Compliance Charge"
								},
								{
									"amount": 36.82,
									"rate": 0,
									"itemFreightClass": null,
									"code": "LH",
									"description": "Linehaul Surcharge - 27.90%"
								},
								{
									"amount": -2267.55,
									"rate": -94.5,
									"itemFreightClass": null,
									"code": "DSC",
									"description": "Discount - 94.50%"
								},
								{
									"amount": 65,
									"rate": 49.25,
									"itemFreightClass": null,
									"code": "FSC",
									"description": "Fuel Surcharge - 49.25%"
								}
							],
							"serviceLevelCode": "STD",
							"isSelectedCarrier": false,
							"transitDays": 4,
							"dispatchTypeId": null,
							"effectiveDate": null,
							"unAcceptedAccessorials": null,
							"carrierVolumeQuoteNum": null,
							"contractId": null,
							"expirationDate": null,
							"truckloadAvailabilityDate": null,
							"priorityMessageHeading": null,
							"priorityMessages": null,
							"displayAPIWarning": false,
							"isHasITMContract": true,
							"apiOutageMessage": "Roadrunner Freight's APIs are currently offline. We have reached out and are working with their team to resolve this. We'll send a follow up message once they are back online."
						}
					],
					"guarantyAM": [],
					"guarantyPM": []
				},
				{
					"formatedGuaranteedDate": "0001-01-01T00:00:00",
					"guarantedDate": "2022-06-23T00:00:00",
					"isHoliday": false,
					"isWeekend": false,
					"holidayName": null,
					"isUnspecifiedTransitCarrier": false,
					"dayDeliveryFormated": {
						"formatedDay": "23",
						"formatedWeekDay": "Thu",
						"formatedMonth": "06"
					},
					"transitDays": 5,
					"standart": [],
					"guarantyAM": [],
					"guarantyPM": []
				},
				{
					"formatedGuaranteedDate": "0001-01-01T00:00:00",
					"guarantedDate": "2022-06-24T00:00:00",
					"isHoliday": false,
					"isWeekend": false,
					"holidayName": null,
					"isUnspecifiedTransitCarrier": false,
					"dayDeliveryFormated": {
						"formatedDay": "24",
						"formatedWeekDay": "Fri",
						"formatedMonth": "06"
					},
					"transitDays": 6,
					"standart": [],
					"guarantyAM": [],
					"guarantyPM": []
				},
				{
					"formatedGuaranteedDate": "0001-01-01T00:00:00",
					"guarantedDate": "2022-06-27T00:00:00",
					"isHoliday": false,
					"isWeekend": false,
					"holidayName": null,
					"isUnspecifiedTransitCarrier": true,
					"dayDeliveryFormated": {
						"formatedDay": "27",
						"formatedWeekDay": "Mon",
						"formatedMonth": "06"
					},
					"transitDays": 7,
					"standart": [],
					"guarantyAM": [],
					"guarantyPM": []
				}
			],
			"selectedCarrier": null,
			"selectedCarrierTransitTime": null
		},
		"step": 2,
		"savedQuoteId": null,
		"isFromSavedQuote": false,
		"pickupDate": "06/16/2022",
		"infoMessage": null,
		"isFavorite": false,
		"isSystemAdmin": false,
		"isKeepAdminChangesModeEnabled": false,
		"orderId": null,
		"quoteErrors": []
	}
	`)
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/Quote/SaveQuoteStep", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "dc.services.visualstudio.com")
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
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInJvbGUiOlsiTXlDYXJyaWVyX1NoaXAiLCJNeUNhcnJpZXJfTWFuYWdlQ29tcGFueSIsIkNhbkFjY2Vzc0hvbWVQYWdlIiwiQ2FuVXNlUXVvdGVQYWdlIiwiQ2FuVXNlU2hpcG1lbnRMaXN0UGFnZSIsIkNhblVzZUFkZHJlc3NCb29rUGFnZSIsIkNhblVzZUxvY2F0aW9uUGFnZSIsIkNhblVzZVByb2R1Y3RQYWdlIiwiQ2FuVXNlQmlsbGluZ0FkZHJlc3NQYWdlIiwiQ2FuVXNlQ2FsZW5kYXJIb21lUGFnZSIsIkNhblVzZVByb2ZpbGVDb25maWd1cmF0aW9uUGFnZSIsIkNhblVzZUN1c3RvbWVyVXNlcnNQYWdlIiwiQ2FuVXNlQ29udGFjdFN1cHBvcnQiLCJDYW5Vc2VTZWxlY3RMb2NhdGlvbkNhcnJpZXJQYWdlIiwiQ2FuVXNlUGFzdFF1b3RlIiwiQ2FuVXNlQnVsa1VwbG9hZEFkZHJlc3NQYWdlIiwiQ2FuVXNlUXVpY2tCb2xQYWdlIiwiQ2FuVXNlQnVsa0VkaXRBZGRyZXNzUGFnZSIsIkNhblVzZVRydWNrbG9hZEFwaSIsIkNhblVzZUN1c3RvbWVyQ2FycmllclJlc291cmNlc1BhZ2UiLCJDYW5Vc2VCdWxrVXBsb2FkT3JkZXJQYWdlIiwiQ2FuVXNlSW52b2ljZVBhZ2UiLCJDdXN0b21lckFkbWluaXN0cmF0b3IiXSwibmJmIjoxNjU1MzQwMjk2LCJleHAiOjE2NTUzNDM4OTYsImlhdCI6MTY1NTM0MDI5Nn0.GmH6WnyGZewf0TQXKDvDOmSO2jxRxFxi_CO7zc-wExA")
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
	fmt.Println("hell")
	bodyBytes, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyBytes))
	defer resp.Body.Close()
}
