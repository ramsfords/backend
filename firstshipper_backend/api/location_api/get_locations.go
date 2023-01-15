package location_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func GinLocation() {

}

func (loc Location) EchoGetLocations(ctx echo.Context) error {
	req := &v1.Id{}
	err := ctx.Bind(req)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := loc.GetLocations(ctx.Request().Context(), req)
	if err == nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)

}

func (loc Location) GetLocations(ctx context.Context, req *v1.Id) (*v1.Locations, error) {
	locations, err := loc.services.GetLocations(ctx, req.BusinessId)
	if err != nil {
		loc.services.Logger.Error("GetAllLocations GetAllLocation : error in getting all locations: %s", err)
		return nil, errs.ErrStoreInternal
	}
	locationsRes := []*v1.Location{}
	for _, loc := range locations {
		locationsRes = append(locationsRes, loc)
	}
	return &v1.Locations{
		Locations: locationsRes,
	}, nil
}
