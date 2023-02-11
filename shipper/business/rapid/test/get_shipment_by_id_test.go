package test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type ShipmentResponse struct {
	IsShipmentException            bool        `json:"isShipmentException"`
	IsShipmentManual               bool        `json:"isShipmentManual"`
	IsAbleFileAClaim               bool        `json:"isAbleFileAClaim"`
	ExceptionText                  interface{} `json:"exceptionText"`
	ShipmentStatusNumber           int         `json:"shipmentStatusNumber"`
	CarrierPRONumber               string      `json:"carrierPRONumber"`
	CustomerBOLNumber              string      `json:"customerBOLNumber"`
	PoNumber                       interface{} `json:"poNumber"`
	ReferenceNumber                interface{} `json:"referenceNumber"`
	GS1CompanyPrefix               interface{} `json:"gS1CompanyPrefix"`
	EstimatedDeliveryDate          string      `json:"estimatedDeliveryDate"`
	TruckloadRequestedDeliveryDate interface{} `json:"truckloadRequestedDeliveryDate"`
	PickupDate                     string      `json:"pickupDate"`
	ShipmentID                     int         `json:"shipmentId"`
	ActivityDateTime               string      `json:"activityDateTime"`
	SecurityKey                    string      `json:"securityKey"`
	CapacityQuoteNumber            string      `json:"capacityQuoteNumber"`
	PickupNumber                   interface{} `json:"pickupNumber"`
	CheckDigit                     string      `json:"checkDigit"`
	IsCanBeCanceled                bool        `json:"isCanBeCanceled"`
	ShipmentTruckloadID            interface{} `json:"shipmentTruckloadId"`
	FreightCharge                  int         `json:"freightCharge"`
	TenderID                       interface{} `json:"tenderId"`
	EquipmentTypeName              interface{} `json:"equipmentTypeName"`
	EquipmentSize                  interface{} `json:"equipmentSize"`
	TargetRate                     interface{} `json:"targetRate"`
	IsHotLoad                      bool        `json:"isHotLoad"`
	ShipmentNote                   interface{} `json:"shipmentNote"`
	HandlingUnitDensity            float64     `json:"handlingUnitDensity"`
	HandlingUnitTotal              int         `json:"handlingUnitTotal"`
	HandlingUnitTotalPackages      int         `json:"handlingUnitTotalPackages"`
	TotalShipmentWeight            int         `json:"totalShipmentWeight"`
	HandlingUnitVolume             float64     `json:"handlingUnitVolume"`
	ActualPickupDate               interface{} `json:"actualPickupDate"`
	UpdatedDeliveryDate            string      `json:"updatedDeliveryDate"`
	BolText                        interface{} `json:"bolText"`
	UnAcceptedAccessorials         interface{} `json:"unAcceptedAccessorials"`
	CustomerID                     int         `json:"customerId"`
	IsMetricMeasurementEnabled     bool        `json:"isMetricMeasurementEnabled"`
	IsCanceledByCarrier            bool        `json:"isCanceledByCarrier"`
	BookedDate                     string      `json:"bookedDate"`
	ServiceType                    int         `json:"serviceType"`
	CoverageValue                  interface{} `json:"coverageValue"`
	IsUnspecifiedTracking          bool        `json:"isUnspecifiedTracking"`
	CustomerCompanyName            string      `json:"customerCompanyName"`
	LastModifyUserName             string      `json:"lastModifyUserName"`
	AvailableDocuments             struct {
		Pod   bool `json:"pod"`
		Bol   bool `json:"bol"`
		Other bool `json:"other"`
	} `json:"availableDocuments"`
	CarrierInfo struct {
		CarrierID                 int         `json:"carrierId"`
		CarrierPartnerID          interface{} `json:"carrierPartnerId"`
		CarrierLogo               string      `json:"carrierLogo"`
		CarrierTruckloadLogo      interface{} `json:"carrierTruckloadLogo"`
		CarrierCode               string      `json:"carrierCode"`
		NetworkPartnerContactName interface{} `json:"networkPartnerContactName"`
		CarrierName               string      `json:"carrierName"`
		ShipmentServiceName       interface{} `json:"shipmentServiceName"`
		Total                     float64     `json:"total"`
		IsGuaranty                bool        `json:"isGuaranty"`
		FromProLength             int         `json:"fromProLength"`
		ToProLength               interface{} `json:"toProLength"`
		ProPrefixLength           int         `json:"proPrefixLength"`
		IsCheckDigit              bool        `json:"isCheckDigit"`
		IncludeCheckDigitTracking bool        `json:"includeCheckDigitTracking"`
		IsTrackingAPIEnabled      bool        `json:"isTrackingAPIEnabled"`
		HasITMContract            bool        `json:"hasITMContract"`
		CheckDigitModuleType      int         `json:"checkDigitModuleType"`
		CarrierContactInfo        struct {
			Phone          string      `json:"phone"`
			TruckloadPhone string      `json:"truckloadPhone"`
			Email          interface{} `json:"email"`
		} `json:"carrierContactInfo"`
	} `json:"carrierInfo"`
	PickupRemarks struct {
		StartTime              string      `json:"startTime"`
		EndTime                string      `json:"endTime"`
		InstructionNote        interface{} `json:"instructionNote"`
		Accessorials           interface{} `json:"accessorials"`
		UnAcceptedAccessorials interface{} `json:"unAcceptedAccessorials"`
	} `json:"pickupRemarks"`
	DeliveryRemarks struct {
		StartTime              interface{} `json:"startTime"`
		EndTime                interface{} `json:"endTime"`
		InstructionNote        interface{} `json:"instructionNote"`
		Accessorials           interface{} `json:"accessorials"`
		UnAcceptedAccessorials interface{} `json:"unAcceptedAccessorials"`
	} `json:"deliveryRemarks"`
	ShipperInfo struct {
		AddressID             int         `json:"addressId"`
		CompanyName           string      `json:"companyName"`
		StreetLine1           string      `json:"streetLine1"`
		StreetLine2           interface{} `json:"streetLine2"`
		City                  string      `json:"city"`
		State                 string      `json:"state"`
		Zip                   string      `json:"zip"`
		Country               string      `json:"country"`
		Name                  string      `json:"name"`
		Email                 string      `json:"email"`
		Phone                 string      `json:"phone"`
		HasSavedInAddressBook bool        `json:"hasSavedInAddressBook"`
	} `json:"shipperInfo"`
	ConsegneeInfo struct {
		AddressID             int    `json:"addressId"`
		CompanyName           string `json:"companyName"`
		StreetLine1           string `json:"streetLine1"`
		StreetLine2           string `json:"streetLine2"`
		City                  string `json:"city"`
		State                 string `json:"state"`
		Zip                   string `json:"zip"`
		Country               string `json:"country"`
		Name                  string `json:"name"`
		Email                 string `json:"email"`
		Phone                 string `json:"phone"`
		HasSavedInAddressBook bool   `json:"hasSavedInAddressBook"`
	} `json:"consegneeInfo"`
	RequesterInfo struct {
		AddressID             int         `json:"addressId"`
		CompanyName           interface{} `json:"companyName"`
		StreetLine1           interface{} `json:"streetLine1"`
		StreetLine2           interface{} `json:"streetLine2"`
		City                  interface{} `json:"city"`
		State                 interface{} `json:"state"`
		Zip                   interface{} `json:"zip"`
		Country               interface{} `json:"country"`
		Name                  interface{} `json:"name"`
		Email                 interface{} `json:"email"`
		Phone                 interface{} `json:"phone"`
		HasSavedInAddressBook bool        `json:"hasSavedInAddressBook"`
	} `json:"requesterInfo"`
	TrackingHistories []interface{} `json:"trackingHistories"`
	HandlingUnits     []struct {
		HandlingUnitNumber int  `json:"handlingUnitNumber"`
		Stackable          bool `json:"stackable"`
		Width              int  `json:"width"`
		Height             int  `json:"height"`
		Length             int  `json:"length"`
		HandlingUnitType   struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"handlingUnitType"`
		Commodities []struct {
			Description          string      `json:"description"`
			NmfcCode             interface{} `json:"nmfcCode"`
			NmfcSubCode          interface{} `json:"nmfcSubCode"`
			PiecesNumber         int         `json:"piecesNumber"`
			TotalWeight          int         `json:"totalWeight"`
			UnCode               string      `json:"unCode"`
			UnNumber             string      `json:"unNumber"`
			PropertyShippingName string      `json:"propertyShippingName"`
			HazardClass          interface{} `json:"hazardClass"`
			PackingGroup         string      `json:"packingGroup"`
			IsHazMat             bool        `json:"isHazMat"`
			FreightClass         struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"freightClass"`
			PackingType         interface{} `json:"packingType"`
			CustomerOrderNumber interface{} `json:"customerOrderNumber"`
			AdditionalInfo      interface{} `json:"additionalInfo"`
		} `json:"commodities"`
	} `json:"handlingUnits"`
	ShipmentPriceDetails []struct {
		Description string `json:"description"`
		Amount      int    `json:"amount"`
	} `json:"shipmentPriceDetails"`
	CargoInsuranceInfo struct {
		IsCanShipWithInshuredEdited           bool        `json:"isCanShipWithInshuredEdited"`
		IsInsuredShipment                     bool        `json:"isInsuredShipment"`
		TotalInsurancePaymentSum              interface{} `json:"totalInsurancePaymentSum"`
		InsurancePremium                      interface{} `json:"insurancePremium"`
		IntegrationFee                        interface{} `json:"integrationFee"`
		FalveyShipmentID                      interface{} `json:"falveyShipmentId"`
		CustomerCargoInsuranceID              int         `json:"customerCargoInsuranceId"`
		CertificateNumber                     interface{} `json:"certificateNumber"`
		InsuranceShipmentTypeID               int         `json:"insuranceShipmentTypeId"`
		IsCanBeInsured                        bool        `json:"isCanBeInsured"`
		IsInsuranceInvoiceEnabled             bool        `json:"isInsuranceInvoiceEnabled"`
		ShippingSumInsured                    int         `json:"shippingSumInsured"`
		IsAnyCustomerPaidInsuranceCertificate bool        `json:"isAnyCustomerPaidInsuranceCertificate"`
		IsCreditCardPaymentType               interface{} `json:"isCreditCardPaymentType"`
		CustomerIntegrationFee                int         `json:"customerIntegrationFee"`
	} `json:"cargoInsuranceInfo"`
	IsShowEditButtonOnInsurance bool        `json:"isShowEditButtonOnInsurance"`
	ClaimFiled                  bool        `json:"claimFiled"`
	TenderStatusID              interface{} `json:"tenderStatusId"`
	IsVicsBolShipment           bool        `json:"isVicsBolShipment"`
}

