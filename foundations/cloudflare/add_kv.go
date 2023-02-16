package cloudflare

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
)

func (cloudFlare Cloudflare) AddTokenToCloudFlareKV(key string, value string) error {
	_, err := cloudFlare.WriteWorkersKVEntry(context.Background(), cloudflare.AccountIdentifier(cloudFlare.AccountId), cloudflare.WriteWorkersKVEntryParams{
		NamespaceID: cloudFlare.NamespaceID,
		Key:         key,
		Value:       []byte(value),
	})
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("error writing to cloudflare: %v", err))
	}
	return nil
}
