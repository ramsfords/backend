package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoUpdateLocation(ctx echo.Context) error {
	req := v1.Address{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := loc.UpdateLocation(ctx.Request().Context(), &req)

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)

}

func (loc Location) UpdateLocation(ctx context.Context, addressReq *v1.Address) (*v1.Ok, error) {
	err := addressReq.Validate()
	if err != nil {
		loc.services.Logger.Error("UpdateLocation Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = loc.services.Db.UpdateLocation(ctx, addressReq.BusinessId, addressReq)
	if err != nil {
		loc.services.Logger.Error("UpdateLocation : error in updating location into the database: %s", err)
		return nil, errs.ErrLocationUpdationFailed
	}
	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
