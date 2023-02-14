package cloudflare

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/ramsfords/backend/configs"
)

func New(conf configs.CloudFlareConfig) *cloudflare.API {
	api, err := cloudflare.New(conf.ApiKey, conf.Email)
	if err != nil {
		panic(fmt.Sprintf("error in cloudflare: %v", err))
	}
	return api
}
