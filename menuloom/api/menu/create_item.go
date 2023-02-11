package menu_api

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/menuloom/api/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (api menuApi) createItem(ctx echo.Context) error {
	data := &v1.Item{}
	err := ctx.Bind(data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := api.CreateItem(ctx.Request().Context(), data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api menuApi) CreateItem(ctx context.Context, data *v1.Item) (*v1.ItemResponse, error) {
	valid := data.Validate()
	if valid != nil {
		return nil, errs.ErrMissingData
	}
	// restaruant, err := api.services.GetRestaurantData(ctx, data.RestaurantId)
	// if err != nil {
	// 	return nil, err
	// }
	for _, url := range data.Images {
		folderName := strings.Split(data.RestaurantId, ".")[0]
		res, err := api.services.Cloudinery.UploadImageFromUrl(url.Url, data.Name, folderName)
		if err != nil {
			return nil, err
		}
		url.Url = res.URL
		url.SecureUrl = res.SecureURL
		url.AssetId = res.AssetID
		url.PublicId = res.PublicID
	}

	err := api.services.Db.CreateItem(ctx, data, data.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &v1.ItemResponse{
		Success: true,
		Message: "Menu created successfully",
	}, nil
}
