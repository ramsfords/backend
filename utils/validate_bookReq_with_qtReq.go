package utils

import (
	"github.com/pkg/errors"
	v1 "github.com/ramsfords/types_gen/v1"
)

func validateBookRequestWithQuoteRequest(qtReq *v1.QuoteRequest, bkReq *v1.QuoteRequest) error {
	if len(qtReq.Commodities) != len(bkReq.Commodities) {
		return errors.New("commodities length not match")
	}
	if qtReq.TotalItems != bkReq.TotalItems {
		return errors.New("total items not match")
	}
	if qtReq.TotalWeight != bkReq.TotalWeight {
		return errors.New("total weight not match")
	}
	if qtReq.LocationServices.DeliveryAppointment != bkReq.LocationServices.DeliveryAppointment {
		return errors.New("delivery appointment not match")
	}
	if qtReq.LocationServices.PickupAppointment != bkReq.LocationServices.PickupAppointment {
		return errors.New("pickup appointment not match")
	}
	if qtReq.LocationServices.LiftGatePickup != bkReq.LocationServices.LiftGatePickup {
		return errors.New("pickup liftgate not match")
	}
	if qtReq.LocationServices.LiftGateDelivery != bkReq.LocationServices.LiftGateDelivery {
		return errors.New("delivery liftgate not match")
	}
	if qtReq.LocationServices.DeliveryLocationWithDock != bkReq.LocationServices.DeliveryLocationWithDock {
		return errors.New("location with dock not match")
	}
	if qtReq.LocationServices.PickupLocationWithDock != bkReq.LocationServices.PickupLocationWithDock {
		return errors.New("pickup location with dock not match")
	}
	if qtReq.LocationServices.DeliveryAppointment != bkReq.LocationServices.DeliveryAppointment {
		return errors.New("delivery appointment not match")
	}
	if qtReq.LocationServices.PickupAppointment != bkReq.LocationServices.PickupAppointment {
		return errors.New("pickup appointment not match")
	}
	if qtReq.LocationServices.DeliveryNotification != bkReq.LocationServices.DeliveryNotification {
		return errors.New("delivery notification not match")
	}
	if qtReq.LocationServices.PickupNotification != bkReq.LocationServices.PickupNotification {
		return errors.New("pickup notification not match")
	}

	for _, j := range qtReq.Commodities {
		bkItem := bkReq.Commodities[j.Index]
		if j.ShipmentDescription != bkItem.ShipmentDescription {
			return errors.New("shipment description not match")
		}
		if j.Quantity != bkItem.Quantity {
			return errors.New("quantity not match")
		}
		if j.Weight != bkItem.Weight {
			return errors.New("weight not match")
		}
		if j.DimensionUOM.CM != bkItem.DimensionUOM.CM {
			return errors.New("dimension uom cm not match")
		}
		if j.DimensionUOM.INCH != bkItem.DimensionUOM.INCH {
			return errors.New("dimension uom inch not match")
		}
		if j.WeightUOM.LB != bkItem.WeightUOM.LB {
			return errors.New("weight uom lb not match")
		}
		if j.PackageType != bkItem.PackageType {
			return errors.New("package type not match")
		}
		if j.Width != bkItem.Width {
			return errors.New("width not match")
		}
		if j.Height != bkItem.Height {
			return errors.New("height not match")
		}
		if j.Length != bkItem.Length {
			return errors.New("length not match")
		}
		if j.CommodityServices.Guaranteed != bkItem.CommodityServices.Guaranteed {
			return errors.New("guaranteed not match")
		}
		if j.CommodityServices.Hazardous != bkItem.CommodityServices.Hazardous {
			return errors.New("commodity hazardous not match")
		}
		if j.CommodityServices.ProtectFromFreeze != bkItem.CommodityServices.ProtectFromFreeze {
			return errors.New("commodity protect from freeze not match")
		}
		if j.CommodityServices.SortAndSegregate != bkItem.CommodityServices.SortAndSegregate {
			return errors.New("commodity sort and segregate from freeze not match")
		}
		if j.CommodityServices.Stackable != bkItem.CommodityServices.Stackable {
			return errors.New("commodity stackable from freeze not match")
		}
	}
	return nil
}
