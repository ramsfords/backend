package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func verifyBookCommodities(bookComodities []*v1.Commodity, quoteCommodity []*v1.Commodity) error {
	for _, j := range bookComodities {
		qc := GetComodityByIndex(int(j.Index), quoteCommodity)
		if qc.DimensionUom != j.DimensionUom {
			return errs.InvalidDimensionUOM
		}
		if qc.WeightUom != j.WeightUom {
			return errs.InvalidWeightUOM
		}
		if qc.Height != j.Height || qc.Length != j.Length || qc.Width != j.Width {
			return errs.ErrInvalidShipmentCommodity
		}
		if qc.ShipmentDescription != j.ShipmentDescription {
			return errs.InvalidShipmentDescription
		}
		if qc.Weight != j.Weight {
			return errs.InvalidWeight
		}

		if len(j.CommodityServices) != len(qc.CommodityServices) {
			return errs.InvalidCommodityServices
		}
		validServices := true
		for _, k := range j.CommodityServices {
			if !IncludesService(toInt32ArrayFromCommodityServices(qc.CommodityServices), int32(*k.Enum())) {
				validServices = false
			}
		}
		if !validServices {
			return errs.InvalidCommodityServices
		}
	}
	return nil
}
