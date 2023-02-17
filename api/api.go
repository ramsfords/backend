package api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/auth"
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
	api := engine.Group("/api")
	auth.New(services, api)
	bol_api.New(services, api)
	location_api.New(services, api)
	quote_api.New(services, api)
	tracking_api.New(services, api)
	user_api.New(services, api)
	business_api.New(services, api)
	booking_api.New(services, api)

}
