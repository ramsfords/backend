package utils

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/logger"
)

func RemoveTokenFormCloudflareKv(conf *configs.Config, cloudFlareCli *cloudflare.API, tokenKey string) error {
	res, err := cloudFlareCli.DeleteWorkersKVEntry(context.Background(), cloudflare.AccountIdentifier(conf.CloudFlareConfig.AccountId), cloudflare.DeleteWorkersKVEntryParams{
		NamespaceID: conf.CloudFlareConfig.NamespaceId,
		Key:         tokenKey,
	})
	if err != nil {
		logger.Error(err, "error writing to cloudflare")
		return err
	}
	if err != nil || len(res.Errors) > 0 {
		logger.Error(err, "error writing to cloudflare: %v")
		return err
	}
	return nil

}
