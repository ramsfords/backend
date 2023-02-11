package models

type CarrierRequiredField struct {
	CarrierRequiredFieldID int  `json:"carrierRequiredFieldId" dynamodbav:"carrierRequiredFieldId"`
	ShipperEmailAddress    bool `json:"shipperEmailAddress" dynamodbav:"shipperEmailAddress"`
	ConsigneeEmailAddress  bool `json:"consigneeEmailAddress" dynamodbav:"consigneeEmailAddress"`
	ShipperLastName        bool `json:"shipperLastName" dynamodbav:"shipperLastName"`
	ConsigneeLastName      bool `json:"consigneeLastName" dynamodbav:"consigneeLastName"`
	NmfcCode               bool `json:"nmfcCode" dynamodbav:"nmfcCode"`
	ConsigneeTime          bool `json:"consigneeTime" dynamodbav:"consigneeTime"`
}
