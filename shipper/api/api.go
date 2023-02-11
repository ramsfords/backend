package api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/shipper/api/bol_api"
	"github.com/ramsfords/backend/shipper/api/booking_api"
	"github.com/ramsfords/backend/shipper/api/business_api"
	"github.com/ramsfords/backend/shipper/api/location_api"
	"github.com/ramsfords/backend/shipper/api/quote_api"
	"github.com/ramsfords/backend/shipper/api/tracking_api"
	"github.com/ramsfords/backend/shipper/api/user_api"
	"github.com/ramsfords/backend/shipper/services"
)

func SetUpAPi(firstShipperGrp *echo.Echo, services *services.Services) {
	// grp.Use(apis.RequireAdminOrRecordAuth())
	firstShipperGrp.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "pong",
			"status":  "ok",
			"code":    200,
		})
	})

	bol_api.New(services, firstShipperGrp)
	location_api.New(services, firstShipperGrp)
	quote_api.New(services, firstShipperGrp)
	tracking_api.New(services, firstShipperGrp)
	user_api.New(services, firstShipperGrp)
	business_api.New(services, firstShipperGrp)
	booking_api.New(services, firstShipperGrp)
}
