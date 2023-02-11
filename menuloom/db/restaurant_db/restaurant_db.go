package restaurant_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type RestaurantDb struct {
	dynamo.DB
	*configs.Config
}
