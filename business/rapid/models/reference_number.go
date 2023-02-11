package models

type ReferenceNumberInfo struct {
	CustomerBOL      *string `json:"customerBOL,omitempty" dynamodbav:"customerBOL"`
	PickupNumber     *string `json:"pickupNumber,omitempty" dynamodbav:"pickupNumber"`
	PoNumber         *string `json:"poNumber,omitempty" dynamodbav:"poNumber"`
	ReferenceNumber  *string `json:"referenceNumber,omitempty" dynamodbav:"referenceNumber"`
	GS1CompanyPrefix *string `json:"gS1CompanyPrefix,omitempty" dynamodbav:"gS1CompanyPrefix"`
	ProNumber        *string `json:"proNumber" dynamodbav:"gS1CompanyPrefix"`
}

func GetNewReferenceNumberInfo(booking Wrapper) ReferenceNumberInfo {
	return ReferenceNumberInfo{
		CustomerBOL:     &booking.Bid.CapacityProviderQuoteId,
		PoNumber:        &booking.Bid.CapacityProviderQuoteId,
		PickupNumber:    &booking.Bid.CapacityProviderQuoteId,
		ReferenceNumber: &booking.Bid.CapacityProviderQuoteId,
	}
}
