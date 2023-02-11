package roadrunner

import (
	v1 "github.com/ramsfords/types_gen/v1"
)

const (
	// RoadRunnerUsername is the username for the RoadRunner API
	LIFTGATEPICKUP                  string = "LGP"
	LIFTGATEDELIVERY                string = "LGD"
	INSIDEDELIVERY                  string = "ID"
	INSIDEPICKUP                    string = "IP"
	LIMITEDACCESSDELIVERY           string = "LTD"
	LIMITEDACCESSPICUP              string = "LTP"
	NOTIFICATIONCHARGE              string = "NC"
	MAINTAINTEMPRETURE              string = "PSM"
	NONBUSINESSHOURSDELIVERY        string = "NBD"
	PROTECTFROMCOLD                 string = "PSC"
	PROTECTFROMHEAT                 string = "PSH"
	RESIDENTIALDELIVERY             string = "RSD"
	RESIDENTIALPICKUP               string = "RSP"
	SORTANDSEGREGATE                string = "SRT"
	CONTAINERFREIGHTSTATIONPICKUP   string = "CFS"
	CONTAINERFREIGHTSTATIONDELIVERY string = "CFD"
	APPOINTMENT                     string = "APT"
)

func GetDeliveryServices(locationServices *v1.LocationServices) []ServiceOption {
	servicesDeliveryOptions := []ServiceOption{}
	if locationServices.LiftGatePickup {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: LIFTGATEPICKUP})
	}
	if locationServices.LiftGateDelivery {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: LIFTGATEDELIVERY})
	}
	if locationServices.InsideDelivery {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: INSIDEDELIVERY})
	}
	if locationServices.InsidePickup {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: INSIDEPICKUP})
	}
	if locationServices.PickupNotification {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: NOTIFICATIONCHARGE})
	}
	if locationServices.DeliveryNotification {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: NOTIFICATIONCHARGE})
	}
	if locationServices.ReceiverPickupNotification {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: NOTIFICATIONCHARGE})
	}
	if locationServices.ShipperDeliveryNotification {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: NOTIFICATIONCHARGE})
	}
	if locationServices.PickupAppointment {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: APPOINTMENT})
	}
	if locationServices.DeliveryAppointment {
		servicesDeliveryOptions = append(servicesDeliveryOptions, ServiceOption{ServiceCode: APPOINTMENT})
	}
	return servicesDeliveryOptions
}
