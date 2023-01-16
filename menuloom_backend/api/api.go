package api

import (
	"github.com/labstack/echo/v5"
	menu_api "github.com/ramsfords/backend/menuloom_backend/api/menu"
	restaurant_api "github.com/ramsfords/backend/menuloom_backend/api/restaurant"
	"github.com/ramsfords/backend/menuloom_backend/api/user_api"
	"github.com/ramsfords/backend/menuloom_backend/api/validate_api"
	"github.com/ramsfords/backend/menuloom_backend/services"
)

func SetUpAPi(menuLoomGrp *echo.Echo, services services.Services) {
	grp := menuLoomGrp.Group("/menuloom")
	grp.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(200, "Pong form Menuloom")
	})
	menu_api.New(grp, services)
	restaurant_api.New(grp, services)
	validate_api.New(grp, services)
	user_api.New(grp, services)
}
