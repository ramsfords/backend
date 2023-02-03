package quote_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (qt Quote) EchoGetBidsByQuoteId(ctx echo.Context) error {
	quoteID := ctx.PathParam("quoteId")
	if len(quoteID) < 5 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newCtx := ctx.Request().Context()
	res, err := qt.services.GetBidsByQuoteId(newCtx, quoteID)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}
