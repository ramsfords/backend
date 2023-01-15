package rapid

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/client"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/foundations/vault"
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
//     }
// ]

type Rapid struct {
	configs.RapidShipLTL `json:"rapid_ship_ltl"`
	*vault.Vault
	*client.HttpClient
}

func New() *Rapid {
	return &Rapid{
		configs.GetConfig().SitesSettings.FirstShipper.RapidShipLTL,
		vault.New(),
		client.New(),
	}

}

func (rapid Rapid) GetCapacityProviderAccountGroup() models.CapacityProviderAccountGroup {
	// "capacityProviderAccountGroup": {
	//     "code": "RamsFordinc_PROD_91762_1039341_202202111959",
	//     "accounts": [
	//         {
	//             "code": "RDFS"
	//         }
	//     ]
	// },
	return models.CapacityProviderAccountGroup{
		Code: "RamsFordinc_PROD_91762_1039341_202202111959",
		Accounts: []models.Account{
			{
				Code: "RDFS",
			},
		},
	}

}
