package models

type Result struct {
	CapacityProviderBolURL string   `json:"capacityProviderBolUrl" dynamodbav:"capacityProviderBolUrl"`
	ShipmentIdentifier     string   `json:"shipmentIdentifier" dynamodbav:"shipmentIdentifier"`
	PickupNote             string   `json:"pickupNote" dynamodbav:"pickupNote"`
	PickupDateTime         string   `json:"pickupDateTime" dynamodbav:"pickupDateTime"`
	Errors                 []string `json:"errors" dynamodbav:"errors"`
	InfoMessages           []string `json:"infoMessages" dynamodbav:"infoMessages"`
}
type DispatchResponse struct {
	ShipmentID           int     `json:"shipmentId" dynamodbav:"shipmentId"`
	SecurityKey          string  `json:"securityKey" dynamodbav:"securityKey"`
	PickupNumber         string  `json:"pickupNumber" dynamodbav:"pickupNumber"`
	CarrierName          string  `json:"carrierName" dynamodbav:"carrierName"`
	CarrierPhone         string  `json:"carrierPhone" dynamodbav:"carrierPhone"`
	CarrierPRONumber     string  `json:"carrierPRONumber" dynamodbav:"carrierPRONumber"`
	HandlingUnitTotal    float64 `json:"handlingUnitTotal" dynamodbav:"handlingUnitTotal"`
	IsShipmentEdit       bool    `json:"isShipmentEdit" dynamodbav:"isShipmentEdit"`
	IsShipmentManual     bool    `json:"isShipmentManual" dynamodbav:"isShipmentManual"`
	ServiceType          int     `json:"serviceType" dynamodbav:"serviceType"`
	IsTrackingEmailSend  bool    `json:"isTrackingEmailSend" dynamodbav:"isTrackingEmailSend"`
	IsTrackingAPIEnabled bool    `json:"isTrackingAPIEnabled" dynamodbav:"isTrackingAPIEnabled"`
	CustomerBOLNumber    string  `json:"customerBOLNumber" dynamodbav:"customerBOLNumber"`
	ShipperEmail         string  `json:"shipperEmail" dynamodbav:"shipperEmail"`
	ConsigneeEmail       string  `json:"consigneeEmail" dynamodbav:"consigneeEmail"`
	Result               Result  `json:"result" dynamodbav:"result"`
}
