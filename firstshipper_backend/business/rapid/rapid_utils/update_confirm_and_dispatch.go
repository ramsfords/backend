package rapid_utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
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
func NewConfirmAndDispatchStep3(quoteReq *v1.QuoteResponse, saveQuote *models.SaveQuote) (*models.SaveQuote, error) {
	//TODO: fix HandlingUnitTotal in a quote
	//TODO: fix HandlingUnitTotalPackages
	// confirmAndDispatch, err := GetConfirmAndDispatch(booking, saveQuote)
	// if err != nil {
	// 	return nil, errors.New("could not make confirm and dispatch step 3")
	// }
	// if confirmAndDispatch != nil {
	// 	saveQuote.ConfirmAndDispatch = confirmAndDispatch
	// } else {
	// 	saveQuote.ConfirmAndDispatch = &models.ConfirmAndDispatch{}
	// 	saveQuote.ConfirmAndDispatch = confirmAndDispatch
	// }
	confirmAndDispatch := &models.ConfirmAndDispatch{}
	rateDetails := GetRateDetails(quoteReq, saveQuote)
	bytes, err := json.Marshal(rateDetails)
	if err != nil {
		return saveQuote, err
	}
	err = json.Unmarshal(bytes, confirmAndDispatch)
	if err != nil {
		return saveQuote, err
	}
	confirmAndDispatch.ShipmentPriceDetails = confirmAndDispatch.RateQuoteDetails
	confirmAndDispatch.RateQuoteDetails = nil
	saveQuote.ConfirmAndDispatch = confirmAndDispatch
	saveQuote.ConfirmAndDispatch.BillingAddress = GetBillingAddress()
	var handlingUnitVolume float32
	var totalShipmentWeight float32
	var handelingUnits int
	var handlingUnitTotalPackages int
	for _, j := range quoteReq.QuoteRequest.Commodities {
		//75*48*40 * 0.0005787
		handlingUnitVolume += j.Length * j.Width * j.Height * 0.0005787
		totalShipmentWeight += j.Weight
		handelingUnits += 1
		handlingUnitTotalPackages += 1

	}
	handlingUnitDensity := totalShipmentWeight / handlingUnitVolume
	saveQuote.ConfirmAndDispatch.HandlingUnitDensity = float64(handlingUnitDensity)
	saveQuote.ConfirmAndDispatch.HandlingUnitTotal = handelingUnits
	saveQuote.ConfirmAndDispatch.HandlingUnitTotalPackages = handlingUnitTotalPackages
	saveQuote.ConfirmAndDispatch.HandlingUnitVolume = float64(handlingUnitVolume)
	saveQuote.ConfirmAndDispatch.IconLogo = rateDetails.IconLogo
	saveQuote.ConfirmAndDispatch.OpportunityID = rateDetails.OpportunityID
	saveQuote.ConfirmAndDispatch.ServiceType = rateDetails.ServiceType
	saveQuote.ConfirmAndDispatch.TotalCost = rateDetails.TotalCost
	saveQuote.ConfirmAndDispatch.FreightCharge = rateDetails.FreightCharge
	saveQuote.ConfirmAndDispatch.EstimateDeliveryDate = rateDetails.EstimateDeliveryDate
	saveQuote.ConfirmAndDispatch.TotalShipmentWeight = int(totalShipmentWeight)
	saveQuote.ConfirmAndDispatch.ShipperDetails = GetShippperPartyDetails(quoteReq, saveQuote)
	saveQuote.ConfirmAndDispatch.ConsigneeDetails = GetConsigneePartyDetails(quoteReq, saveQuote)
	saveQuote.ConfirmAndDispatch.EmergencyContactPerson = GetEmmergencyContact(quoteReq)
	bol := "BOL" + quoteReq.QuoteRequest.QuoteId
	reference := "bid" + quoteReq.QuoteRequest.QuoteId
	saveQuote.ConfirmAndDispatch.ReferenceNumberInfo = models.ReferenceNumberInfo{
		CustomerBOL:     &bol,
		ReferenceNumber: &reference,
	}
	nim := strings.Split(strings.Split(time.Now().UTC().String(), " ")[0], "-")
	bookDate := nim[2] + "/" + nim[1] + "/" + nim[0]
	// t1, _ := time.ParseInLocation(RFC3339local, nim, loc)
	// fmt.Println(t1)
	// now, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC1123Z))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	saveQuote.ConfirmAndDispatch.BookedDate = &bookDate
	saveQuote.ConfirmAndDispatch.ConsigneeDetails.InstructionNote = &quoteReq.QuoteRequest.ReceiverInstructions

	return saveQuote, nil
}

func GetRateDetails(quoteRes *v1.QuoteResponse, saveQuote *models.SaveQuote) models.Standard {
	var standard models.Standard
	for _, j := range saveQuote.QuoteRate.DayDeliveries {
		for _, k := range j.Standart {
			for _, j := range quoteRes.Bids {
				if *k.CarrierName == j.CarrierName {
					standard = k
					break
				}
			}
		}
	}
	return standard
}

// func StrPtr(value string) *string {
// 	return &value
// }

// func GetConfirmAndDispatch(booking *v1.Booking, saveQuote *models.SaveQuote) (*models.ConfirmAndDispatch, error) {
// 	confirmAndDispatch := &saveQuote.ConfirmAndDispatch
// 	standard := GetRateDetails(booking, saveQuote)
// 	bytes, err := json.Marshal(standard)
// 	if err != nil {
// 		return nil, errors.New("could not marshall saveQuote to Standard")
// 	}
// 	err = json.Unmarshal(bytes, confirmAndDispatch)
// 	if err != nil {
// 		return nil, errors.New("could not unmarshall standard to confirm and dispatch")
// 	}
// 	return *confirmAndDispatch, nil
// }
