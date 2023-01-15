package tracking_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (trac Tracking) GinGetTracking(ctx echo.Context) error {
	req := &v1.Id{}
	err := ctx.Bind(req)
	if err != nil {
		return errs.ErrInvalidInputs
	}

	res, err := trac.TrackAShipment(ctx.Request().Context(), req)

	if err != nil {
		err = errs.ErrLocationUpdationFailed
		return err
	}
	return ctx.JSON(http.StatusOK, res)

}
func (trac Tracking) TrackAShipment(ctx context.Context, id *v1.Id) (*v1.ShipmentStatus, error) {
	return nil, nil

}
