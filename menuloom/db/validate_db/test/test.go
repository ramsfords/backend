package test

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

var conf = configs.GetConfig()
var db = dynamo.New(conf)

func init() {
	conf = configs.GetConfig()
	db = dynamo.New(conf)
}
