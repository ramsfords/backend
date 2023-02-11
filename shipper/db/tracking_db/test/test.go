package test

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type TrackingDb struct {
	*configs.Config
	dynamo.DB
}

var trackingDb TrackingDb

func init() {
	conf := configs.GetConfig()
	db := dynamo.New(conf)
	trackingDb = TrackingDb{
		Config: conf,
		DB:     db,
	}
}
