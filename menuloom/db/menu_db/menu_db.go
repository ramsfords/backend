package menu_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type MenuDb struct {
	dynamo.DB
	*configs.Config
}
