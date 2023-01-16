package tracking_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Tracking struct {
	services *services.Services
	radpid   *rapid.Rapid
}

func New(services *services.Services, echo *echo.Group, rapid *rapid.Rapid) {
	tracking := Tracking{
		services: services,
		radpid:   rapid,
	}
	protectedTrackGroup := echo.Group("/tracking/v1/")
	protectedTrackGroup.GET(":id", tracking.GinGetTracking)
	protectedTrackGroup.PATCH(":id", tracking.GinUpdateTracking)
}
