package menu_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (menu menuApi) getMenu(ctx echo.Context) error {
	id := ctx.PathParam("id")
	data, err := menu.services.Repository.GetMenu(context.Background(), id)
	if err != nil {
		ctx.NoContent(http.StatusBadRequest)

	}
	return ctx.JSON(http.StatusOK, data)
}
