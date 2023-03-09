package booking_api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	rapid "github.com/ramsfords/backend/business/rapid/rapid_utils/book"
	books "github.com/ramsfords/backend/foundations/books"
	"github.com/ramsfords/backend/foundations/errs"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

type BookReq struct {
	QuoteRequest *v1.QuoteRequest `json:"quoteRequest"`
	Bid          *v1.Bid          `json:"bid"`
}

func (bookApi BookingApi) EchoCreateBooking(ctx echo.Context) error {
	quote := &BookReq{}
	err := ctx.Bind(quote)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	newCtx := ctx.Request().Context()
	ctx.Request().Header.Set("Cache-Control", "no-cache")
	res, err := bookApi.CreateNewBook(newCtx, quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, &res)
}

func (bookingApi BookingApi) CreateNewBook(ctxx context.Context, bkReq *BookReq) (*v1.BookingResponse, error) {
	business, err := bookingApi.services.Db.GetBusiness(ctxx, bkReq.QuoteRequest.BusinessId)
	if err != nil {
		return nil, fmt.Errorf("business not found")
	}
	oldQuote, err := bookingApi.services.Db.GetQuoteByQuoteId(ctxx, bkReq.QuoteRequest.QuoteId)
	if err != nil {
		return nil, fmt.Errorf("bid not found")
	}
	//updates oldQuote with new quoteRequest which is only updated value coming from frontend
	oldQuote.QuoteRequest = bkReq.QuoteRequest
	oldQuote.Business = business
	if !business.AllowBooking {
		return nil, errs.ErrNotAllowed
	}
	if business.BooksOpened {
		// business, err := bookingApi.services.Db.GetBusiness(ctxx, bkReq.QuoteRequest.BusinessId)
		// if err != nil {
		// 	return nil, errors.New("user not found")
		// }
		contact := books.NewContact(oldQuote.QuoteRequest.Business)
		contact, err = bookingApi.books.CreateContact(contact)
		if err != nil {
			logger.Error(err, "could not create zoho books for user")
		} else {
			business.BooksOpened = true
			business.BooksContactId = contact.ContactID
			business.ContactPersonsIds = GetContactPersonsIds(contact.ContactPersons)
			err = bookingApi.services.Db.UpdateBusiness(ctxx, business.BusinessId, *business)
			if err != nil {
				logger.Error(err, "could not update busiess with books contact id")
			}
		}
		fmt.Println(contact)
	}
	// books.InsertProspects(oldQuote.QuoteRequest.Delivery, *bookingApi.books)
	bid := getBidFormBids(oldQuote.Bids, bkReq.Bid.BidId)
	if bid.BidId == "" {
		return nil, fmt.Errorf("bid not found")
	}
	oldQuote.Bid = bid
	err = utils.ValidateBookRequest(oldQuote.QuoteRequest, bkReq.QuoteRequest, bid)
	if err != nil {
		return nil, err
	}

	// make saveQuoteStep3 Data
	err = rapid.SaveQuoteStep3(oldQuote, bid)
	if err != nil {
		// just log the error not Need to return error
		logger.Error(err, "could not save quote step 3")
	}
	saveQuoteRes, err := bookingApi.services.Rapid.SaveQuoteStep(oldQuote.RapidSaveQuote)
	if err != nil {
		// just log the error not Need to return error
		logger.Error(err, "could not get save quote response")
	}
	oldQuote.RapidSaveQuote.SavedQuoteID = saveQuoteRes.SavedQuoteID
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.SavedQuoteID = &saveQuoteRes.SavedQuoteID
	disPatchResponse, err := bookingApi.services.Rapid.Dispatch(oldQuote.RapidSaveQuote.ConfirmAndDispatch)
	if err != nil {
		// just log the error not Need to return error
		logger.Error(err, "could not get dispatch response")
		return nil, err
	}
	serviceType := fmt.Sprintf("%d", bid.ServiceType)
	shipmentId := fmt.Sprintf("%d", disPatchResponse.ShipmentID)
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.ShipmentID = &shipmentId
	oldQuote.RapidBooking = disPatchResponse
	bolNumber := "BOL" + oldQuote.QuoteRequest.QuoteId

	// hash the user password

	fileName := "BOL" + bid.BidId + "-" + strings.ToLower(utils.GenerateString(4)) + ".pdf"
	if err != nil {
		logger.Error(err, "Error in created hashed bol")
	}
	bolUrl := "https://firstshipperbol.s3.us-west-1.amazonaws.com/" + fileName
	carrier := utils.GetCarrierContact(disPatchResponse.CarrierName)
	oldQuote.BookingInfo = &v1.BookingInfo{
		ShipmentId:            int32(disPatchResponse.ShipmentID),
		FirstShipperBolNumber: bolNumber,
		FreightTerm:           1,
		CarrierName:           disPatchResponse.CarrierName,
		CarrierPhone:          carrier.Phone,
		CarrierEmail:          carrier.Email,
		CarrierProNumber:      disPatchResponse.CarrierProNumber,
		CarrierLogoUrl:        bid.VendorLogo,
		CarrierBolNumber:      disPatchResponse.CustomerBOLNumber,
		CarrierReference:      disPatchResponse.CarrierProNumber,
		PickupNumber:          disPatchResponse.CarrierProNumber,
		ServiceType:           serviceType,
		BolUrl:                bolUrl,
	}

	url := "https://bwipjs-api.metafloor.com/?bcid=code128&text={poNumber}"
	url = strings.ReplaceAll(url, "{poNumber}", oldQuote.BookingInfo.CarrierProNumber)
	oldQuote.BookingInfo.SvgData = url
	// go bookingApi.adobe.UrlToPdf(bid, fileName)
	err = bookingApi.services.Db.SaveBooking(ctxx, oldQuote)
	if err != nil {
		// just log the error not Need to return error
		logger.Error(err, "could not save booking")
		return nil, err
	}

	oldQuote.BookingInfo.SvgData = url
	oldQuote.BookingInfo.BolUrl = bolUrl
	outRes := &v1.BookingResponse{
		QuoteRequest: oldQuote.QuoteRequest,
		BookingInfo:  oldQuote.BookingInfo,
		Bid:          oldQuote.Bid,
		SvgData:      oldQuote.BookingInfo.SvgData,
	}
	res, err := bookingApi.services.AdobeCli.UrlToPdf(outRes, fileName)
	if err != nil {
		logger.Error(err, "could not save booking")
	}
	fmt.Println(res)
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

//	func makeBOlGenGetRequest(conf *configs.Config, bookingData *v1.BookingResponse) error {
//		url := ""
//		if conf.Env == "dev" {
//			url = "http://127.0.0.1:5173/bol" + bookingData.Bid.BidId
//		} else {
//			url = "" + "/bol" + bookingData.Bid.BidId
//		}
//		resp, err := http.Get(url)
//		if err != nil || resp.StatusCode != http.StatusOK {
//			return err
//		}
//		// marshall the response
//		defer resp.Body.Close()
//		res, err := ioutil.ReadAll(resp.Body)
//		if err != nil {
//			return err
//		}
//		fmt.Println(string(res))
//		return nil
//	}
func GetContactPersonsIds(contactPerson []books.ContactPerson) []string {
	ids := []string{}
	for _, j := range contactPerson {
		ids = append(ids, j.ContactPersonID)
	}
	return ids
}
