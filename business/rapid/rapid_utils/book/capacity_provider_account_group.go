package book

import (
	"github.com/ramsfords/backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GetCapacityProviderAccountGroup(bid *v1.Bid) models.CapacityProviderAccountGroup {
	if bid.CarrierCode == "RDFS" {
		return models.CapacityProviderAccountGroup{
			Code: "RamsFordinc_PROD_91762_1039341_202202111959",
			Accounts: []models.Account{
				{
					Code: "RDFS",
				},
			},
		}
	}
	if bid.CarrierCode == "XPOL" {
		return models.CapacityProviderAccountGroup{
			Code: "RamsFordinc_PROD_91762_1039341_202202111959",
			Accounts: []models.Account{
				{
					Code: "XPOL",
				},
			},
		}
	}
	if bid.CarrierCode == "CLNI" {
		return models.CapacityProviderAccountGroup{
			Code: "RamsFordinc_PROD_91762_1039341_202202111959",
			Accounts: []models.Account{
				{
					Code: "CLNI",
				},
			},
		}
	}
	return models.CapacityProviderAccountGroup{}

}
