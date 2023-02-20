package authapi

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"
)

type AuthApi struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	auth := AuthApi{
		services: services,
	}
	authApiGroup := echo.Group("/auth")
	authApiGroup.POST("/:code", auth.EchoLogin)
	authApiGroup.POST("/logout", auth.EchoLogout)
	authApiGroup.POST("/signup", auth.EchoSignUp)
	authApiGroup.POST("/confirm-email", auth.ConfirmEmail)
	authApiGroup.POST("/redirect-login", auth.RedirectLogin)
}
