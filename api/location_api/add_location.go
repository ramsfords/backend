package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoAddLocation(ctx echo.Context) error {
	req := v1.Location{}
	err := ctx.Bind(&req)
	if err != nil {
		err = errs.ErrInvalidInputs
		ctx.NoContent(http.StatusBadRequest)
	}
	res, err := loc.AddLocation(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = ctx.JSON(http.StatusOK, res)
	return err
}

func (loc Location) AddLocation(ctx context.Context, locationReq *v1.Location) (*v1.Ok, error) {
	err := locationReq.Validate()
	if err != nil {
		logger.Error(err, "AddLocation Validate : req data validation failed")
		return nil, errs.ErrInvalidInputs
	}
	err = loc.services.Db.SaveLocation(ctx, locationReq.BusinessId, locationReq)
	if err != nil {
		logger.Error(err, "AddLocation InsertLocation : error in inserting location into the database")
		return nil, errs.ErrLocationCreationFailed
	}

	return &v1.Ok{Success: true, Code: 201, Message: "Success"}, nil
}
