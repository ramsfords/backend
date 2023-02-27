package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) EchoCloseBusiness(ctx echo.Context) error {
	res, err := business.CloseBusiness(ctx.Request().Context(), " &req")
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)

}

func (business Business) CloseBusiness(ctx context.Context, req string) (*v1.Ok, error) {

	//TO DO

	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
