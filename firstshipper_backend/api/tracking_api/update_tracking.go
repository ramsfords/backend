package tracking_api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v5"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (trac Tracking) GinUpdateTracking(ctx echo.Context) error {
	req := &v1.ShipmentStatus{}
	err := ctx.Bind(&req)
	if err != nil {
		err = errs.ErrInvalidInputs
		return err
	}

	res, err := trac.UpdateShipmentTracking(ctx.Request().Context(), req)

	if err != nil {
		err = errs.ErrLocationUpdationFailed
		return err
	}
	err = ctx.JSON(http.StatusOK, res)
	return err
}
func (trac Tracking) UpdateShipmentTracking(ctx context.Context, tracking *v1.ShipmentStatus) (*v1.Ok, error) {
	return nil, nil

}
