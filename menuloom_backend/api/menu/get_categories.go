package menu_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (api menuApi) getCategories(ctx echo.Context) error {
	id := ctx.PathParam("id")
	res, err := api.GetCategories(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api menuApi) GetCategories(ctx context.Context, id string) (data []*v1.Category, err error) {
	categories, err := api.services.Repository.GetCategories(ctx, id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
