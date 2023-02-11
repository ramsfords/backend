package user_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/menuloom_backend/services"
)

type UserApi struct {
	services *services.Services
}
type AuthToken struct {
	AuthToken string `json:"authToken"`
}

func (user UserApi) EchoLogout(ctx echo.Context) error {
	auth := AuthToken{}
	if err := ctx.Bind(&auth); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.NoContent(http.StatusOK)
}

func New(echo *echo.Group, services *services.Services) {
	userApi := UserApi{
		services: services,
	}
	grp := echo.Group("/user")
	grp.POST("/logout", userApi.EchoLogout)
}
