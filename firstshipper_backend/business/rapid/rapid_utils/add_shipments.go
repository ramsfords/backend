package rapid_utils

import (
	"fmt"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/business/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func AddShipments(baseCommodity []*v1.Commodity) *[]models.ShipmentItem {
	items := []models.ShipmentItem{}
	for _, j := range baseCommodity {
		newCommodity := models.ShipmentItem{
			HandlingUnitType:   GetPackageType(*j),
			HandlingUnitNumber: int(j.Quantity),
			Dimensions: models.Dimensions{
				Length:        fmt.Sprintf("%v", int(j.Length)),
				Width:         fmt.Sprintf("%v", int(j.Width)),
				Height:        fmt.Sprintf("%v", int(j.Height)),
				CurrentLength: fmt.Sprintf("%v", int(j.Length)),
				CurrentWidth:  fmt.Sprintf("%v", int(j.Width)),
				CurrentHeight: fmt.Sprintf("%v", int(j.Height)),
			},
			Stackable: j.Stackable,
			Commodities: []models.Commodity{
				{
					Description:        utils.StrPtr(j.ShipmentDescription),
					CommodityClass:     utils.StrPtr(GetClass(*j)),
					Pieces:             int(j.Quantity),
					TotalWeight:        utils.StrPtr(fmt.Sprintf("%v", int(j.Weight))),
					CurrentTotalWeight: utils.StrPtr(fmt.Sprintf("%v", int(j.Weight))),
					HazmatDetailInfo: models.HazmatDetailInfo{
						UnCode: utils.StrPtr("UN"),
					},
					IsHazardous: false,
				},
			},
		}
		items = append(items, newCommodity)
	}
	return &items
}
