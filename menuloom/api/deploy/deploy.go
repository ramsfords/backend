package deploy

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/menuloom/services"
)

type Deploy interface {
	create(ctx echo.Context) error
	update(ctx echo.Context) error
	delete(ctx echo.Context) error
}
type deploy struct {
	services services.Services
}

func New(echo *echo.Echo, services services.Services) {
	var deployHandler deploy = deploy{
		services: services,
	}
	grp := echo.Group("/deploy")
	grp.POST("create", deployHandler.create)
	grp.PUT(":id", deployHandler.update)
	grp.DELETE(":id", deployHandler.delete)
}
