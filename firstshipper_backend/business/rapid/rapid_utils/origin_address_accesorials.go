package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func FixOriginAddressAccessorials(baseOrigin *v1.Location, newQuote *models.OriginShippingDetails) {
	// for _, j := range baseOrigin.LocationServices {
	// 	var currentService models.AddressAccessorial
	// 	switch j {
	// 	case 2:
	// 		currentService = InsidePickup()
	// 	case 3:
	// 		currentService = OriginLiftGate()
	// 		// case 4:
	// 		// 	currentService = PickupNotification()
	// 	}

	// 	if currentService.AccessorialID != 0 {
	// 		newQuote.AddressAccessorials = append(newQuote.AddressAccessorials, currentService)
	// 	}

	// }
	for _, pickupLocationServices := range baseOrigin.PickupLocationServices {
		if pickupLocationServices == v1.PickupLocationServices_LIFTGATE_PICKUP {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, OriginLiftGate())
		}
		if pickupLocationServices == v1.PickupLocationServices_INSIDE_PICKUP {
			newQuote.Address.AddressAccessorials = append(newQuote.Address.AddressAccessorials, InsidePickup())
		}
		if pickupLocationServices == v1.PickupLocationServices_PICKUP_APPOINTMENT {
			// newQuote.AddressAccessorials = append(newQuote.AddressAccessorials, DeliveryAppt())
		}

	}

}
