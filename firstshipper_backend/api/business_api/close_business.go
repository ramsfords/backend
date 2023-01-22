package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) GinCloseBusiness(ctx echo.Context) error {
	req := v1.Id{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := business.CloseBusiness(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)

}

func (business Business) CloseBusiness(ctx context.Context, req *v1.Id) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		business.services.Logger.Errorf("CloseBusiness Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}

	//TO DO

	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
