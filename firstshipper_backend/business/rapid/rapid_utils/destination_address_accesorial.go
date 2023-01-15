package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// LocationService_inside_delivery            LocationService = 5
// LocationService_liftgate_delivery          LocationService = 6
// LocationService_delivery_notification      LocationService = 7
// CommodityServices_stackable           CommodityServices = 0
// CommodityServices_protect_from_freeze CommodityServices = 1
// CommodityServices_sort_and_segregate  CommodityServices = 2
// CommodityServices_guaranteed          CommodityServices = 3
// CommodityServices_hazardous           CommodityServices = 4
func FixDestinationAddressAccesorial(deliveryLocation *v1.Location, newQuote *models.DestinationShippingDetails) {
	for _, deliveryLocationServices := range deliveryLocation.DeliveryLocationServices {
		if deliveryLocationServices == v1.DeliveryLocationServices_LIFTGATE_DELIVERY {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, LiftGateDelivery())
		}
		if deliveryLocationServices == v1.DeliveryLocationServices_INSIDE_DELIVERY {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, InsideDelivery())
		}
		if deliveryLocationServices == v1.DeliveryLocationServices_DELIVERY_APPOINTMENT {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, DeliveryAppt())
		}
		if deliveryLocationServices == v1.DeliveryLocationServices_DELIVERY_NOTIFICATION {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, NotifyBeforeDelivery())
		}
	}

	// for _, j := range baseOrigin.LocationServices {
	// 	var currentService models.AddressAccessorial
	// 	switch j {
	// 	case 5:
	// 		currentService = InsideDelivery()
	// 	case 6:
	// 		currentService = LiftGateDelivery()
	// 	case 7:
	// 		currentService = NotifyBeforeDelivery()
	// 	case 8:
	// 		currentService = DeliveryAppt()
	// 	}
	// 	if currentService.AccessorialID != 0 {
	// 		newQuote.AddressAccessorials = append(newQuote.AddressAccessorials, currentService)
	// 	}

	// }
}
