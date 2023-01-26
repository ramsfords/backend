package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (business Business) UpdateDefaultPickupAddress(ctx echo.Context) error {
	address := &v1.Address{}
	//err := server.unMarshall(ctx, signUpReq)
	err := ctx.Bind(address)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	businessId := ctx.QueryParam("businessId")
	if businessId == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = business.services.UpdateLocation(ctx.Request().Context(), businessId, address)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusAccepted)
}
