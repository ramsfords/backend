package cloudflare

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/ramsfords/backend/configs"
)

type Cloudflare struct {
	*cloudflare.API
	AccountId   string
	NamespaceID string
}

func New(conf configs.CloudFlareConfig) *Cloudflare {
	api, err := cloudflare.New(conf.ApiKey, conf.Email)
	if err != nil {
		panic(fmt.Sprintf("error in cloudflare: %v", err))
	}
	cloudFlare := &Cloudflare{
		API:         api,
		AccountId:   conf.AccountId,
		NamespaceID: conf.NamespaceId,
	}

	return cloudFlare
}
