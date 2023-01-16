package api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/api/bol_api"
	"github.com/ramsfords/backend/firstshipper_backend/api/location_api"
	"github.com/ramsfords/backend/firstshipper_backend/api/quote_api"
	"github.com/ramsfords/backend/firstshipper_backend/api/tracking_api"
	"github.com/ramsfords/backend/firstshipper_backend/api/user_api"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

func SetUpAPi(echo *echo.Echo, services *services.Services, rapid *rapid.Rapid) {
	grp := echo.Group("/firstshipper")
	bol_api.New(services, grp)
	location_api.New(services, grp)
	quote_api.New(services, grp, rapid)
	tracking_api.New(services, grp, rapid)
	user_api.New(services, grp)
}
