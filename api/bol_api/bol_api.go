package bol_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"
)

type Bol struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	bol := Bol{
		services: services,
	}
	protectedBolGroup := echo.Group("/bol")
	protectedBolGroup.GET("", bol.EchoGetBOL)

}
