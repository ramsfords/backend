package book

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func ToV1DispatchResponse(model models.DispatchResponse) v1.DispatchResponse {
	res := v1.DispatchResponse{
		ShipmentID:           int32(model.ShipmentID),
		SecurityKey:          model.SecurityKey,
		PickupNumber:         model.PickupNumber,
		CarrierName:          model.CarrierName,
		CarrierPhone:         model.CarrierPhone,
		CarrierPRONumber:     model.CarrierPRONumber,
		HandlingUnitTotal:    model.HandlingUnitTotal,
		IsShipmentEdit:       model.IsShipmentEdit,
		IsShipmentManual:     model.IsShipmentManual,
		ServiceType:          int32(model.ServiceType),
		IsTrackingEmailSend:  model.IsTrackingEmailSend,
		IsTrackingAPIEnabled: model.IsTrackingAPIEnabled,
		CustomerBOLNumber:    model.CustomerBOLNumber,
		ConsigneeEmail:       model.ShipperEmail,
		ShipperEmail:         model.ConsigneeEmail,
		Result: &v1.Result{
			CapacityProviderBolUrl: model.Result.CapacityProviderBolURL,
			ShipmentIdentifier:     model.Result.ShipmentIdentifier,
			PickupNote:             model.Result.PickupNote,
			PickupDateTime:         model.Result.PickupDateTime,
			Errors:                 model.Result.Errors,
			InfoMessages:           model.Result.InfoMessages,
		},
	}
	return res
}
