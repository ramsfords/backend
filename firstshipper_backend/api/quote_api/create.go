package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v5"
	"github.com/pkg/errors"
	"github.com/ramsfords/backend/firstshipper_backend/api/utils"
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
	res, err := qt.GetNewQuote(ctx.Request().Context(), quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (qt Quote) GetNewQuote(ctxx context.Context, qtReq *v1.QuoteRequest) (*v1.QuoteRequest, error) {
	ctx, ok := ctxx.(*gin.Context)
	if !ok {
		return nil, errs.ErrInvalidInputs
	}
	err := utils.ValidateQuoteRequest(qtReq)
	if err != nil {
		return nil, err
	}
	quotId := utils.GenerateString(10)
	quoteId := fmt.Sprint(quotId)
	qtReq.QuoteId = quoteId
	qtReq.ShipmentDetails.QuoteId = quoteId
	for _, i := range qtReq.Commodities {
		i.QuoteId = quoteId

	}
	err = qt.services.SaveQuote(ctx, *qtReq)
	if err != nil {
		return nil, errs.ErrStoreInternal
	}
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
	err = qt.services.SaveRapidQuote(ctx, *res, *qtReq)
	if err != nil {
		qt.services.Logger.Error(err)
		return nil, errs.ErrInternal
	}
	saveQuote := rapid_utils.NewSaveQuoteStep2(rapidQuote, res)
	err = qt.services.SaveRapidQuote(ctx, *res, *qtReq)
	if err != nil {
		qt.services.Logger.Error(err)
		return nil, errors.New("could not save quote in rapid")
	}
	bidRes := rapid_utils.MakeBid(saveQuote, qtReq)
	if bidRes == nil {
		return nil, errs.ErrInternal
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
	return &v1.QuoteRequest{
		ShipmentDetails: qtReq.ShipmentDetails,
		Pickup:          qtReq.Pickup,
		Delivery:        qtReq.Delivery,
		Commodities:     qtReq.Commodities,
		Bids:            bidRes,
		QuoteId:         quoteId,
		Type:            "quote",
	}, nil
}
