package models

type Standard struct {
	CarrierCode                  *string                       `json:"carrierCode"`
	CarrierCodeAdditional        *string                       `json:"carrierCodeAdditional"`
	CarrierName                  *string                       `json:"carrierName"`
	NetworkPartnerID             *string                       `json:"networkPartnerId,omitempty"`
	IsGuaranty                   bool                          `json:"isGuaranty"`
	CarrierDeliveryDate          *string                       `json:"carrierDeliveryDate"`
	CarrierDeliveryTime          *string                       `json:"carrierDeliveryTime,omitempty"`
	EstimateDeliveryDate         *string                       `json:"estimateDeliveryDate"`
	DeliveryTime                 *string                       `json:"deliveryTime,omitempty"`
	CapacityProviderQuoteNumber  *string                       `json:"capacityProviderQuoteNumber"`
	TruckloadIconLogo            *string                       `json:"truckloadIconLogo,omitempty"`
	CarrierRequiredField         *CarrierRequiredField         `json:"carrierRequiredField,omitempty"`
	CarrierAPIs                  *CarrierAPIs                  `json:"carrierAPIs,omitempty"`
	OnTimeRisk                   *float64                      `json:"onTimeRisk,omitempty"`
	Logo                         *string                       `json:"logo,omitempty"`
	LargeLogo                    *string                       `json:"largeLogo,omitempty"`
	QuoteNumber                  *string                       `json:"quoteNumber,omitempty"`
	InfoMessage                  *string                       `json:"infoMessage,omitempty"`
	CurrencyCode                 *string                       `json:"currencyCode,omitempty"`
	Total                        float64                       `json:"total,omitempty"`
	ServiceName                  *string                       `json:"serviceName"`
	BolText                      *string                       `json:"bolText"`
	CarrierID                    int                           `json:"carrierId"`
	IsTermsAndConditionsEnabled  *bool                         `json:"isTermsAndConditionsEnabled,omitempty"`
	TermAndCondition             interface{}                   `json:"termAndCondition,omitempty"`
	LaneType                     int                           `json:"laneType"`
	LaneTypeName                 *string                       `json:"laneTypeName,omitempty"`
	CapacityProviderAccountGroup *CapacityProviderAccountGroup `json:"capacityProviderAccountGroup"`
	HandlingUnitVolume           float64                       `json:"handlingUnitVolume"`
	HandlingUnitDensity          float64                       `json:"handlingUnitDensity"`
	HandlingUnitTotal            int                           `json:"handlingUnitTotal"`
	TotalCost                    float64                       `json:"totalCost"`
	HandlingUnitTotalPackages    int                           `json:"handlingUnitTotalPackages"`
	TotalShipmentWeight          int                           `json:"totalShipmentWeight"`
	IconLogo                     *string                       `json:"iconLogo"`
	FreightCharge                int                           `json:"freightCharge"`
	QuoteID                      int                           `json:"quoteId"`
	OriginLocationID             int                           `json:"originLocationId"`
	DestinationLocationID        int                           `json:"destinationLocationId"`
	BillingAddressID             int                           `json:"billingAddressId"`
	TransitTime                  int                           `json:"transitTime"`
	ShipmentID                   *string                       `json:"shipmentId"`
	ServiceLevelCode             *string                       `json:"serviceLevelCode"`
	SpecialInstruction           *string                       `json:"specialInstruction"`
	ShipmentPriceDetails         *[]RateQuoteDetails           `json:"shipmentPriceDetails,omitempty"`
	RateQuoteDetails             *[]RateQuoteDetails           `json:"rateQuoteDetails,omitempty"`
	IsSelectedCarrier            *bool                         `json:"isSelectedCarrier,omitempty"`
	TransitDays                  *int                          `json:"transitDays,omitempty"`
	DispatchTypeID               *string                       `json:"dispatchTypeId"`
	EffectiveDate                *string                       `json:"effectiveDate,omitempty"`
	UnAcceptedAccessorials       []*AddressAccessorial         `json:"unAcceptedAccessorials"`
	CarrierVolumeQuoteNum        interface{}                   `json:"carrierVolumeQuoteNum,omitempty"`
	ContractID                   *string                       `json:"contractId,omitempty"`
	ExpirationDate               *string                       `json:"expirationDate,omitempty"`
	TruckloadAvailabilityDate    *string                       `json:"truckloadAvailabilityDate,omitempty"`
	PriorityMessageHeading       *string                       `json:"priorityMessageHeading,omitempty"`
	PriorityMessages             interface{}                   `json:"priorityMessages,omitempty"`
	DisplayAPIWarning            *bool                         `json:"displayAPIWarning,omitempty"`
	IsHasITMContract             bool                          `json:"isHasITMContract,omitempty"`
	APIOutageMessage             *string                       `json:"apiOutageMessage,omitempty"`
	ServiceType                  int                           `json:"serviceType"`
	LinearFeet                   interface{}                   `json:"linearFeet"`
	EquipmentTypeID              interface{}                   `json:"equipmentTypeId"`
	EquipmentSize                interface{}                   `json:"equipmentSize"`
	TargetRate                   interface{}                   `json:"targetRate"`
	ShipmentNote                 *string                       `json:"shipmentNote"`
	IsHotLoad                    bool                          `json:"isHotLoad"`
	OpportunityID                int                           `json:"opportunityId"`
	OptionID                     *string                       `json:"optionId"`
	TruckloadAccessorials        []*AddressAccessorial         `json:"truckloadAccessorials"`
	ShipmentTruckloadID          *string                       `json:"shipmentTruckloadId"`
	CargoValue                   interface{}                   `json:"cargoValue"`
	IsFromSavedQuote             bool                          `json:"isFromSavedQuote"`
	IsVicsBolShipment            bool                          `json:"isVicsBolShipment"`
	OriginCode                   *string                       `json:"originCode"`
	ReferenceID                  *string                       `json:"referenceId"`

	// new added field
	ShipperDetails         *PartyDetails           `json:"shipperDetails,omitempty"`
	ConsigneeDetails       *PartyDetails           `json:"consigneeDetails,omitempty"`
	BillingAddress         *Address                `json:"billingAddress,omitempty"`
	ReferenceNumberInfo    *ReferenceNumberInfo    `json:"referenceNumberInfo,omitempty"`
	EmergencyContactPerson *EmergencyContactPerson `json:"emergencyContactPerson,omitempty"`
	HandlingUnits          *[]ShipmentItem         `json:"handlingUnits,omitempty"`
	ShipmentItems          *[]ShipmentItem         `json:"shipmentItems,omitempty"`
}
