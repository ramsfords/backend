package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

// [
//     {
//         "accessorialId": 28,
//         "name": "Inside Pickup",
//         "code": "INPU",
//         "destinationCode": null,
//         "isOnlyForCanada": false,
//         "isOnlyForUSA": false,
//         "dataType": null,
//         "value": null
//     },
//     {
//         "accessorialId": 29,
//         "name": "Liftgate Pickup",
//         "code": "LGPU",
//         "destinationCode": null,
//         "isOnlyForCanada": false,
//         "isOnlyForUSA": false,
//         "dataType": null,
//         "value": null
//     },
//     {
//         "accessorialId": 30,
//         "name": "Protect From Freeze",
//         "code": "PFZ",
//         "destinationCode": null,
//         "isOnlyForCanada": false,
//         "isOnlyForUSA": true,
//         "dataType": null,
//         "value": null
//     },
//     {
//         "accessorialId": 96,
//         "name": "Single Shipment",
//         "code": "SINGSHIP",
//         "destinationCode": null,
//         "isOnlyForCanada": false,
//         "isOnlyForUSA": false,
//         "dataType": null,
//         "value": null
//     }
// ]

func FixAddressAccesorial(baseQuoteReq *v1.QuoteRequest, rapidQuote *models.QuoteDetails) {
	// pickup fixes
	if len(baseQuoteReq.Pickup.PickupLocationServices) < 1 && len(baseQuoteReq.Delivery.PickupLocationServices) < 1 {
		return
	}
	pickup_services := []models.AddressAccessorial{}
	for _, j := range baseQuoteReq.Pickup.PickupLocationServices {
		if j == v1.PickupLocationServices_PICKUP_LOCATION_WITH_DOCK {
			continue
		}
		if j == v1.PickupLocationServices_INSIDE_PICKUP {
			pickup_services = append(pickup_services, models.AddressAccessorial{
				AccessorialID:   28,
				Name:            "Inside Pickup",
				Code:            "INPU",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
				DataType:        nil,
				Value:           nil,
			})
		}
		if j == v1.PickupLocationServices_LIFTGATE_PICKUP {
			pickup_services = append(pickup_services, models.AddressAccessorial{
				AccessorialID:   29,
				Name:            "Liftgate Pickup",
				Code:            "LGPU",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
				DataType:        nil,
				Value:           nil,
			})
		}
	}
	for _, j := range baseQuoteReq.Commodities {
		for _, k := range j.CommodityServices {
			if k == v1.CommodityServices_PROTECT_FROM_FREEZE {
				pickup_services = append(pickup_services, models.AddressAccessorial{
					AccessorialID:   30,
					Name:            "Protect From Freeze",
					Code:            "PFZ",
					IsOnlyForCanada: false,
					IsOnlyForUSA:    false,
					DataType:        nil,
					Value:           nil,
				})
			}
		}
	}
	rapidQuote.OriginShippingDetails.Address.AddressAccessorials = pickup_services
	// delivery
	delivery_services := []models.AddressAccessorial{}
	for _, j := range baseQuoteReq.Delivery.DeliveryLocationServices {
		if j == v1.DeliveryLocationServices_DELIVERY_LOCATION_WITH_DOCK {
			continue
		}
		if j == v1.DeliveryLocationServices_INSIDE_DELIVERY {
			delivery_services = append(delivery_services, models.AddressAccessorial{
				AccessorialID:   38,
				Name:            "Inside Delivery",
				DestinationCode: "INDEL",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
				DataType:        nil,
				Value:           nil,
			})
		}
		if j == v1.DeliveryLocationServices_DELIVERY_APPOINTMENT {
			delivery_services = append(delivery_services, models.AddressAccessorial{
				AccessorialID:   120,
				Name:            "Delivery Appointment",
				DestinationCode: "APPTDEL",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
				DataType:        0,
			})
		}
		if j == v1.DeliveryLocationServices_LIFTGATE_DELIVERY {
			delivery_services = append(delivery_services, models.AddressAccessorial{
				AccessorialID:   35,
				Name:            "Liftgate Delivery",
				DestinationCode: "LGDEL",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
			})
		}
		if j == v1.DeliveryLocationServices_DELIVERY_NOTIFICATION {
			delivery_services = append(delivery_services, models.AddressAccessorial{
				AccessorialID:   32,
				Name:            "Notify Before Delivery",
				DestinationCode: "NOTIFY",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
			})
		}

	}
	for _, j := range baseQuoteReq.Commodities {
		for _, k := range j.CommodityServices {
			if k == v1.CommodityServices_SORT_AND_SEGREGATE {
				delivery_services = append(delivery_services, models.AddressAccessorial{
					AccessorialID:   36,
					Name:            "Sort/Segregate Delivery",
					DestinationCode: "SORTDEL",
					IsOnlyForCanada: false,
					IsOnlyForUSA:    false,
				})
			}
		}
	}
	rapidQuote.DestinationShippingDetails.Address.AddressAccessorials = delivery_services
	// [
	// {
	//     "accessorialId": 32,
	//     "name": "Notify Before Delivery",
	//     "code": null,
	//     "destinationCode": "NOTIFY",
	//     "isOnlyForCanada": false,
	//     "isOnlyForUSA": false,
	//     "dataType": null,
	//     "value": null
	// },
	// {
	//     "accessorialId": 35,
	//     "name": "Liftgate Delivery",
	//     "code": null,
	//     "destinationCode": "LGDEL",
	//     "isOnlyForCanada": false,
	//     "isOnlyForUSA": false,
	//     "dataType": null,
	//     "value": null
	// },
	// {
	//     "accessorialId": 36,
	//     "name": "Sort/Segregate Delivery",
	//     "code": null,
	//     "destinationCode": "SORTDEL",
	//     "isOnlyForCanada": false,
	//     "isOnlyForUSA": false,
	//     "dataType": null,
	//     "value": null
	// },
	// {
	//     "accessorialId": 38,
	//     "name": "Inside Delivery",
	//     "code": null,
	//     "destinationCode": "INDEL",
	//     "isOnlyForCanada": false,
	//     "isOnlyForUSA": false,
	//     "dataType": null,
	//     "value": null
	// },
	// {
	//     "accessorialId": 120,
	//     "name": "Delivery Appointment",
	//     "code": null,
	//     "destinationCode": "APPTDEL",
	//     "isOnlyForCanada": false,
	//     "isOnlyForUSA": false,
	//     "dataType": 0,
	//     "value": null
	// }
	// ]

}
