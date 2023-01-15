package rapid_utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func MakeQuoteDetails(baseQuote *v1.QuoteRequest) (*models.QuoteDetails, error) {
	totalWeight := fmt.Sprintf("%f", baseQuote.Commodities[0].Weight)
	totalWeight = strings.Split(totalWeight, ".")[0] + ".00"
	qtReq := &models.QuoteDetails{
		BillingAddress: models.Address{
			AddressID:           1039341,
			AddressAccessorials: []models.AddressAccessorial{},
		},
		FreightCharge:                  1,
		ServiceType:                    1,
		IsShowVLTLToggle:               true,
		IsAutoEmailTrackingEnabled:     true,
		IsManualDispatchSettingEnabled: true,
		IsExistConnectedCarriers:       true,
		CustomerCarrierID:              1070,
		TruckloadCapacityProviders:     []string{"EMRG"},
		TotalWeight:                    totalWeight,
		QuoteDate:                      time.Now().Format("01/02/2006"),
	}
	pickupDate, err := time.Parse(time.RFC3339, baseQuote.ShipmentDetails.PickupDate)
	if err != nil {
		return nil, errors.New("can not convert pick up date to rapid pickup date")
	}
	origin, err := FixOriginShippingDetails(baseQuote.Pickup)
	if err != nil {
		fmt.Println("could not parse origin details for base quote origin to rapid quote origin")
	}
	pickupDateFormat := pickupDate.Format("01/02/2006")
	origin.PickupDate = pickupDateFormat
	qtReq.OriginShippingDetails = *origin
	destination, err := FixDestinationShippingDetails(baseQuote.Delivery)
	if err != nil {
		return nil, fmt.Errorf("could not parse origin details for base quote origin to rapid quote origin")
	}
	qtReq.DestinationShippingDetails = *destination
	qtReq.ShipmentItems = AddShipments(baseQuote.Commodities)
	FixAddressAccesorial(baseQuote, qtReq)
	return qtReq, nil
}