func Test_GetShipmentById(t *testing.T) {
	req, err := http.NewRequest("GET", "https://rapidshipltl.mycarriertms.com/MyCarrierAPI//api/Shipment/GetShipmentById?shipmentId=3438544", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Authority", "rapidshipltl.mycarriertms.com")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1bmlxdWVfbmFtZSI6ImthbmRlbHN1cmVuQGdtYWlsLmNvbSIsInJvbGUiOlsiTXlDYXJyaWVyX1NoaXAiLCJNeUNhcnJpZXJfTWFuYWdlQ29tcGFueSIsIkNhbkFjY2Vzc0hvbWVQYWdlIiwiQ2FuVXNlUXVvdGVQYWdlIiwiQ2FuVXNlU2hpcG1lbnRMaXN0UGFnZSIsIkNhblVzZUFkZHJlc3NCb29rUGFnZSIsIkNhblVzZUxvY2F0aW9uUGFnZSIsIkNhblVzZVByb2R1Y3RQYWdlIiwiQ2FuVXNlQmlsbGluZ0FkZHJlc3NQYWdlIiwiQ2FuVXNlQ2FsZW5kYXJIb21lUGFnZSIsIkNhblVzZVByb2ZpbGVDb25maWd1cmF0aW9uUGFnZSIsIkNhblVzZUN1c3RvbWVyVXNlcnNQYWdlIiwiQ2FuVXNlQ29udGFjdFN1cHBvcnQiLCJDYW5Vc2VTZWxlY3RMb2NhdGlvbkNhcnJpZXJQYWdlIiwiQ2FuVXNlUGFzdFF1b3RlIiwiQ2FuVXNlQnVsa1VwbG9hZEFkZHJlc3NQYWdlIiwiQ2FuVXNlUXVpY2tCb2xQYWdlIiwiQ2FuVXNlQnVsa0VkaXRBZGRyZXNzUGFnZSIsIkNhblVzZVRydWNrbG9hZEFwaSIsIkNhblVzZUN1c3RvbWVyQ2FycmllclJlc291cmNlc1BhZ2UiLCJDYW5Vc2VCdWxrVXBsb2FkT3JkZXJQYWdlIiwiQ2FuVXNlSW52b2ljZVBhZ2UiLCJDdXN0b21lckFkbWluaXN0cmF0b3IiXSwibmJmIjoxNjU2NDY5MDA5LCJleHAiOjE2NTY0NzI2MDksImlhdCI6MTY1NjQ2OTAwOX0.-HyPzk2R8z_B3URP5EWWKb9Ynj0j1TsyzGFFt7sKFEA")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "ai_user=Oi4VWkdRdayuaM8bsUTtLr|2022-05-12T16:52:15.909Z; _gcl_au=1.1.611107233.1655681811; ASPSESSIONIDSGBBABTD=NPLCKCCDJJODINFPGLADMIMB; ASPSESSIONIDSGBDSDSS=FABDOGMAPFILKLMJLDGKJLGL; ASPSESSIONIDQGBCCATD=JNHDFMLAOENHAOHFCOCFEGAB; _gid=GA1.2.419302534.1656369639; ASPSESSIONIDCERTDBQA=ENICOGMAFEILPOLEFDLHJPMD; ASPSESSIONIDQEBBACTD=BMCCANBBMNJOPJAPPGPCNLKD; ASPSESSIONIDCEDSACSB=LIPOJPGBBDAKNGPMOLIDLPJI; hubspotutk=e783a15325a2b1d4a32e6caf15455034; __hssrc=1; intercom-id-c9oc6fab=8476f2dc-fe2a-4ca5-ad9a-2e49fe216bba; ASPSESSIONIDSGCDDDQD=AHBDLLIBNMBGNJNIOMNBDMHC; ASPSESSIONIDCEAQABQD=MEFJLFMBGGFJIJIHOCBHELJB; __hstc=192659957.e783a15325a2b1d4a32e6caf15455034.1656434261965.1656461762245.1656465852461.5; ASPSESSIONIDQEBCBBTC=GELILFMBBOEPKLLFMDCHPGIA; _ga=GA1.1.2101301807.1655681811; __hssc=192659957.8.1656465852461; intercom-session-c9oc6fab=NlU3a2dBSGNIZHhqOGJ5SUtyVjJKbTI3UlhrVVlXeTdYTTk5MXc5eTBld0t3SzhZRy85WmQ2N08xZW5CQmRCMS0tOGxNM2dUUEJrUXZtaFJDb2FDODFMdz09--4b8973eaf8e78372a8e75048f15d04fbac6860c0; _ga_7EFE8PPXQV=GS1.1.1656465286.39.1.1656469009.0; ASPSESSIONIDCWSDADCC=IPFILFMBKJIIOANMGFJPPPIK; ai_session=hPZ+Oflk5bbeE8+NVqz5d4|1656461759986|1656469096979")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Environment", "PROD")
	req.Header.Set("Notificationresponseurl", "https://inf-prod-signalrhandler.azurewebsites.net/")
	req.Header.Set("Orderprocessingurl", "https://app-orderprocessing-prod-bulkapi.azurewebsites.net/")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://rapidshipltl.mycarriertms.com/customers/quote")
	req.Header.Set("Request-Context", "appId=cid-v1:cf6939be-b751-46bb-b434-d14afeb50826")
	req.Header.Set("Request-Id", "|48433c13f10f462dba0dafdd12a2d249.f7bf8b9b03f64605")
	req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"102\", \"Google Chrome\";v=\"102\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Timezone", "420")
	req.Header.Set("Traceparent", "00-48433c13f10f462dba0dafdd12a2d249-f7bf8b9b03f64605-01")
	req.Header.Set("Truckloadfunctionurl", "https://app-truckload-prod-api.azurewebsites.net/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	shipment := &ShipmentResponse{}
	bodyBytes, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, shipment)
	defer resp.Body.Close()
}
