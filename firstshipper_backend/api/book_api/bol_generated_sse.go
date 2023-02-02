package book_api

import (
	"fmt"

	"github.com/labstack/echo/v5"
)

func (bol *Booking) EchoInformBOlGeneratedHandler(ctx echo.Context) error {
	quoteId := ctx.PathParam("bidId")
	fmt.Print(quoteId)
	return nil
}
