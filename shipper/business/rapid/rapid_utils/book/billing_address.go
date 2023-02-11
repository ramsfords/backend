package book

import "github.com/ramsfords/backend/shipper/business/rapid/models"

func getBillingAddress() models.Address {
	return models.Address{
		AddressID:            1039341,
		PickUpInstructions:   "",
		DeliveryInstructions: "",
		City:                 "Ontario",
		CompanyName:          "RamsFord inc",
		PrimaryContactPerson: &models.Contact{
			AddressContactID: 1082586,
			Name:             "Surendra Kandel",
			Email:            "kandelsuren@gmail.com",
			FirstName:        "Surendra",
			IsPrimary:        true,
			LastName:         "Kandel",
			Phone:            "7135162836",
			Position:         nil,
			Ext:              nil,
		},
		AddressContacts: []*models.Contact{
			{
				AddressContactID: 1082586,
				Name:             "Surendra Kandel",
				FirstName:        "Surendra",
				LastName:         "Kandel",
				Phone:            "7135162836",
				Ext:              nil,
				Email:            "kandelsuren@gmail.com",
				Position:         nil,
				IsPrimary:        true,
			},
		},
		Country:          "USA",
		CountryCode:      "US",
		DeliveryFromTime: "8:00:00 AM",
		DeliveryToTime:   "5:30:00 PM",
		IsCanada:         false,
		PostalCode:       "91762",
		ShippingFromTime: "8:00:00 AM",
		ShippingToTime:   "5:30:00 PM",
		State:            "California",
		StateCode:        "CA",
		StreetLine1:      "1131 West 6th Street",
		StreetLine2:      "",
		Location: &models.Location{
			AddressID: 1039341,
			Name:      "RamsFord inc",
			IsDefault: false,
		},
		CommercialType:      nil,
		AddressAccessorials: []models.AddressAccessorial{},
		Lat:                 34.0845839,
		Long:                -117.6719935,
		HasSaiaIntegration:  false,
	}
}
