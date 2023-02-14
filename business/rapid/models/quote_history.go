package models

type QuoteHistory struct {
	SavedQuoteID         string        `json:"savedQuoteId,omitempty"`
	Step                 int           `json:"step,omitempty"`
	ReferenceID          string        `json:"referenceId,omitempty"`
	PickupDate           string        `json:"pickupDate,omitempty"`
	ShipperCompanyName   interface{}   `json:"shipperCompanyName,omitempty"`
	OriginAddress        string        `json:"originAddress,omitempty"`
	ConsigneeCompanyName interface{}   `json:"consigneeCompanyName,omitempty"`
	DestinationAddress   string        `json:"destinationAddress,omitempty"`
	CarrierName          interface{}   `json:"carrierName,omitempty"`
	TotalCost            string        `json:"totalCost,omitempty"`
	ServiceType          int           `json:"serviceType,omitempty"`
	ServiceTypeName      string        `json:"serviceTypeName,omitempty"`
	QuoteID              int           `json:"quoteId,omitempty"`
	TotalWeight          int           `json:"totalWeight,omitempty"`
	TimeCreated          interface{}   `json:"timeCreated,omitempty"`
	ShipmentID           interface{}   `json:"shipmentId,omitempty"`
	IsFavorite           bool          `json:"isFavorite,omitempty"`
	IsArchivedQuote      bool          `json:"isArchivedQuote,omitempty"`
	IsShipmentCanceled   bool          `json:"isShipmentCanceled,omitempty"`
	OrderID              interface{}   `json:"orderId,omitempty"`
	ErrorSteps           []interface{} `json:"errorSteps,omitempty"`
	OriginCode           interface{}   `json:"originCode,omitempty"`
}
