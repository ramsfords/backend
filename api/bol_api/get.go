package bol_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (bol Bol) EchoGetBOL(ctx echo.Context) error {
	quoteId := ctx.QueryParam("quoteId")
	if len(quoteId) < 1 || quoteId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	businessId := ctx.QueryParam("businessId")
	if len(businessId) < 1 || businessId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	ctxx := ctx.Request().Context()
	qtReq, err := bol.services.Db.GetBooking(ctxx, "23122")
	if err != nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, qtReq)

}
