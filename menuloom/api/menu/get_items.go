package menu_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (api menuApi) getItems(ctx echo.Context) error {
	id := "pk#" + ctx.PathParam("id")
	record, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)
	fmt.Println(record)
	res, err := api.GetCategories(ctx.Request().Context(), id)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api menuApi) GetItems(ctx context.Context, id string) (data []*v1.Category, err error) {
	categories, err := api.services.Db.GetCategories(ctx, id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
