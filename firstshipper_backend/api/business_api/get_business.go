package business_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) GinGetAllBusinesses(ctx echo.Context) error {
	req := v1.Id{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := business.GetBusiness(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusAccepted, res)

}

func (business Business) GetBusiness(ctx context.Context, req *v1.Id) (*v1.Business, error) {
	err := req.Validate()
	if err != nil {
		business.services.Logger.Errorf("GetBusiness Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	businessData, err := business.services.GetBusiness(ctx, req.Id)
	if err != nil {
		business.services.Logger.Errorf("GetAllBusinesses : error in getting all businesses: %s", err)
		return nil, errs.ErrStoreInternal
	}

	return businessData, nil
}
