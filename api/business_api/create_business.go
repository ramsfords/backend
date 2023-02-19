package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) GinCreateBusiness(ctx echo.Context) error {
	req := v1.Business{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := business.CreateBusiness(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (business Business) CreateBusiness(ctx context.Context, req *v1.Business) (*v1.Ok, error) {
	err := req.Validate()
	if err != nil {
		business.services.Logger.Errorf("CreateBusiness Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = business.services.Db.SaveBusiness(ctx, *req, req.BusinessId)
	if err != nil {
		business.services.Logger.Errorf("CreateBusiness InsertBusiness : error in inserting business into the database: %s", err)
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
