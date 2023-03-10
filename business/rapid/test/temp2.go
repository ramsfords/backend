package test

type AutoGeneratedss struct {
	CarrierCode                 string      `json:"carrierCode"`
	CarrierCodeAdditional       string      `json:"carrierCodeAdditional"`
	CarrierName                 string      `json:"carrierName"`
	NetworkPartnerID            interface{} `json:"networkPartnerId"`
	IsGuaranty                  bool        `json:"isGuaranty"`
	CarrierDeliveryDate         string      `json:"carrierDeliveryDate"`
	CarrierDeliveryTime         interface{} `json:"carrierDeliveryTime"`
	EstimateDeliveryDate        string      `json:"estimateDeliveryDate"`
	DeliveryTime                interface{} `json:"deliveryTime"`
	CapacityProviderQuoteNumber string      `json:"capacityProviderQuoteNumber"`
	TruckloadIconLogo           interface{} `json:"truckloadIconLogo"`
	CarrierRequiredField        struct {
		CarrierRequiredFieldID int  `json:"carrierRequiredFieldId"`
		ShipperEmailAddress    bool `json:"shipperEmailAddress"`
		ConsigneeEmailAddress  bool `json:"consigneeEmailAddress"`
		ShipperLastName        bool `json:"shipperLastName"`
		ConsigneeLastName      bool `json:"consigneeLastName"`
		NmfcCode               bool `json:"nmfcCode"`
		ConsigneeTime          bool `json:"consigneeTime"`
	} `json:"carrierRequiredField"`
	CarrierAPIs struct {
		CarrierAPIID             int  `json:"carrierAPIId"`
		QuotingAPIEnabled        bool `json:"quotingAPIEnabled"`
		DispatchAPIEnabled       bool `json:"dispatchAPIEnabled"`
		TrackingAPIEnabled       bool `json:"trackingAPIEnabled"`
		DocumentImagesAPIEnabled bool `json:"documentImagesAPIEnabled"`
	} `json:"carrierAPIs"`
	OnTimeRisk                  float64     `json:"onTimeRisk"`
	Logo                        string      `json:"logo"`
	LargeLogo                   string      `json:"largeLogo"`
	QuoteNumber                 string      `json:"quoteNumber"`
	InfoMessage                 interface{} `json:"infoMessage"`
	CurrencyCode                string      `json:"currencyCode"`
	Total                       float64     `json:"total"`
	ServiceName                 interface{} `json:"serviceName"`
	BolText                     interface{} `json:"bolText"`
	CarrierID                   int         `json:"carrierId"`
	IsTermsAndConditionsEnabled bool        `json:"isTermsAndConditionsEnabled"`
	TermAndCondition            interface{} `json:"termAndCondition"`
	LaneType                    int         `json:"laneType"`
	LaneTypeName                string      `json:"laneTypeName"`

	/////
	CapacityProviderAccountGroup struct {
		Code     string `json:"code"`
		Accounts []struct {
			Code string `json:"code"`
		} `json:"accounts"`
	} `json:"capacityProviderAccountGroup"`
	HandlingUnitVolume        float64     `json:"handlingUnitVolume"`
	HandlingUnitDensity       float64     `json:"handlingUnitDensity"`
	HandlingUnitTotal         int         `json:"handlingUnitTotal"`
	TotalCost                 float64     `json:"totalCost"`
	HandlingUnitTotalPackages int         `json:"handlingUnitTotalPackages"`
	TotalShipmentWeight       int         `json:"totalShipmentWeight"`
	IconLogo                  string      `json:"iconLogo"`
	FreightCharge             int         `json:"freightCharge"`
	QuoteID                   int         `json:"quoteId"`
	OriginLocationID          int         `json:"originLocationId"`
	DestinationLocationID     int         `json:"destinationLocationId"`
	BillingAddressID          int         `json:"billingAddressId"`
	TransitTime               int         `json:"transitTime"`
	ShipmentID                interface{} `json:"shipmentId"`
	ServiceLevelCode          string      `json:"serviceLevelCode"`
	SpecialInstruction        interface{} `json:"specialInstruction"`
	ShipmentPriceDetails      []struct {
		Amount           float64     `json:"amount"`
		Rate             float64     `json:"rate"`
		ItemFreightClass interface{} `json:"itemFreightClass"`
		Code             string      `json:"code"`
		Description      string      `json:"description"`
	} `json:"shipmentPriceDetails"`
	RateQuoteDetails []struct {
		Amount           float64     `json:"amount"`
		Rate             float64     `json:"rate"`
		ItemFreightClass interface{} `json:"itemFreightClass"`
		Code             string      `json:"code"`
		Description      string      `json:"description"`
	} `json:"rateQuoteDetails"`
	IsSelectedCarrier         bool        `json:"isSelectedCarrier"`
	TransitDays               int         `json:"transitDays"`
	DispatchTypeID            interface{} `json:"dispatchTypeId"`
	EffectiveDate             interface{} `json:"effectiveDate"`
	UnAcceptedAccessorials    interface{} `json:"unAcceptedAccessorials"`
	CarrierVolumeQuoteNum     interface{} `json:"carrierVolumeQuoteNum"`
	ContractID                interface{} `json:"contractId"`
	ExpirationDate            interface{} `json:"expirationDate"`
	TruckloadAvailabilityDate interface{} `json:"truckloadAvailabilityDate"`
	PriorityMessageHeading    interface{} `json:"priorityMessageHeading"`
	PriorityMessages          interface{} `json:"priorityMessages"`
	DisplayAPIWarning         bool        `json:"displayAPIWarning"`
	IsHasITMContract          bool        `json:"isHasITMContract"`
	APIOutageMessage          string      `json:"apiOutageMessage"`
	ServiceType               int         `json:"serviceType"`
	LinearFeet                interface{} `json:"linearFeet"`
	EquipmentTypeID           interface{} `json:"equipmentTypeId"`
	EquipmentSize             interface{} `json:"equipmentSize"`
	TargetRate                interface{} `json:"targetRate"`
	ShipmentNote              interface{} `json:"shipmentNote"`
	IsHotLoad                 bool        `json:"isHotLoad"`
	OpportunityID             int         `json:"opportunityId"`
	OptionID                  interface{} `json:"optionId"`
	TruckloadAccessorials     interface{} `json:"truckloadAccessorials"`
	ShipmentTruckloadID       interface{} `json:"shipmentTruckloadId"`
	CargoValue                interface{} `json:"cargoValue"`
	IsFromSavedQuote          bool        `json:"isFromSavedQuote"`
	IsVicsBolShipment         bool        `json:"isVicsBolShipment"`
	OriginCode                interface{} `json:"originCode"`
	ReferenceID               interface{} `json:"referenceId"`
}
