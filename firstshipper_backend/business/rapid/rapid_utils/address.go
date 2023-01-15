package rapid_utils

import (
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func FixOriginAddress(location *v1.Location, newLocation *models.Address) {
	// "addressId": 1034341,
	// "postalCode": "90013",
	newLocation.AddressID = 1034341
	newLocation.PostalCode = "91762"
	newLocation.CommercialType = &models.CommercialType{
		AccessorialID: 72,
		Name:          "Business",
	}
}

func FixDestinationAddress(location *v1.Location, newLocation *models.Address) {
	newLocation.PostalCode = location.Address.ZipCode
	newLocation.CommercialType = &models.CommercialType{
		AccessorialID: 72,
		Name:          "Business",
	}
}
func FixOriginShippingDetails(origin *v1.Location) (*models.OriginShippingDetails, error) {
	rapidOrigin := &models.OriginShippingDetails{
		Address: models.Address{
			AddressID:  1039341,
			PostalCode: origin.Address.ZipCode,
			CommercialType: &models.CommercialType{
				AccessorialID: 72,
				Name:          "Business",
			},
			AddressAccessorials: []models.AddressAccessorial{},
		},
	}
	return rapidOrigin, nil
}
func FixDestinationShippingDetails(baseQuote *v1.Location) (*models.DestinationShippingDetails, error) {
	origin := &models.DestinationShippingDetails{
		Address: models.Address{
			PostalCode: baseQuote.Address.ZipCode,
			CommercialType: &models.CommercialType{
				AccessorialID: 72,
				Name:          "Business",
			},
			AddressAccessorials: []models.AddressAccessorial{},
		},
	}
	return origin, nil
}
