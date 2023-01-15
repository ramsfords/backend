package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func FixCommodityDeliveryAccessorial(commodities []*v1.Commodity, delivery *models.DestinationShippingDetails) {
	for _, j := range commodities {
		if j.SortAndSegregate {
			delivery.Address.AddressAccessorials = append(delivery.Address.AddressAccessorials, SortAndSegregateDelivery())
			return
		}
	}

}
