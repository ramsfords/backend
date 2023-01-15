package menu_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/models"
	"github.com/ramsfords/backend/menuloom_backend/api/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (api menuApi) createItems(ctx echo.Context) error {
	data := &v1.Categories{}
	err := ctx.Bind(data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	record, _ := ctx.Get(apis.ContextAuthRecordKey).(*models.Record)
	fmt.Println(record)
	restaurantId, ok := record.Get("restaurantIds").(string)
	if !ok {
		return ctx.NoContent(http.StatusBadRequest)
	}
	for _, v := range data.Categories {
		v.RestaurantId = restaurantId
	}
	res, err := api.CreateCategories(ctx.Request().Context(), data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api menuApi) CreateItems(ctx context.Context, data *v1.Categories) (*v1.ItemResponse, error) {
	valid := data.Validate()
	if valid != nil {
		return nil, errs.ErrMissingData
	}
	err := api.services.Repository.CreateCategories(ctx, data.Categories[0].RestaurantId, data.Categories)
	if err != nil {
		return nil, err
	}
	return &v1.ItemResponse{
		Success: true,
		Message: "Menu created successfully",
	}, nil
}
