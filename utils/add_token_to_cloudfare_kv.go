package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/pocketbase/pocketbase/core"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func AddTokenToCloudFlareKV(conf *configs.Config, logger logger.Logger, cloudFlareCli *cloudflare.API) func(e *core.RecordAuthEvent) error {
	return func(e *core.RecordAuthEvent) error {
		// e record json
		if e.Record.Collection().Name != "firstshipper_auth" {
			return nil
		}
		originalData := e.Record.ColumnValueMap()
		fmt.Println(originalData)
		jsonRecord, err := json.Marshal(originalData)
		if err != nil {
			logger.Errorf("error marshalling record: %v", err)
		}
		newData := v1.User{}
		err = json.Unmarshal(jsonRecord, &newData)
		if err != nil {
			logger.Errorf("error unmarshalling record: %v", err)
		}
		newData.Token = e.Token
		newUserBytes, err := json.Marshal(newData)
		if err != nil {
			logger.Errorf("error marshalling newUser: %v", err)
		}
		res, err := cloudFlareCli.WriteWorkersKVEntry(context.Background(), cloudflare.AccountIdentifier(conf.SitesSettings.Menuloom.CloudFlareConfig.AccountId), cloudflare.WriteWorkersKVEntryParams{
			NamespaceID: conf.SitesSettings.Menuloom.CloudFlareConfig.NamespaceID,
			Key:         e.Token,
			Value:       newUserBytes,
		})
		if err != nil {
			logger.Errorf("error writing to cloudflare: %v", err)
		}
		if err != nil || len(res.Errors) > 0 {
			logger.Errorf("error writing to cloudflare: %v", err)
		}
		return nil
	}
}
