package models

type CarrierAPIs struct {
	CarrierAPIID             int  `json:"carrierAPIId" dynamodbav:"carrierAPIId"`
	QuotingAPIEnabled        bool `json:"quotingAPIEnabled" dynamodbav:"quotingAPIEnabled"`
	DispatchAPIEnabled       bool `json:"dispatchAPIEnabled" dynamodbav:"dispatchAPIEnabled"`
	TrackingAPIEnabled       bool `json:"trackingAPIEnabled" dynamodbav:"trackingAPIEnabled"`
	DocumentImagesAPIEnabled bool `json:"documentImagesAPIEnabled" dynamodbav:"documentImagesAPIEnabled"`
}
