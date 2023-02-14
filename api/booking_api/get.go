package booking_api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

func (bookingApi *BookingApi) EchoGetBooking(ctx echo.Context) error {
	bookingId := ctx.PathParam("bookingId")
	if len(bookingId) < 5 {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	bookingId = strings.Split(bookingId, "-")[0]
	ctxx := ctx.Request().Context()
	qtReq, err := bookingApi.services.Db.GetBooking(ctxx, bookingId)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	fmt.Print(qtReq)
	if len(qtReq.BookingInfo.CarrierProNumber) < 4 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	return ctx.JSON(http.StatusOK, qtReq)

}
