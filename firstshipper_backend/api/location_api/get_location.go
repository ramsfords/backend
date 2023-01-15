package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (loc Location) EchoGetLocation(ctx echo.Context) error {
	req := v1.Id{}
	err := ctx.Bind(&req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	res, err := loc.GetLocation(ctx.Request().Context(), &req)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (loc Location) GetLocation(ctx context.Context, locReq *v1.Id) (*v1.Location, error) {
	err := locReq.Validate()
	if err != nil {
		loc.services.Logger.Error("GetLocation Validate : req data validation failed: %s", err)
		return nil, errs.ErrInputDataNotValid
	}
	_, err = loc.services.GetLocation(ctx, locReq.BusinessId, locReq.Id)
	if err != nil {
		loc.services.Logger.Error("GetAllLocations GetAllLocation : error in getting allproviderations: %s", err)
		return nil, errs.ErrStoreInternal
	}

	return nil, nil
}
