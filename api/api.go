package api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/authapi"
	"github.com/ramsfords/backend/api/bol_api"
	"github.com/ramsfords/backend/api/booking_api"
	"github.com/ramsfords/backend/api/business_api"
	"github.com/ramsfords/backend/api/location_api"
	"github.com/ramsfords/backend/api/quote_api"
	"github.com/ramsfords/backend/api/tracking_api"
	"github.com/ramsfords/backend/api/user_api"
	"github.com/ramsfords/backend/services"
)

func SetUpAPi(engine *echo.Echo, services *services.Services) {
	authapi.New(services, engine)
	bol_api.New(services, engine)
	location_api.New(services, engine)
	quote_api.New(services, engine)
	tracking_api.New(services, engine)
	user_api.New(services, engine)
	business_api.New(services, engine)
	booking_api.New(services, engine)

}
