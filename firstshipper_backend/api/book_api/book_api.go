package book_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Booking struct {
	services *services.Services
}

func New(services *services.Services, echoClient *echo.Echo) {
	bol := Booking{
		services: services,
	}
	protectedBolGroup := echoClient.Group("/book")
	protectedBolGroup.GET("/hello", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello from book, World!")
	})
	protectedBolGroup.GET("/bid/:bidId", bol.EchoGetBooking)
	protectedBolGroup.POST("", bol.EchoCreateBooking)
	protectedBolGroup.GET("/inform_bol_gen", bol.EchoInformBOlGeneratedHandler)
}
