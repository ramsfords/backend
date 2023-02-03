package quote_api

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (quote Quote) EchoGetQuoteWithBidByBidId(ctx echo.Context) error {
	bidId := ctx.PathParam("bidId")
	if len(bidId) < 5 || bidId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	quoteId := strings.Split(bidId, "-")[0]

	// admin, _ := ctx.Get(ContextAdminKey).(*models.Admin)
	// record, _ := ctx.Get(ContextAuthRecordKey).(*models.Record)

	// if admin == nil && record == nil {
	// 	return ctx.NoContent(http.StatusUnauthorized)
	// }
	// optCollectionNames := []string{"firstshipper_auth"}

	// if record != nil && len(optCollectionNames) > 0 && !list.ExistInSlice(record.Collection().Name, optCollectionNames) {
	// 	return ctx.NoContent(http.StatusUnauthorized)
	// }

	// ctxs := ctx.Request().Context()
	quoteWithBids, err := quote.services.GetBidsWithQuoteByQuoteId(context.Background(), quoteId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	quoteWithBids.Bid = BidByBidId(quoteWithBids.Bids, bidId)
	quoteWithBids.Bids = nil
	return ctx.JSON(http.StatusOK, quoteWithBids)

}
func BidByBidId(bids []*v1.Bid, bidId string) *v1.Bid {
	for _, bid := range bids {
		if bid.BidId == bidId {
			return bid
		}
	}
	return &v1.Bid{}
}
