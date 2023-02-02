package book_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Booking struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	bol := Booking{
		services: services,
	}
	protectedBolGroup := echo.Group("/book")
	protectedBolGroup.GET("/:bidId", bol.EchoGetBooking)
	protectedBolGroup.POST("", bol.EchoCreateBooking)
	protectedBolGroup.GET("/inform_bol_gen", bol.EchoInformBOlGeneratedHandler)
}
