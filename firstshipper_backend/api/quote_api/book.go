package quote_api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/api/utils"
	"github.com/ramsfords/backend/firstshipper_backend/business/core/model"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/rapid_utils/book"
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
	return ctx.JSON(http.StatusCreated, &res)
}

func (qt Quote) CreateNewBook(ctxx context.Context, bkReq *v1.BookRequest) (*model.QuoteRequest, error) {
	oldQuote, err := qt.services.GetQuoteByQuoteId(ctxx, bkReq.QuoteRequest.QuoteId, bkReq.QuoteRequest.BusinessId)
	if err != nil {
		return nil, fmt.Errorf("bid not found")
	}
	bid := getBidFormBids(oldQuote.Bids, bkReq.BidId)
	if bid.BidId == "" {
		return nil, fmt.Errorf("bid not found")
	}
	oldQuote.Bid = bid
	err = utils.ValidateBookRequest(oldQuote.QuoteRequest, bkReq.QuoteRequest, bid)
	if err != nil {
		return nil, err
	}
	//updates oldQuote with new quoteRequest which is only updated value coming from frontend
	oldQuote.QuoteRequest = bkReq.QuoteRequest
	// make saveQuoteStep3 Data
	err = book.SaveQuoteStep3(oldQuote)
	if err != nil {
		// just log the error not Need to return error
		qt.services.Logger.Error(err)
	}
	saveQuoteRes, err := qt.rapid.SaveQuoteStep(oldQuote.RapidSaveQuote)
	if err != nil {
		// just log the error not Need to return error
		qt.services.Logger.Error(err)
	}
	oldQuote.RapidSaveQuote.SavedQuoteID = saveQuoteRes.SavedQuoteID
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.SavedQuoteID = &saveQuoteRes.SavedQuoteID
	res, err := qt.rapid.Dispatch(oldQuote.RapidSaveQuote.ConfirmAndDispatch)
	if err != nil {
		// just log the error not Need to return error
		qt.services.Logger.Error(err)
		return nil, err
	}
	shipmentId := fmt.Sprintf("%d", res.ShipmentID)
	serviceType := fmt.Sprintf("%d", bid.ServiceType)
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.ShipmentID = &shipmentId
	oldQuote.RapidBooking = res

	oldQuote.BookingInfo = &v1.BookingInfo{
		ShipmentId:            int32(res.ShipmentID),
		FirstShipperBolNumber: "BOL" + shipmentId,
		FreightTerm:           1,
		CarrierName:           res.CarrierName,
		CarrierPhone:          res.CarrierPhone,
		CarrierEmail:          "",
		CarrierProNumber:      res.CarrierPRONumber,
		CarrierLogoUrl:        bid.VendorLogo,
		CarrierBolNumber:      res.CustomerBOLNumber,
		CarrierReference:      res.CarrierPRONumber,
		PickupNumber:          res.CarrierPRONumber,
		ServiceType:           serviceType,
		BolUrl:                "https://firstshipper.com/admin/bol/" + shipmentId,
	}
	err = qt.services.SaveBooking(ctxx, oldQuote)
	if err != nil {
		// just log the error not Need to return error
		qt.services.Logger.Error(err)
		return nil, err
	}
	outRes := &model.QuoteRequest{
		QuoteRequest: oldQuote.QuoteRequest,
		BookingInfo:  oldQuote.BookingInfo,
	}
	return outRes, nil
}
func getBidFormBids(bids []*v1.Bid, bidId string) *v1.Bid {
	for _, bid := range bids {
		if bid.BidId == bidId {
			return bid
		}
	}
	return nil
}
