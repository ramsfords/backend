package tracking_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Tracking struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	tracking := Tracking{
		services: services,
	}
	protectedTrackGroup := echo.Group("/tracking")
	protectedTrackGroup.GET("/:shipmentId", tracking.EchoGetTracking)
	protectedTrackGroup.PATCH("/:shipmentId", tracking.GinUpdateTracking)
}
