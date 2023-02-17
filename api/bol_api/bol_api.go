package bol_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
)

type Bol struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Group) {
	bol := Bol{
		services: services,
	}
	protectedBolGroup := echo.Group("/bol", mid.Protected(services))
	protectedBolGroup.GET("/hello", bol.EchoGetBOL)

}
