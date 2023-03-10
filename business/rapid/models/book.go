package models

type Book struct {
	CapacityProviderAccountGroup CapacityProviderAccountGroup `json:"capacityProviderAccountGroup,omitempty"`
	CapacityProviderQuoteNumber  string                       `json:"capacityProviderQuoteNumber,omitempty"`
	CarrierCode                  string                       `json:"carrierCode,omitempty"`
	CarrierCodeAdditional        string                       `json:"carrierCodeAdditional,omitempty"`
	ShipperDetails               PartyDetails                 `json:"shipperDetails,omitempty"`
	ConsigneeDetails             PartyDetails                 `json:"consigneeDetails,omitempty"`
	BillingAddress               Address                      `json:"billingAddress,omitempty"`
	ReferenceNumberInfo          `json:"referenceNumberInfo,omitempty"`
	EmergencyContactPerson       EmergencyContactPerson `json:"emergencyContactPerson,omitempty"`
	HandlingUnits                []ShipmentItem         `json:"handlingUnits,omitempty"`
	HandlingUnitVolume           float64                `json:"handlingUnitVolume,omitempty"`
	HandlingUnitDensity          float64                `json:"handlingUnitDensity,omitempty"`
	HandlingUnitTotal            int                    `json:"handlingUnitTotal,omitempty"`
	TotalCost                    float64                `json:"totalCost,omitempty"`
	EstimateDeliveryDate         string                 `json:"estimateDeliveryDate,omitempty"`
	CarrierName                  string                 `json:"carrierName,omitempty"`
	ServiceName                  string                 `json:"serviceName,omitempty"`
	HandlingUnitTotalPackages    int                    `json:"handlingUnitTotalPackages,omitempty"`
	TotalShipmentWeight          int                    `json:"totalShipmentWeight,omitempty"`
	IconLogo                     string                 `json:"iconLogo,omitempty"`
	CarrierDeliveryDate          string                 `json:"carrierDeliveryDate,omitempty"`
	IsGuaranty                   bool                   `json:"isGuaranty,omitempty"`
	FreightCharge                int                    `json:"freightCharge,omitempty"`
	BolText                      string                 `json:"bolText,omitempty"`
	QuoteID                      int                    `json:"quoteId,omitempty"`
	OriginLocationID             int                    `json:"originLocationId,omitempty"`
	DestinationLocationID        int                    `json:"destinationLocationId,omitempty"`
	BillingAddressID             int                    `json:"billingAddressId,omitempty"`
	BookedDate                   string                 `json:"bookedDate,omitempty"`
	CarrierID                    int                    `json:"carrierId,omitempty"`
	LaneType                     int                    `json:"laneType,omitempty"`
	TransitTime                  int                    `json:"transitTime,omitempty"`
	ShipmentID                   string                 `json:"shipmentId,omitempty"`
	ServiceLevelCode             string                 `json:"serviceLevelCode,omitempty"`
	SpecialInstruction           string                 `json:"specialInstruction,omitempty"`
	ShipmentPriceDetails         RateQuoteDetails       `json:"shipmentPriceDetails,omitempty"`
	SavedQuoteID                 string                 `json:"savedQuoteId,omitempty"`
	ServiceType                  int                    `json:"serviceType,omitempty"`
	LinearFeet                   string                 `json:"linearFeet,omitempty"`
	UnAcceptedAccessorials       []AddressAccessorial   `json:"unAcceptedAccessorials,omitempty"`
	DispatchTypeID               string                 `json:"dispatchTypeId,omitempty"`
	EquipmentTypeID              string                 `json:"equipmentTypeId,omitempty"`
	EquipmentSize                interface{}            `json:"equipmentSize,omitempty"`
	TargetRate                   interface{}            `json:"targetRate,omitempty"`
	ShipmentNote                 interface{}            `json:"shipmentNote,omitempty"`
	IsHotLoad                    bool                   `json:"isHotLoad,omitempty"`
	OpportunityID                int                    `json:"opportunityId,omitempty"`
	OptionID                     interface{}            `json:"optionId,omitempty"`
	TruckloadAccessorials        interface{}            `json:"truckloadAccessorials,omitempty"`
	ShipmentTruckloadID          interface{}            `json:"shipmentTruckloadId,omitempty"`
	CargoValue                   interface{}            `json:"cargoValue,omitempty"`
	IsFromSavedQuote             bool                   `json:"isFromSavedQuote,omitempty"`
	IsVicsBolShipment            bool                   `json:"isVicsBolShipment,omitempty"`
	OriginCode                   string                 `json:"originCode,omitempty"`
	ReferenceID                  string                 `json:"referenceId,omitempty"`
}
