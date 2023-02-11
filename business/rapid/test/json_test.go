package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/goccy/go-json"
	"github.com/ramsfords/backend/business/rapid/models"
)

type Test struct {
	Name string `json:"name"`
}
type AnotherTest struct {
	Name string `json:"name"`
}

func Test_Json(t *testing.T) {
	quoteRateStanderd := []byte(`{
		"carrierCode": "RDFS",
		"carrierCodeAdditional": "RDFS",
		"carrierName": "Roadrunner Freight",
		"networkPartnerId": null,
		"isGuaranty": false,
		"carrierDeliveryDate": "06/23",
		"carrierDeliveryTime": null,
		"estimateDeliveryDate": "2022-06-23T00:00:00",
		"deliveryTime": null,
		"capacityProviderQuoteNumber": "4478445",
		"truckloadIconLogo": null,
		"capacityProviderAccountGroup": {
			"code": "RamsFordinc_PROD_91762_1039341_202202111959",
			"accounts": [{
				"code": "RDFS"
			}]
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
		"quoteNumber": "4478445",
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
		"rateQuoteDetails": [{
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
		"apiOutageMessage": "Roadrunner Freight's APIs are currently offline. We have reached out and are working with their team to resolve this. We'll send a follow up message once they are back online.",
		"transitTime": 4
	}`)
	standard1 := &models.Standard{}
	err := json.Unmarshal(quoteRateStanderd, standard1)

	confirmStandard := []byte(`{
		"capacityProviderAccountGroup": {
			"code": "RamsFordinc_PROD_91762_1039341_202202111959",
			"accounts": [
				{
					"code": "RDFS"
				}
			]
		},
		"capacityProviderQuoteNumber": "4478445",
		"carrierCode": "RDFS",
		"carrierCodeAdditional": "RDFS",
		"handlingUnitVolume": 83.33333333333334,
		"handlingUnitDensity": 9.6,
		"handlingUnitTotal": 1,
		"totalCost": 293.79,
		"estimateDeliveryDate": "06/23/2022",
		"carrierName": "Roadrunner Freight",
		"serviceName": null,
		"handlingUnitTotalPackages": 1,
		"totalShipmentWeight": 800,
		"iconLogo": "https://content.mycarriertms.com/carriers/651af2b5-638a-463e-918d-9dedb42da42b.png",
		"carrierDeliveryDate": null,
		"isGuaranty": false,
		"freightCharge": 1,
		"bolText": null,
		"quoteId": 7444451,
		"originLocationId": 1039341,
		"destinationLocationId": 0,
		"billingAddressId": 1039341,
		"carrierId": 1070,
		"laneType": 0,
		"transitTime": 4,
		"shipmentId": null,
		"serviceLevelCode": "STD",
		"specialInstruction": null,
		"shipmentPriceDetails": [
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
		"serviceType": 1,
		"linearFeet": null,
		"unAcceptedAccessorials": null,
		"dispatchTypeId": null,
		"equipmentTypeId": null,
		"equipmentSize": null,
		"targetRate": null,
		"shipmentNote": null,
		"isHotLoad": false,
		"opportunityId": 7444451,
		"optionId": null,
		"truckloadAccessorials": null,
		"shipmentTruckloadId": null,
		"cargoValue": null,
		"isFromSavedQuote": false,
		"isVicsBolShipment": false,
		"originCode": null,
		"referenceId": null
	}`)
	fmt.Println(err)
	standard2 := &models.Standard{}
	err = json.Unmarshal(confirmStandard, standard2)
	standard3 := &models.ConfirmAndDispatch{}
	confirmAndDispatch := []byte(`{
		"carrierCode": "RDFS",
		"carrierCodeAdditional": "RDFS",
		"carrierName": "Roadrunner Freight",
		"networkPartnerId": null,
		"isGuaranty": false,
		"carrierDeliveryDate": "06/23",
		"carrierDeliveryTime": null,
		"estimateDeliveryDate": "2022-06-23T00:00:00",
		"deliveryTime": null,
		"capacityProviderQuoteNumber": "4478445",
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
		"quoteNumber": "4478445",
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
		"apiOutageMessage": "Roadrunner Freight's APIs are currently offline. We have reached out and are working with their team to resolve this. We'll send a follow up message once they are back online.",
		"transitTime": 4
	}`)
	err = json.Unmarshal(confirmAndDispatch, standard3)
	jsonValue, err := json.Marshal(standard3)
	fmt.Println(string(jsonValue))
	fmt.Println(standard3)
	fmt.Println(err)
	fmt.Println(standard1)
	fmt.Println(standard2)
}

func Test_Nson(t *testing.T) {
	now, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		fmt.Println(err)
	}
	pickupDate := now.Format("01/02/2006")
	fmt.Println(pickupDate)
}
