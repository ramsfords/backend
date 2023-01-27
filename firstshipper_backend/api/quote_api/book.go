package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/api/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoCreateBook(ctx echo.Context) error {
	quote := &v1.BookRequest{}
	err := ctx.Bind(quote)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	ctx.Request().Header.Set("Cache-Control", "no-cache")
	newCtx := ctx.Request().Context()
	res, err := qt.CreateNewBook(newCtx, quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, map[string]string{"message": *res})
}

func (qt Quote) CreateNewBook(ctxx context.Context, bkReq *v1.BookRequest) (*string, error) {
	qtReq, err := qt.services.GetQuoteByQuoteId(ctxx, bkReq.QuoteRequest.QuoteId, bkReq.QuoteRequest.BusinessId)
	if err != nil {
		return nil, fmt.Errorf("bid not found")
	}
	bid := getBidFormBids(qtReq.Bids, bkReq.BidId)
	if bid.BidId == "" {
		return nil, fmt.Errorf("bid not found")
	}
	err = utils.ValidateBookRequest(qtReq.QuoteRequest, bkReq.QuoteRequest, qtReq.Bids[0])
	if err != nil {
		return nil, err
	}
	quotId := utils.GenerateString(10)
	quoteId := fmt.Sprint(quotId)

	fmt.Println(quoteId)

	return aws.String("booking was succesfull"), nil
}
func getBidFormBids(bids []v1.Bid, bidId string) v1.Bid {
	for _, bid := range bids {
		if bid.BidId == bidId {
			return bid
		}
	}
	return v1.Bid{}
}
