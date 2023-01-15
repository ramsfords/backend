package rapid_utils

// LocationService_inside_pickup              LocationService = 2
// LocationService_liftgate_pickup            LocationService = 3
// LocationService_pickup_notification        LocationService = 4
// func CommodityAccesorials(quote v1.Quote) []models.AddressAccessorial {
// 	addressAsso := []models.AddressAccessorial{}
// 	for _, j := range quote.Origin.LocationServices {
// 		var currentService models.AddressAccessorial
// 		switch j {
// 		case 2:
// 			currentService = InsidePickup()
// 		case 3:
// 			currentService = OriginLiftGate()
// 		}
// 		addressAsso = append(addressAsso, currentService)

// 	}
// 	protectFromFreeze := false
// 	for _, k := range quote.Commodities {
// 		for _, j := range k.CommodityServices {
// 			if j == 1 {
// 				protectFromFreeze = true
// 			}
// 		}
// 	}
// 	if protectFromFreeze {
// 		addressAsso = append(addressAsso, ProtectFromFreeze())
// 	}
// 	return addressAsso
// }
