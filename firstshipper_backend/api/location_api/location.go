package location_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Location struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Group) {
	loc := Location{
		services: services,
	}
	locationGrp := echo.Group("/location")
	locationGrp.DELETE("/:id", loc.EchoDeleteLocation)
	locationGrp.DELETE("/deleteall", loc.EchoDeleteLocations)
	locationGrp.POST("", loc.EchoCreateLocation)
	locationGrp.POST("/add", loc.EchoAddLocation)
	locationGrp.GET("", loc.EchoGetLocations)
	locationGrp.GET("/:id", loc.EchoGetLocation)
	locationGrp.PATCH("", loc.EchoUpdateLocation)
}
