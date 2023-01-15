package utils

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/logger"
)

func RemoveTokenFormCloudflareKv(conf *configs.Config, logger logger.Logger, cloudFlareCli *cloudflare.API, tokenKey string) error {
	res, err := cloudFlareCli.DeleteWorkersKVEntry(context.Background(), cloudflare.AccountIdentifier(conf.SitesSettings.Menuloom.CloudFlareConfig.AccountId), cloudflare.DeleteWorkersKVEntryParams{
		NamespaceID: conf.SitesSettings.Menuloom.CloudFlareConfig.NamespaceID,
		Key:         tokenKey,
	})
	if err != nil {
		logger.Errorf("error writing to cloudflare: %v", err)
		return err
	}
	if err != nil || len(res.Errors) > 0 {
		logger.Errorf("error writing to cloudflare: %v", err)
		return err
	}
	return nil

}
