package book

import (
	"time"

	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// missing or mismatched
// "billingAddressId": 0 =>  "billingAddressId": 1039341
// "estimateDeliveryDate": "2022-06-23T00:00:00", => "estimateDeliveryDate": "06/23/2022"
// "freightCharge": 1,
// "handlingUnitDensity": 9.6,
// "handlingUnitTotal": 1,
// "handlingUnitTotalPackages": 1,
// "handlingUnitVolume": 83.33333333333334,
// "freightCharge": 0,
// "handlingUnitDensity": 0,
// "handlingUnitTotal": 0,
// "handlingUnitTotalPackages": 0,
// "handlingUnitVolume": 0,
// "iconLogo": "https://content.mycarriertms.com/carriers/651af2b5-638a-463e-918d-9dedb42da42b.png",
// "iconLogo": null,
// "opportunityId": 7444451,
// "originLocationId": 1039341,
// "quoteId": 7444451,
// "serviceType": 1,
// "billingAddressId": 0,
// "totalCost": 293.79,
// "totalShipmentWeight": 800,
// "opportunityId": 0,
// "originLocationId": 0,
// "quoteId": 0,
// "totalCost": 0,
// "totalShipmentWeight": 0,
func newConfirmAndDispatch(quoteRequest *model.QuoteRequest, bid *v1.Bid) error {
	billingAddress := getBillingAddress()
	var handlingUnitVolume float32
	var totalShipmentWeight float32
	var handelingUnits int
	var handlingUnitTotalPackages int
	for _, j := range quoteRequest.QuoteRequest.Commodities {
		//75*48*40 * 0.0005787
		handlingUnitVolume += j.Length * j.Width * j.Height * 0.0005787
		totalShipmentWeight += j.Weight
		handelingUnits += 1
		handlingUnitTotalPackages += 1

	}
	handlingUnitDensity := totalShipmentWeight / handlingUnitVolume
	rateDetails := getRateDetails(quoteRequest, bid)

	bol := "BOL" + quoteRequest.QuoteRequest.QuoteId
	poNumber := "PO" + quoteRequest.QuoteRequest.QuoteId
	bookedDate := time.Now().Format("01/02/2006")
	confirmAndDispatch := &models.ConfirmAndDispatch{
		CapacityProviderAccountGroup: GetCapacityProviderAccountGroup(bid),
		CapacityProviderQuoteNumber:  rateDetails.CapacityProviderQuoteNumber,
		CarrierCode:                  rateDetails.CarrierCode,
		CarrierCodeAdditional:        rateDetails.CarrierCode,
		ShipperDetails:               getShippperPartyDetails(quoteRequest),
		ConsigneeDetails:             getConsigneePartyDetails(quoteRequest),
		BillingAddress:               billingAddress,
		ReferenceNumberInfo: models.ReferenceNumberInfo{
			CustomerBOL:     &bol,
			ReferenceNumber: &bid.BidId,
			PoNumber:        &poNumber,
			PickupNumber:    &poNumber,
		},
		EmergencyContactPerson:    getEmmergencyContact(quoteRequest),
		HandlingUnits:             quoteRequest.RapidSaveQuote.QuoteDetails.ShipmentItems,
		HandlingUnitTotal:         handelingUnits,
		HandlingUnitTotalPackages: handlingUnitTotalPackages,
		HandlingUnitVolume:        float64(handlingUnitVolume),
		HandlingUnitDensity:       float64(handlingUnitDensity),
		TotalShipmentWeight:       int(totalShipmentWeight),
		BillingAddressID:          billingAddress.AddressID,
		CarrierID:                 rateDetails.CarrierID,
		CarrierName:               rateDetails.CarrierName,
		IconLogo:                  rateDetails.IconLogo,
		OpportunityID:             rateDetails.OpportunityID,
		QuoteID:                   rateDetails.QuoteID,
		ServiceLevelCode:          rateDetails.ServiceLevelCode,
		ServiceType:               rateDetails.OpportunityID,
		ShipmentNote:              &quoteRequest.QuoteRequest.SpecialInstruction,
		TransitTime:               *rateDetails.TransitDays,
		TotalCost:                 rateDetails.TotalCost,
		EstimateDeliveryDate:      rateDetails.EstimateDeliveryDate,
		FreightCharge:             rateDetails.FreightCharge,
		BookedDate:                &bookedDate,
		SpecialInstruction:        &quoteRequest.QuoteRequest.SpecialInstruction,
		ShipmentPriceDetails:      rateDetails.ShipmentPriceDetails,
	}
	quoteRequest.RapidSaveQuote.ConfirmAndDispatch = confirmAndDispatch
	return nil
}

func getRateDetails(quoteReq *model.QuoteRequest, bid *v1.Bid) models.Standard {
	var standard models.Standard
	for _, j := range quoteReq.RapidSaveQuote.QuoteRate.DayDeliveries {
		for _, k := range j.Standart {
			if *k.CarrierName == bid.CarrierName {
				standard = k
				break
			}
		}

	}
	return standard
}
