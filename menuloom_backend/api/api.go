package api

import (
	"github.com/labstack/echo/v5"
	menu_api "github.com/ramsfords/backend/menuloom_backend/api/menu"
	restaurant_api "github.com/ramsfords/backend/menuloom_backend/api/restaurant"
	"github.com/ramsfords/backend/menuloom_backend/api/user_api"
	"github.com/ramsfords/backend/menuloom_backend/api/validate_api"
	"github.com/ramsfords/backend/menuloom_backend/services"
)

func SetUpAPi(echo *echo.Echo, services services.Services) {
	menu_api.New(echo, services)
	restaurant_api.New(echo, services)
	validate_api.New(echo, services)
	user_api.New(echo, services)
}
