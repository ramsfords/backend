package quote_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type QuoteDb struct {
	*configs.Config
	dynamo.DB
}
