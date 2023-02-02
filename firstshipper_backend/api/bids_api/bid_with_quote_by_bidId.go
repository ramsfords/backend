package bids_api

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (bid Bids) EchoGetBidWithQuoteByBidId(ctx echo.Context) error {
	bidId := ctx.QueryParam("bidId")
	if len(bidId) < 5 || bidId == "" {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	quoteId := strings.Split(bidId, "-")[0]
	if len(quoteId) < 4 {
		return ctx.NoContent(http.StatusBadRequest)
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

	// ctxs := ctx.Request().Context()
	quoteRequest, err := bid.services.GetBidsWithQuoteByQuoteId(context.Background(), quoteId)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	selectedBid := getBidFromBids(bidId, quoteRequest.Bids)
	quoteRequest.Bid = selectedBid
	// dont need to send bids
	quoteRequest.Bids = nil
	return ctx.JSON(http.StatusOK, quoteRequest)

}
func getBidFromBids(bidId string, bids []*v1.Bid) *v1.Bid {
	for _, j := range bids {
		if j.BidId == bidId {
			return j
		}
	}
	return nil
}
