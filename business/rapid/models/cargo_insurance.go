package models

type AdminCargoInsuranceInfo struct {
	IntegrationFee    int     `json:"integrationFee,omitempty"  dynamodbav:""`
	LtlMinimumPremium int     `json:"ltlMinimumPremium,omitempty"  dynamodbav:""`
	PremiumRate       float64 `json:"premiumRate,omitempty"  dynamodbav:""`
}

type CargoInsuranceQuoteInfo struct {
	IsAllowedPromptForCargoValue        bool                    `json:"isAllowedPromptForCargoValue,omitempty"  dynamodbav:"isAllowedPromptForCargoValue"`
	MaxCargoValue                       int                     `json:"maxCargoValue,omitempty"  dynamodbav:"maxCargoValue"`
	AdminCargoInsuranceInfo             AdminCargoInsuranceInfo `json:"adminCargoInsuranceInfo,omitempty"  dynamodbav:"adminCargoInsuranceInfo"`
	CustomerIntegrationFee              int                     `json:"customerIntegrationFee,omitempty"  dynamodbav:"customerIntegrationFee"`
	IsInsureShipmentsEnabledByCarrier   bool                    `json:"isInsureShipmentsEnabledByCarrier,omitempty"  dynamodbav:"isInsureShipmentsEnabledByCarrier"`
	IsSaiaCustomer                      bool                    `json:"isSaiaCustomer,omitempty"  dynamodbav:"isSaiaCustomer"`
	CarrierIntegrationFee               int                     `json:"carrierIntegrationFee,omitempty"  dynamodbav:"carrierIntegrationFee"`
	CarrierInsuranceCoeffecientCoverage int                     `json:"carrierInsuranceCoeffecientCoverage,omitempty"  dynamodbav:"carrierInsuranceCoeffecientCoverage"`
}
