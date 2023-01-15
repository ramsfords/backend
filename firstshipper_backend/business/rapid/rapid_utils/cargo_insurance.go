package rapid_utils

import "github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"

func GetNewCargoInsuranceQuoteInfo() models.CargoInsuranceQuoteInfo {
	return models.CargoInsuranceQuoteInfo{
		IsAllowedPromptForCargoValue: true,
		MaxCargoValue:                240000,
		AdminCargoInsuranceInfo: models.AdminCargoInsuranceInfo{
			IntegrationFee:    30,
			LtlMinimumPremium: 7,
			PremiumRate:       0.003,
		},
	}
}
