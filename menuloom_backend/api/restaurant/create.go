package restaurant_api

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/menuloom_backend/api/errs"

	v1 "github.com/ramsfords/types_gen/v1"
)

func (api restaurantApi) createRestaurant(ctx echo.Context) error {
	data := &v1.CreateRestaurantData{}
	err := ctx.Bind(data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := api.CreateRestaurant(ctx.Request().Context(), data)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, res)
}
func (api restaurantApi) CreateRestaurant(ctx context.Context, data *v1.CreateRestaurantData) (*v1.RestaurantResponse, error) {
	data.RestaurantWebUrl = strings.Split(data.RestaurantWebUrl, ":")[0]
	valid := data.Validate()
	if valid != nil {
		return nil, errs.ErrMissingData
	}
	// resS3, err := api.services.S3Client.Upload(data.RestaurantWebUrl)
	// if err != nil {
	// 	return nil, err
	// }
	// s3Location := strings.Split(resS3.Location, "/sample")[0]
	// data.RestaurantS3StaticProdUrl = s3Location
	// _, err = api.services.S3Client.UploadSampleImage(data.RestaurantWebUrl)
	// if err != nil {
	// 	return nil, err
	// }
	// data.S3StaticImagesUrl = s3Location + "/images"
	err := api.services.Db.CreateRestaurant(ctx, data)
	if err != nil {
		return nil, err
	}

	// api.services.Db.CreateUser(ctx, newUser)

	return &v1.RestaurantResponse{
		Id:      data.Id,
		Message: "Menu created successfully",
	}, nil

}
