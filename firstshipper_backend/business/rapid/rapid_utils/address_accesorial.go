package rapid_utils

import "github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"

func InsidePickup() models.AddressAccessorial {
	/*
		{
		"accessorialId": 28,
		"name": "Inside Pickup",
		"code": "INPU",
		"destinationCode": null,
		"isOnlyForCanada": false,
		"isOnlyForUSA": false,
		"dataType": null,
		"value": null
		}
	*/
	return models.AddressAccessorial{
		AccessorialID: 28,
		Code:          "INPU",
		Name:          "Inside Pickup",
	}

}
func OriginLiftGate() models.AddressAccessorial {
	/*
			{
		                    "accessorialId": 29,
		                    "name": "Liftgate Pickup",
		                    "code": "LGPU",
		                    "destinationCode": null,
		                    "isOnlyForCanada": false,
		                    "isOnlyForUSA": false,
		                    "dataType": null,
		                    "value": null
		                }
	*/
	return models.AddressAccessorial{
		AccessorialID: 29,
		Name:          "Liftgate Pickup",
		Code:          "LGPU",
	}

}

func ProtectFromFreeze() models.AddressAccessorial {
	/*
		{
			"accessorialId": 30,
			"name": "Protect From Freeze",
			"code": "PFZ",
			"destinationCode": null,
			"isOnlyForCanada": false,
			"isOnlyForUSA": true,
			"dataType": null,
			"value": null
		}
	*/
	return models.AddressAccessorial{
		AccessorialID: 30,
		Name:          "Protect From Freeze",
		Code:          "PFZ",
		IsOnlyForUSA:  true,
	}
}

func AddOriginInBondFreight() models.AddressAccessorial {
	/*
		{
			"accessorialId": 121,
			"name": "In Bond Freight",
			"code": "BOND",
			"destinationCode": null,
			"isOnlyForCanada": false,
			"isOnlyForUSA": false,
			"dataType": null,
			"value": null
		}
	*/
	return models.AddressAccessorial{
		AccessorialID: 121,
		Name:          "In Bond Freight",
		Code:          "BOND",
	}
}

func InsideDelivery() models.AddressAccessorial {
	/*
		{
			accessorialId: 38
			code: null
			dataType: null
			destinationCode: "INDEL"
			isOnlyForCanada: false
			isOnlyForUSA: false
			name: "Inside Delivery"
			value: null
		},
	*/
	return models.AddressAccessorial{
		AccessorialID:   38,
		Name:            "Inside Delivery",
		IsOnlyForUSA:    false,
		DestinationCode: "INDEL",
	}
}
func LiftGateDelivery() models.AddressAccessorial {
	/*
		{
			accessorialId: 35
			code: null
			dataType: null
			destinationCode: "LGDEL"
			isOnlyForCanada: false
			isOnlyForUSA: false
			name: "Liftgate Delivery"
			value: null
		}
	*/
	return models.AddressAccessorial{
		AccessorialID:   35,
		Name:            "Liftgate Delivery",
		IsOnlyForUSA:    false,
		DestinationCode: "LGDEL",
	}

}
func NotifyBeforeDelivery() models.AddressAccessorial {
	/*
		accessorialId: 32
		code: null
		dataType: null
		destinationCode: "NOTIFY"
		isOnlyForCanada: false
		isOnlyForUSA: false
		name: "Notify Before Delivery"
		value: null
	*/
	return models.AddressAccessorial{
		AccessorialID:   32,
		Name:            "Notify Before Delivery",
		DestinationCode: "NOTIFY",
	}

}
func SortAndSegregateDelivery() models.AddressAccessorial {
	/*
		accessorialId: 36
		code: null
		dataType: null
		destinationCode: "SORTDEL"
		isOnlyForCanada: false
		isOnlyForUSA: false
		name: "Sort/Segregate Delivery"
		value: null
	*/
	return models.AddressAccessorial{
		AccessorialID:   36,
		Name:            "Sort/Segregate Delivery",
		IsOnlyForUSA:    false,
		DestinationCode: "SORTDEL",
	}
}
func DeliveryAppt() models.AddressAccessorial {
	/*
		{
			accessorialId: 120
			code: null
			dataType: 0
			destinationCode: "APPTDEL"
			isOnlyForCanada: false
			isOnlyForUSA: false
			name: "Delivery Appointment"
			value: null
		}
	*/
	return models.AddressAccessorial{
		AccessorialID:   120,
		Name:            "Delivery Appointment",
		DestinationCode: "APPTDEL",
	}
}
