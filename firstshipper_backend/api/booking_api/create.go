package booking_api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/api/utils"
	rapid "github.com/ramsfords/backend/firstshipper_backend/business/rapid/rapid_utils/book"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (bookApi BookingApi) EchoCreateBooking(ctx echo.Context) error {
	quote := &v1.BookRequest{}
	err := ctx.Bind(quote)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// quote.QuoteRequest = &v1.QuoteRequest{
	// 	QuoteId:    "23122",
	// 	BusinessId: "kandelsuren@gmail.com",
	// }
	ctx.Request().Header.Set("Cache-Control", "no-cache")
	newCtx := ctx.Request().Context()
	res, err := bookApi.CreateNewBook(newCtx, quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, &res)
}

func (bookingApi BookingApi) CreateNewBook(ctxx context.Context, bkReq *v1.BookRequest) (*v1.BookingResponse, error) {
	oldQuote, err := bookingApi.services.GetQuoteByQuoteId(ctxx, bkReq.QuoteRequest.QuoteId)
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
	err = rapid.SaveQuoteStep3(oldQuote, bid)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
	}
	saveQuoteRes, err := bookingApi.rapid.SaveQuoteStep(oldQuote.RapidSaveQuote)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
	}
	oldQuote.RapidSaveQuote.SavedQuoteID = saveQuoteRes.SavedQuoteID
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.SavedQuoteID = &saveQuoteRes.SavedQuoteID
	disPatchResponse, err := bookingApi.rapid.Dispatch(oldQuote.RapidSaveQuote.ConfirmAndDispatch)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
		return nil, err
	}
	serviceType := fmt.Sprintf("%d", bid.ServiceType)
	shipmentId := fmt.Sprintf("%d", disPatchResponse.ShipmentID)
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.ShipmentID = &shipmentId
	oldQuote.RapidBooking = disPatchResponse
	bolNumber := "BOL" + oldQuote.QuoteRequest.QuoteId
	bolUrl := "https://firstshipperbol.s3.us-west-1.amazonaws.com/" + bolNumber + ".pdf"
	url := "https://bwipjs-api.metafloor.com/?bcid=code128&text={poNumber}"
	url = strings.ReplaceAll(url, "{poNumber}", oldQuote.BookingInfo.CarrierProNumber)
	oldQuote.BookingInfo = &v1.BookingInfo{
		ShipmentId:            int32(disPatchResponse.ShipmentID),
		FirstShipperBolNumber: bolNumber,
		FreightTerm:           1,
		CarrierName:           disPatchResponse.CarrierName,
		CarrierPhone:          disPatchResponse.CarrierPhone,
		CarrierEmail:          "",
		CarrierProNumber:      disPatchResponse.CarrierProNumber,
		CarrierLogoUrl:        bid.VendorLogo,
		CarrierBolNumber:      disPatchResponse.CustomerBOLNumber,
		CarrierReference:      disPatchResponse.CarrierProNumber,
		PickupNumber:          disPatchResponse.CarrierProNumber,
		ServiceType:           serviceType,
		BolUrl:                bolUrl,
	}

	oldQuote.BookingInfo.SvgData = url
	go bookingApi.adobe.UrlToPdf(bid.BidId, oldQuote.QuoteRequest.BusinessId)
	err = bookingApi.services.SaveBooking(ctxx, oldQuote)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
		return nil, err
	}
	oldQuote.BookingInfo.SvgData = url
	oldQuote.BookingInfo.BolUrl = bolUrl
	outRes := &v1.BookingResponse{
		QuoteRequest: oldQuote.QuoteRequest,
		BookingInfo:  oldQuote.BookingInfo,
		SvgData:      oldQuote.BookingInfo.SvgData,
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
