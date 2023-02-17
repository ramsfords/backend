package location_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
)

type Location struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Group) {
	loc := Location{
		services: services,
	}
	locationGrp := echo.Group("/location", mid.Protected(services))
	locationGrp.POST("", loc.EchoCreateLocation)
	locationGrp.POST("/add", loc.EchoAddLocation)
	locationGrp.PATCH("", loc.EchoUpdateLocation)
}
