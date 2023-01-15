package utils

import v1 "github.com/ramsfords/types_gen/v1"

func toInt32ArrayFromCommodityServices(services []v1.CommodityServices) []int32 {
	arr := []int32{}
	for _, j := range services {
		arr = append(arr, int32(*j.Enum()))
	}
	return arr
}
func toInt32ArrayFromPickupLocationServices(services []v1.PickupLocationServices) []int32 {
	arr := []int32{}
	for _, j := range services {
		arr = append(arr, int32(*j.Enum()))
	}
	return arr
}
func toInt32ArrayFromDeliveryLocationServices(services []v1.DeliveryLocationServices) []int32 {
	arr := []int32{}
	for _, j := range services {
		arr = append(arr, int32(*j.Enum()))
	}
	return arr
}
