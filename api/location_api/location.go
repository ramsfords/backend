package location_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/services"
)

type Location struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	loc := Location{
		services: services,
	}
	locationGrp := echo.Group("/location")
	locationGrp.POST("", loc.EchoCreateLocation)
	locationGrp.POST("/add", loc.EchoAddLocation)
	locationGrp.PATCH("", loc.EchoUpdateLocation)
}
