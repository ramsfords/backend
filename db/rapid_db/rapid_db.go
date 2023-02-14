package rapid_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type RapidDb struct {
	*configs.Config
	dynamo.DB
}
