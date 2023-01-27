package utils

import (
	"errors"

	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateQuoteCommodities(qtReq *v1.QuoteRequest) error {
	if len(qtReq.Commodities) < 1 {
		return errs.NewApiError(400, "items not valid", errors.New("invalid shipping items"))
	}
	var totalWeight int32
	var totalItems int32

	for _, j := range qtReq.Commodities {
		totalItems += int32(j.Quantity)
		totalWeight += int32(j.Weight)
		if (j.DimensionUOM.CM || !j.DimensionUOM.INCH) && (!j.DimensionUOM.CM || j.DimensionUOM.INCH) {
			return errs.InvalidDimensionUOM
		}
		if !j.WeightUOM.LB {
			return errs.InvalidWeightUOM
		}
		if j.PackageType == v1.PackageType_PACKAGENONE {
			return errs.InvalidPackageType
		}
		if len(j.ShipmentDescription) < 2 {
			return errs.InvalidShipmentDescription
		}
		if j.Height < 5 || j.Height > 76 || j.Length < 5 || j.Length > 76 || j.Width < 5 || j.Width > 76 {
			return errs.InvalidDimension
		}
		if j.Weight < 0 || j.Weight > 10000 {
			return errs.InvalidWeight
		}
	}
	qtReq.TotalItems = totalItems
	qtReq.TotalWeight = float32(totalWeight)
	return nil
}
