package test

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type QuoteDb struct {
	*configs.Config
	dynamo.DB
}

var quoteDb QuoteDb

func init() {
	conf := configs.GetConfig()
	db := dynamo.New(conf)
	quoteDb = QuoteDb{
		Config: conf,
		DB:     db,
	}
}
