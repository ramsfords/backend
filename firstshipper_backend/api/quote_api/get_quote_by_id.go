package quote_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (qt Quote) EchoGetQuoteById(ctx echo.Context) error {
	id := ctx.PathParam("quoteId")
	if id == "" {
		return ctx.NoContent(http.StatusBadRequest)
	}
	res, err := qt.services.GetQuoteByQuoteId(ctx.Request().Context(), "2700", "id")
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}
