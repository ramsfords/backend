package relic

import (
	"fmt"
	"log"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/ramsfords/backend/configs"
)

func New() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("grpcauth"),
		newrelic.ConfigEnabled(true),
		newrelic.ConfigLicense(configs.GetConfig().NewRelic.License),
	)
	//nrzap.ConfigLogger(logger.Log.Named("newrelic"))
	if err != nil {
		log.Fatal(fmt.Errorf("could not start relic %s", err))
	}
	return app

}
