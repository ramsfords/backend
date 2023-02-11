package menu_api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"

	"github.com/ramsfords/backend/menuloom_backend/api/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (api menuApi) createCategory(ctx echo.Context) error {
	data := &v1.Category{}
	// values, err := ctx.FormValues()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(values)
	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}
	intRank, err := strconv.ParseInt(form.Value["rank"][0], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	data.Name = strings.ToLower(form.Value["name"][0])
	data.Description = strings.ToLower(form.Value["description"][0])
	data.RestaurantId = strings.ToLower(form.Value["restaurantId"][0])
	data.Pk = strings.ToLower(form.Value["pk"][0])
	data.Type = "category"
	data.ServingTime = v1.ServingTime(v1.ServingTime_value[form.Value["servingTime"][0]])
	data.Rank = int32(intRank)
	if err != nil {
		return err
	}
	fmt.Println(form)
	for _, file := range form.File {
		data.Images = make([]*v1.Image, len(file))
		for index, f := range file {
			fileXExt := strings.Split(f.Filename, ".")[0]
			folderName := strings.Split(data.RestaurantId, ".")[0]
			fileData, err := f.Open()
			if err != nil {
				fmt.Println(err)
			}
			res, err := api.services.Cloudinery.UploadImageFromUrl(fileData, fileXExt, folderName)
			if err != nil {
				fmt.Println(err)
			}
			data.Images[index] = &v1.Image{
				Url:       res.URL,
				SecureUrl: res.SecureURL,
				PublicId:  res.PublicID,
				AssetId:   res.AssetID,
			}

		}

	}

	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := api.CreateCategory(ctx.Request().Context(), data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api menuApi) CreateCategory(ctx context.Context, data *v1.Category) (*v1.ItemResponse, error) {
	valid := data.Validate()
	if valid != nil {
		return nil, errs.ErrMissingData
	}
	err := api.services.Db.CreateCategory(ctx, data, data.RestaurantId)
	if err != nil {
		return nil, err
	}
	return &v1.ItemResponse{
		Success: true,
		Message: "Menu created successfully",
	}, nil
}
