package tracking_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (trac Tracking) EchoGetTracking(ctx echo.Context) error {
	req := ctx.PathParam("shipmentId")
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
func (trac Tracking) TrackAShipment(ctx context.Context, id string) (*v1.ShipmentTracking, error) {
	return nil, nil

}
