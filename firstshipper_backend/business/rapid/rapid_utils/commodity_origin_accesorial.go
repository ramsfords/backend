package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func FixCommodityOriginAccessorial(commodities []*v1.Commodity, origin *models.OriginShippingDetails) {
	for _, j := range commodities {
		if j.ProtectFromFreeze {
			origin.Address.AddressAccessorials = append(origin.Address.AddressAccessorials, ProtectFromFreeze())
			return
		}
	}

}
