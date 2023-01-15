package bol_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Bol struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	bol := Bol{
		services: services,
	}
	protectedBolGroup := echo.Group("/bol/v1/")
	protectedBolGroup.GET(":id", bol.GinGetBOL)
	protectedBolGroup.POST("", bol.GinCreateBOL)
}
