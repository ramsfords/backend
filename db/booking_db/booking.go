package booking_db

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/dynamo"
)

type BookingDb struct {
	*configs.Config
	dynamo.DB
}

func New(conf *configs.Config, db dynamo.DB) BookingDb {
	booking := BookingDb{
		conf,
		db,
	}
	return booking

}
