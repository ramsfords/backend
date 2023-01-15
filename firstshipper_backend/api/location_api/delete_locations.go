package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoDeleteLocations(ctx echo.Context) error {
	req := v1.Ids{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)

	}

	res, err := loc.DeleteLocations(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusAccepted, res)
}
func (loc Location) DeleteLocations(ctx context.Context, locationReq *v1.Ids) (*v1.Ok, error) {
	err := locationReq.Validate()
	if err != nil {
		loc.services.Logger.Error("DeleteLocations Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	if len(locationReq.Ids) > 0 {
		for _, j := range locationReq.Ids {
			err = loc.services.DeleteLocation(ctx, j.BusinessId, j.Id)
			if err != nil {
				loc.services.Logger.Error("DeleteLocationsByLocationIds : error in deleting location IDs: %s", err)
				return nil, errs.ErrLocationCreationFailed
			}
		}
	}
	res := &v1.Ok{
		Success: true,
		Code:    204,
		Message: "Success",
	}
	return res, nil
}
