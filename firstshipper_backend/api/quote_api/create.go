package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/api/utils"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/rapid_utils"
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
	quotId := utils.GenerateString(10)
	quoteId := fmt.Sprint(quotId)
	qtReq.QuoteId = quoteId

	// make compatible rapid quote to send to rapid for quote rates
	rapidQuote, err := rapid_utils.MakeQuoteDetails(qtReq)
	if err != nil {
		return nil, errs.ErrInvalidInputs
	}

	// get quote from rapid
	res, err := qt.rapid.GetQuote(rapidQuote)
	if err != nil {
		qt.services.Logger.Error(err)
		return nil, errs.ErrInternal
	}

	saveQuote := rapid_utils.NewSaveQuoteStep2(rapidQuote, res)
	// err = qt.services.SaveRapidQuote(ctxx, *res, *qtReq)
	// if err != nil {
	// 	qt.services.Logger.Error(err)
	// 	return nil, errors.New("could not save quote in rapid")
	// }

	bidRes := rapid_utils.MakeBid(saveQuote, qtReq)
	if bidRes == nil {
		return nil, errs.ErrInternal
	}
	// save quote with rapid quote
	quoteRate := &model.QuoteRequest{
		QuoteRequest:   qtReq,
		RapidSaveQuote: saveQuote,
		Bids:           bidRes,
	}
	err = qt.services.SaveQuote(ctxx, quoteRate)
	if err != nil {
		return nil, errs.ErrStoreInternal
	}
	// validUntil := time.Now().Add(10 * time.Minute)
	// pickupDate, err := time.Parse(time.RFC3339, qtReq.ShipmentDetails.PickupDate)
	// if err != nil {
	// 	return nil, errs.ErrInternal
	// }
	// validUntilStr := ""

	// fmt.Println(pickupDate.Day())
	// if pickupDate.Day() == validUntil.Day() {
	// 	hour, _, _ := validUntil.Clock()
	// 	fmt.Println(qt.provider.Conf().LastPickupTime)
	// 	if hour < qt.provider.Conf().LastPickupTime {
	// 		validUntilStr = validUntil.Format(time.RFC3339)
	// 	} else {
	// 		validUntilStr = ""
	// 	}
	// } else {
	// 	validUntilStr = validUntil.Format(time.RFC3339)
	// }
	return &model.QuoteRequest{
		QuoteRequest: qtReq,
		Bids:         bidRes,
	}, nil
}
