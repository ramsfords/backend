package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/utils"
	"github.com/ramsfords/backend/business/core/model"
	rapid "github.com/ramsfords/backend/business/rapid/rapid_utils/quote"
	errs "github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (qt Quote) EchoCreateQuote(ctx echo.Context) error {
	quote := &v1.QuoteRequest{}
	err := ctx.Bind(quote)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	ctx.Request().Header.Set("Cache-Control", "no-cache")
	newCtx := ctx.Request().Context()
	res, err := qt.GetNewQuote(newCtx, quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (qt Quote) GetNewQuote(ctxx context.Context, qtReq *v1.QuoteRequest) (*model.QuoteRequest, error) {
	err := utils.ValidateQuoteRequest(qtReq)
	if err != nil {
		return nil, err
	}
	qtReq.QuoteId = fmt.Sprint(qt.services.Db.GetQuoteCount())

	// make compatible rapid quote to send to rapid for quote rates
	rapidQuoteRequest, err := rapid.MakeQuoteDetails(qtReq)
	if err != nil {
		return nil, errs.ErrInvalidInputs
	}

	// get quote from rapid
	res, err := qt.services.Rapid.GetQuote(rapidQuoteRequest)
	if err != nil {
		qt.services.Logger.Error(err)
		return nil, errs.ErrInternal
	}

	saveQuote := rapid.SaveQuoteStep2(rapidQuoteRequest, res)
	saveQuoteRes, err := qt.services.Rapid.SaveQuoteStep(saveQuote)
	if err != nil {
		// just log the error not Need to return error
		qt.services.Logger.Error(err)
	}

	bidRes := rapid.MakeBid(qtReq, res.DayDeliveries, qt.Mutex)
	if bidRes == nil {
		return nil, errs.ErrInternal
	}
	// save quote with rapid quote
	quoteRate := &model.QuoteRequest{
		QuoteRequest:      qtReq,
		RapidSaveQuote:    saveQuote,
		Bids:              bidRes,
		SaveQuoteResponse: saveQuoteRes,
		RapidBooking:      nil,
		Bid:               nil,
	}
	err = qt.services.Db.SaveQuote(ctxx, quoteRate)
	go qt.services.Db.IncreateQuoteCount()
	if err != nil {
		return nil, errs.ErrStoreInternal
	}
	return &model.QuoteRequest{
		QuoteRequest: qtReq,
		Bids:         bidRes,
	}, nil
}
