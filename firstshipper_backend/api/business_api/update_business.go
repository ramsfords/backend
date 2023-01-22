package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) GinUpdateBusiness(ctx echo.Context) error {
	req := v1.Business{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := business.UpdateBusiness(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusAccepted, res)

}

func (business Business) UpdateBusiness(ctx context.Context, businessReq *v1.Business) (*v1.Ok, error) {
	err := businessReq.Validate()
	if err != nil {
		business.services.Logger.Errorf("UpdateBusiness Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}

	err = business.services.UpdateBusiness(ctx, businessReq.BusinessId, *businessReq)
	if err != nil {
		business.services.Logger.Errorf("UpdateBusiness : error in updating business into the database: %s", err)
		return nil, errs.ErrBusinessUpdationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
