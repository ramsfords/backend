package quote

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeQuoteDetails(quoteRequest *v1.QuoteRequest) (*models.QuoteDetails, error) {
	totalWeight := fmt.Sprintf("%f", quoteRequest.Commodities[0].Weight)
	totalWeight = strings.Split(totalWeight, ".")[0] + ".00"
	pickupDate, err := time.Parse(time.RFC3339, quoteRequest.PickupDate)
	if err != nil {
		return nil, errors.New("can not convert pick up date to rapid pickup date")
	}
	pickupDateFormated := pickupDate.Format("01/02/2006")
	rapidQuoteRequest := &models.QuoteDetails{
		BillingAddress:                 getBillingAddress(),
		FreightCharge:                  1,
		ServiceType:                    1,
		IsShowVLTLToggle:               true,
		IsAutoEmailTrackingEnabled:     true,
		IsManualDispatchSettingEnabled: true,
		IsExistConnectedCarriers:       true,
		CustomerCarrierID:              1070,
		TotalWeight:                    totalWeight,
		QuoteDate:                      time.Now().Format("01/02/2006"),
		OriginShippingDetails: models.OriginShippingDetails{
			Address: models.Address{
				AddressID:  1039341,
				PostalCode: quoteRequest.Pickup.Address.ZipCode,
				CommercialType: &models.CommercialType{
					AccessorialID: 72,
					Name:          "Business",
				},
				AddressAccessorials: []models.AddressAccessorial{},
			},
			PickupDate: pickupDateFormated,
		},
		DestinationShippingDetails: models.DestinationShippingDetails{
			Address: models.Address{
				PostalCode: quoteRequest.Delivery.Address.ZipCode,
				CommercialType: &models.CommercialType{
					AccessorialID: 72,
					Name:          "Business",
				},
				AddressAccessorials: []models.AddressAccessorial{},
			},
		},
		CargoInsuranceQuoteInfo: getNewCargoInsuranceQuoteInfo(),
	}
	AddShipments(quoteRequest, rapidQuoteRequest)
	FixAddressAccesorial(quoteRequest, rapidQuoteRequest)
	return rapidQuoteRequest, nil
}
