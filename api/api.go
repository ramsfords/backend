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

func SetUpAPi(firstShipperGrp *echo.Echo, services *services.Services) {
	bol_api.New(services, firstShipperGrp)
	location_api.New(services, firstShipperGrp)
	quote_api.New(services, firstShipperGrp)
	tracking_api.New(services, firstShipperGrp)
	user_api.New(services, firstShipperGrp)
	business_api.New(services, firstShipperGrp)
	booking_api.New(services, firstShipperGrp)
	auth.New(services, firstShipperGrp)
}
