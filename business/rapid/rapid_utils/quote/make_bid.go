package quote

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/configs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeBid(quoteRequest *v1.QuoteRequest, bidsData []models.DayDelivery, conf *configs.Config) []*v1.Bid {
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	if len(bidsData) == 0 {
		return nil
	}
	bids := []*v1.Bid{}
	counter := 0
	for _, j := range bidsData {
		for _, k := range j.Standart {
			estimatedDeliveryDate := strings.Split(*k.EstimateDeliveryDate, "T")[0]
			transitTime := fmt.Sprintf("%d", *k.TransitDays)
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
			quoteNumber := ""
			if k.QuoteNumber != nil {
				quoteNumber = *k.QuoteNumber
			}
			index := fmt.Sprintf("%d", counter)
			bid := v1.Bid{
				TransitTime:    transitTime,
				Guranteed:      true,
				VendorQuoteId:  quoteNumber,
				VendorLogo:     logo,
				DeliveryDate:   estimatedDeliveryDate,
				VendorName:     carrierName,
				CarrierName:    carrierName,
				CarrierCode:    carrierCode,
				Carrier:        carrierCodeAddtional,
				BidId:          quoteRequest.QuoteId + "-" + index,
				QuoteId:        quoteRequest.QuoteId,
				CarrierID:      int64(k.CarrierID),
				CarrierQuoteId: capacityProviderQuoteNumber,
				OpportunityId:  int64(k.OpportunityID),
				Amount: &v1.Amount{
					FullAmount: k.Total + float64(conf.Margin),
					NetAmount:  k.Total + 15,
				},
			}
			counter++
			bids = append(bids, &bid)
		}

	}
	return bids
}
