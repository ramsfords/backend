package quote

import (
	"github.com/ramsfords/backend/business/rapid/models"
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
	//pickup fixes
	pickup_services := []models.AddressAccessorial{}
	if baseQuoteReq.LocationServices.InsidePickup {
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
	if baseQuoteReq.LocationServices.LiftGatePickup {
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
	if baseQuoteReq.LocationServices.LiftGatePickup {
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

	rapidQuote.OriginShippingDetails.Address.AddressAccessorials = pickup_services
	// delivery
	delivery_services := []models.AddressAccessorial{}
	if baseQuoteReq.LocationServices.InsideDelivery {
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
	if baseQuoteReq.LocationServices.DeliveryAppointment {
		delivery_services = append(delivery_services, models.AddressAccessorial{
			AccessorialID:   120,
			Name:            "Delivery Appointment",
			DestinationCode: "APPTDEL",
			IsOnlyForCanada: false,
			IsOnlyForUSA:    false,
			DataType:        0,
		})
	}
	if baseQuoteReq.LocationServices.LiftGateDelivery {
		delivery_services = append(delivery_services, models.AddressAccessorial{
			AccessorialID:   35,
			Name:            "Liftgate Delivery",
			DestinationCode: "LGDEL",
			IsOnlyForCanada: false,
			IsOnlyForUSA:    false,
		})
	}
	if baseQuoteReq.LocationServices.DeliveryAppointment {
		delivery_services = append(delivery_services, models.AddressAccessorial{
			AccessorialID:   32,
			Name:            "Notify Before Delivery",
			DestinationCode: "NOTIFY",
			IsOnlyForCanada: false,
			IsOnlyForUSA:    false,
		})
	}

	for _, j := range baseQuoteReq.Commodities {
		if j.CommodityServices.SortAndSegregate {
			delivery_services = append(delivery_services, models.AddressAccessorial{
				AccessorialID:   36,
				Name:            "Sort/Segregate Delivery",
				DestinationCode: "SORTDEL",
				IsOnlyForCanada: false,
				IsOnlyForUSA:    false,
			})
		}
	}
	rapidQuote.DestinationShippingDetails.Address.AddressAccessorials = delivery_services
	for _, j := range baseQuoteReq.Commodities {
		if j.CommodityServices.ProtectFromFreeze {
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
