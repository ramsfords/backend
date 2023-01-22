package quote_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (qt Quote) EchoGetBidsByQuoteId(ctx echo.Context) error {
	quoteID := ctx.PathParam("quoteId")
	businessID := ctx.PathParam("businessId")
	if quoteID == "" || businessID == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newCtx := ctx.Request().Context()
	res, err := qt.services.GetBidsByQuoteId(newCtx, quoteID, businessID)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}
