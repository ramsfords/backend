package user_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type UserDb struct {
	dynamo.DB
	*configs.Config
}
