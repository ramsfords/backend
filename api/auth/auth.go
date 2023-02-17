package auth

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"
)

type Auth struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Group) {
	auth := Auth{
		services: services,
	}
	protectedBolGroup := echo.Group("/auth-callback")
	protectedBolGroup.GET("/:code", auth.EchoLogin)
	protectedBolGroup.GET("/logout", auth.EchoLogout)
}
