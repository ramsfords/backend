package rapid

import (
	"github.com/ramsfords/backend/business/core/client"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/vault"
)

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
