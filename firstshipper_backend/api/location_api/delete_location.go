package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoDeleteLocation(ctx echo.Context) error {
	req := &v1.Id{}
	err := ctx.Bind(req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := loc.DeleteLocation(ctx.Request().Context(), req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)

}
func (loc Location) DeleteLocation(ctx context.Context, locationReq *v1.Id) (*v1.Ok, error) {
	err := locationReq.Validate()
	if err != nil {
		loc.services.Logger.Error("DeleteLocation Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	err = loc.services.DeleteLocation(ctx, locationReq.Id, locationReq.Id)
	if err != nil {
		loc.services.Logger.Error("DeleteLocationsByLocationIds : error in deleting location ID: %s", err)
		return nil, errs.ErrLocationCreationFailed
	}

	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
