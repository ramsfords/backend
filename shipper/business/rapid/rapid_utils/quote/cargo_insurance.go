package quote

import "github.com/ramsfords/backend/shipper/business/rapid/models"

func getNewCargoInsuranceQuoteInfo() *models.CargoInsuranceQuoteInfo {
	return &models.CargoInsuranceQuoteInfo{
		IsAllowedPromptForCargoValue: true,
		MaxCargoValue:                240000,
		AdminCargoInsuranceInfo: models.AdminCargoInsuranceInfo{
			IntegrationFee:    30,
			LtlMinimumPremium: 7,
			PremiumRate:       0.003,
		},
	}
}
