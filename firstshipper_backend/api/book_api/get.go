package book_api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func (booking *Booking) EchoGetBooking(ctx echo.Context) error {
	bidId := ctx.PathParam("bidId")
	if len(bidId) < 5 {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	quoteId := strings.Split(bidId, "-")[0]
	ctxx := ctx.Request().Context()
	qtReq, err := booking.services.GetBooking(ctxx, quoteId)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	fmt.Print(qtReq)
	if len(qtReq.BookingInfo.CarrierProNumber) < 4 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, qtReq)

}
