package cloudflare

import (
	"testing"

	"github.com/ramsfords/backend/configs"
)

func TestAddKV(t *testing.T) {
	conf := configs.GetConfig()
	cloudflare := New(conf.CloudFlareConfig)
	cloudflare.AddTokenToCloudFlareKV("test", "test")

}
