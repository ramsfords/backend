package restaurant_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (menu restaurantApi) get(ctx echo.Context) error {
	id := ctx.PathParam("id")
	data, err := menu.services.Db.GetRestaurant(context.Background(), id)
	if err != nil {
		ctx.NoContent(http.StatusBadRequest)

	}
	return ctx.JSON(http.StatusOK, data)
}
