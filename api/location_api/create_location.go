package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
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
		logger.Error(err, "CreateLocation Validate : req data validation failed")
		return nil, errs.ErrInputDataNotValid
	}
	err = loc.services.Db.SaveLocation(ctx, locationReq.BusinessId, locationReq)
	if err != nil {
		logger.Error(err, "CreateLocation InsertLocation : error in inserting location into the database")
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    200,
		Message: "Success",
	}
	return res, nil
}
