package quote_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (qt Quote) EchoGetBidByBidId(ctx echo.Context) error {
	quoteID := ctx.PathParam("quoteId")
	businessID := ctx.PathParam("businessId")
	bidId := ctx.PathParam("bidId")
	if quoteID == "" || businessID == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newCtx := ctx.Request().Context()
	res, err := qt.services.Db.GetBidByBidID(newCtx, businessID, quoteID, bidId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}
