package quote_api

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func (quote Quote) EchoGetQuoteWithBidsByQuoteId(ctx echo.Context) error {
	quoteId := ctx.QueryParam("quoteId")
	if len(quoteId) < 5 || quoteId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	businessId := ctx.QueryParam("businessId")
	if len(businessId) < 5 || businessId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	// admin, _ := ctx.Get(ContextAdminKey).(*models.Admin)
	// record, _ := ctx.Get(ContextAuthRecordKey).(*models.Record)

	// if admin == nil && record == nil {
	// 	return ctx.NoContent(http.StatusUnauthorized)
	// }
	// optCollectionNames := []string{"firstshipper_auth"}

	// if record != nil && len(optCollectionNames) > 0 && !list.ExistInSlice(record.Collection().Name, optCollectionNames) {
	// 	return ctx.NoContent(http.StatusUnauthorized)
	// }

	ctxs := ctx.Request().Context()
	bids, err := quote.services.Db.GetBidsWithQuoteByQuoteId(ctxs, quoteId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, bids)

}
