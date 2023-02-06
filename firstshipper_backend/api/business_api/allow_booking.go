package business_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type AllowBooking struct {
	Allow      bool   `json:"allowBooking"`
	BusinessId string `json:"businessId"`
}

func (business Business) AllowBooking(ctx echo.Context) error {
	allowBooking := &AllowBooking{}
	err := ctx.Bind(allowBooking)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	if len(allowBooking.BusinessId) < 3 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newContext := ctx.Request().Context()
	_, err = business.services.UpdateAllowBooking(newContext, allowBooking.BusinessId, allowBooking.Allow)
	if err != nil {
		business.services.Logger.Errorf("adding address to database failded: %s", err)
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, allowBooking)
}
