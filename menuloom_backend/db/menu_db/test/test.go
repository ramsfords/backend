package test

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

var db, conf = getDB()

func getDB() (db dynamo.DB, conf configs.Config) {
	confs := configs.GetConfig()
	dbClient := dynamo.New(confs)
	return dbClient, *confs
}
func GetProdDb() dynamo.DB {
	confs := configs.GetConfig()
	dbClient := dynamo.NewProdDb(confs)
	return dbClient
}
