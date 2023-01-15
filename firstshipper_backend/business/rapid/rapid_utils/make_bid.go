package rapid_utils

import (
	"fmt"
	"strings"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeBid(saveQuote *models.SaveQuote, qtReq *v1.QuoteRequest) []*v1.Bid {
	if !saveQuote.QuoteRate.IsValid {
		return nil
	}
	bids := make([]*v1.Bid, 0)
	for _, j := range saveQuote.QuoteRate.DayDeliveries {
		for i, k := range j.Standart {
			estimatedDeliveryDate := strings.Split(*k.EstimateDeliveryDate, "T")[0]
			transitTime := fmt.Sprintf("%d", *k.TransitDays)
			largeLogo := ""
			if k.LargeLogo != nil {
				largeLogo = *k.LargeLogo
			}
			logo := ""
			if k.Logo != nil {
				logo = *k.Logo
			}
			carrierName := ""
			if k.CarrierName != nil {
				carrierName = *k.CarrierName
			}
			carrierCode := ""
			if k.CarrierCode != nil {
				carrierCode = *k.CarrierCode
			}
			carrierCodeAddtional := ""
			if k.CarrierCodeAdditional != nil {
				carrierCodeAddtional = *k.CarrierCodeAdditional
			}
			capacityProviderQuoteNumber := ""
			if k.CapacityProviderQuoteNumber != nil {
				capacityProviderQuoteNumber = *k.CapacityProviderQuoteNumber
			}
			bid := v1.Bid{
				TransitTime:         transitTime,
				Guranteed:           true,
				VendorQuoteId:       saveQuote.SavedQuoteID,
				QuoteId:             qtReq.QuoteId,
				CompanyLargeLogoUrl: largeLogo,
				CompanySmallLogoUrl: logo,
				DeliveryDate:        estimatedDeliveryDate,
				VendorName:          "rapid",
				CarrierName:         carrierName,
				CarrierCode:         carrierCode,
				Carrier:             carrierCodeAddtional,
				BidId:               qtReq.QuoteId + "#" + fmt.Sprint(i),
				BusinessId:          qtReq.BusinessId,
				CarrierQuoteId:      capacityProviderQuoteNumber,
				Amount: &v1.Amount{
					FullAmount: k.Total + 15,
					NetAmount:  k.Total + 15,
				},
			}
			bids = append(bids, &bid)
		}
	}
	return bids
}
