package business_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type BusinessDb struct {
	*configs.Config
	dynamo.DB
}
