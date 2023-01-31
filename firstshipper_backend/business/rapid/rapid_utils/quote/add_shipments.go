package quote

import (
	"fmt"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/business/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func AddShipments(quoteRequest *v1.QuoteRequest, rapidQuoteRequest *models.QuoteDetails) {
	items := []*models.ShipmentItem{}
	for _, j := range quoteRequest.Commodities {
		newCommodity := &models.ShipmentItem{
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
	rapidQuoteRequest.ShipmentItems = items
}

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
			HazmatDetailInfo:    models.HazmatDetailInfo{},
			CustomerOrderNumber: nil,
			AdditionalInfo:      nil,
		}
		commodities = append(commodities, commodity)
	}
	return commodities
}
