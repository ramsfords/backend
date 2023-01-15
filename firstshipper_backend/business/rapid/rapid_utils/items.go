package rapid_utils

import (
	"fmt"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/business/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeItems(shippingItems []v1.Commodity) []models.Commodity {
	commodities := []models.Commodity{}
	for _, j := range shippingItems {
		commodity := models.Commodity{
			Description:         &j.ShipmentDescription,
			Nmfc:                nil,
			NmfcSub:             nil,
			CommodityClass:      utils.StrPtr(GetClass(j)),
			Pieces:              int(j.Quantity),
			TotalWeight:         utils.StrPtr(fmt.Sprintf("%v", (j.Weight))),
			CurrentTotalWeight:  utils.StrPtr(fmt.Sprintf("%v", (j.Weight))),
			IsHazardous:         j.Hazardous,
			HazmatDetailInfo:    models.HazmatDetailInfo{},
			CustomerOrderNumber: nil,
			AdditionalInfo:      nil,
		}
		commodities = append(commodities, commodity)
	}
	return commodities
}
