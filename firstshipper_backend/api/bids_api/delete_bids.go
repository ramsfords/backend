package bids_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (bol Bids) EchoDeleteBids(ctx echo.Context) error {
	quoteId := ctx.QueryParam("quoteId")
	if len(quoteId) < 5 || quoteId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	businessId := ctx.QueryParam("businessId")
	if len(businessId) < 5 || businessId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	ctx.Request().Header.Set("Cache-Control", "max-age=604800")
	return ctx.JSON(http.StatusOK, "application/pdf")

}
