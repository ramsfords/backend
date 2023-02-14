package tracking_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type LocationDb struct {
	*configs.Config
	dynamo.DB
}
