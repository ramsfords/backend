package utils

import v1 "github.com/ramsfords/types_gen/v1"

func FixAdditionalServices(req v1.CommodityServices) []v1.CommodityServices {
	services := []v1.CommodityServices{}
	if req == v1.CommodityServices_STACKABLE {
		services = append(services, v1.CommodityServices_STACKABLE)
	}
	if req == v1.CommodityServices_PROTECT_FROM_FREEZE {
		services = append(services, v1.CommodityServices_PROTECT_FROM_FREEZE)
	}
	if req == v1.CommodityServices_SORT_AND_SEGREGATE {
		services = append(services, v1.CommodityServices_SORT_AND_SEGREGATE)
	}
	if req == v1.CommodityServices_GUARANTEED {
		services = append(services, v1.CommodityServices_GUARANTEED)
	}
	if req == v1.CommodityServices_HAZARDOUS {
		services = append(services, v1.CommodityServices_HAZARDOUS)
	}
	return services
}
