package utils

import (
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookingCommodities(quoteReq *v1.QuoteRequest, quoteCommodity *v1.QuoteRequest) error {
	if len(quoteReq.Commodities) < 1 {
		return &errs.ErrInvalidShipmentCommodity
	}
	var totalWeight int32
	var totalItems int32

	for _, j := range quoteReq.Commodities {
		totalItems += int32(j.Quantity)
		totalWeight += int32(j.Weight)
		// if j.DimensionUOM != 0 {
		// 	return &errs.InvalidDimensionUOM
		// }
		// if j.WeightUOM != 0 {
		// 	return &errs.InvalidWeightUOM
		// }
		if j.PackageType == v1.PackageType_PACKAGENONE {
			return &errs.InvalidPackageType
		}
		if j.ShipmentDescription == "" || len(j.ShipmentDescription) < 2 {
			return &errs.InvalidShipmentDescription
		}
		if j.Height < 5 || j.Height > 100 || j.Length < 5 || j.Length > 100 || j.Width < 5 || j.Width > 100 {
			return &errs.InvalidDimension
		}
		if j.Weight < 0 || j.Weight > 10000 {
			return &errs.InvalidWeight
		}
	}
	quoteReq.TotalItems = totalItems
	quoteReq.TotalWeight = float32(totalWeight)
	err := verifyBookCommodities(quoteReq.Commodities, quoteCommodity.Commodities)
	if err != nil {
		return &errs.InvalidMismatchBookingAndQuoteCommodity
	}
	return nil
}
