package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoCreateLocation(ctx echo.Context) error {
	req := v1.Location{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := loc.CreateLocation(ctx.Request().Context(), &req)

	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	err = ctx.JSON(http.StatusOK, res)
	return nil
}

func (loc Location) CreateLocation(ctx context.Context, locationReq *v1.Location) (*v1.Ok, error) {
	err := locationReq.Validate()
	if err != nil {
		loc.services.Logger.Error("CreateLocation Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = loc.services.Db.SaveLocation(ctx, locationReq.BusinessId, locationReq)
	if err != nil {
		loc.services.Logger.Error("CreateLocation InsertLocation : error in inserting location into the database: %s", err)
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
