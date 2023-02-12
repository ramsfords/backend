package booking_api

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/utils"
	rapid "github.com/ramsfords/backend/business/rapid/rapid_utils/book"
	"github.com/ramsfords/backend/configs"
	template "github.com/ramsfords/backend/email"
	books "github.com/ramsfords/backend/foundations/books"
	errs "github.com/ramsfords/backend/foundations/error"
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
	newCtx := ctx.Request().Context()

	ctx.Request().Header.Set("Cache-Control", "no-cache")

	res, err := bookApi.CreateNewBook(newCtx, quote)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusCreated, &res)
}

func (bookingApi BookingApi) CreateNewBook(ctxx context.Context, bkReq *v1.BookRequest) (*v1.BookingResponse, error) {
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
	if !business.BooksOpened {
		// business, err := bookingApi.services.Db.GetBusiness(ctxx, bkReq.QuoteRequest.BusinessId)
		// if err != nil {
		// 	return nil, errors.New("user not found")
		// }
		contact := books.NewContact(oldQuote.QuoteRequest.Business)
		contact, err = bookingApi.books.CreateContact(contact)
		if err != nil {
			bookingApi.services.Logger.Errorf("could not create zoho books for user: ", oldQuote.Business.AdminUser)
		} else {
			business.BooksOpened = true
			business.BooksContactId = contact.ContactID
			business.ContactPersonsIds = GetContactPersonsIds(contact.ContactPersons)
			err = bookingApi.services.Db.UpdateBusiness(ctxx, business.BusinessId, *business)
			if err != nil {
				bookingApi.services.Logger.Errorf("could not update busiess with books contact id", oldQuote.Business)
			}
		}
		fmt.Println(contact)
	}
	books.InsertProspects(oldQuote.QuoteRequest.Delivery, *bookingApi.books)
	bid := getBidFormBids(oldQuote.Bids, bkReq.BidId)
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
		bookingApi.services.Logger.Error(err)
	}
	saveQuoteRes, err := bookingApi.services.Rapid.SaveQuoteStep(oldQuote.RapidSaveQuote)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
	}
	oldQuote.RapidSaveQuote.SavedQuoteID = saveQuoteRes.SavedQuoteID
	oldQuote.RapidSaveQuote.ConfirmAndDispatch.SavedQuoteID = &saveQuoteRes.SavedQuoteID
	disPatchResponse, err := bookingApi.services.Rapid.Dispatch(oldQuote.RapidSaveQuote.ConfirmAndDispatch)
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

	// hash the user password

	fileName := bid.BidId + "-" + utils.GenerateString(4) + ".pdf"
	if err != nil {
		bookingApi.services.Logger.Errorf("Error in created hashed bol %v", err)
	}
	bolUrl := "https://firstshipperbol.s3.us-west-1.amazonaws.com/" + fileName

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
	url := "https://bwipjs-api.metafloor.com/?bcid=code128&text={poNumber}"
	url = strings.ReplaceAll(url, "{poNumber}", oldQuote.BookingInfo.CarrierProNumber)
	oldQuote.BookingInfo.SvgData = url
	// go bookingApi.adobe.UrlToPdf(bid, fileName)
	err = bookingApi.services.Db.SaveBooking(ctxx, oldQuote)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
		return nil, err
	}
	err = makeBOlGenGetRequest(bookingApi.services.Conf, fileName)
	if err != nil {
		// just log the error not Need to return error
		bookingApi.services.Logger.Error(err)
	}
	oldQuote.BookingInfo.SvgData = url
	oldQuote.BookingInfo.BolUrl = bolUrl
	outRes := &v1.BookingResponse{
		QuoteRequest: oldQuote.QuoteRequest,
		BookingInfo:  oldQuote.BookingInfo,
		SvgData:      oldQuote.BookingInfo.SvgData,
	}
	emailSubject := "FirstShipper: Booking Confirmation " + "Pickup Number: " + oldQuote.BookingInfo.CarrierProNumber + " " + "BOL Number: " + oldQuote.BookingInfo.FirstShipperBolNumber
	data := template.Data{
		To:          []string{oldQuote.QuoteRequest.Pickup.Contact.EmailAddress},
		Subject:     emailSubject,
		From:        "quotes@firstshipper.com",
		ContentType: "text/html",
		Body:        "Please find your BOL attached",
		Attachments: []template.Attachment{
			{
				Path: "firstshipperbol/" + fileName,
				Type: template.AttachmentTypeS3,
			},
		},
	}
	go func() {
		fmt.Println("data to send emaill is", data)
		time.Sleep(5 * time.Second)
		bolSentRes, err := template.Send(context.Background(), data)
		if err != nil {
			bookingApi.services.Logger.Errorf("error sending bol to user ", err, outRes)
		}
		bookingApi.services.Logger.Infof("bol sent", bolSentRes, outRes)
	}()
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
func makeBOlGenGetRequest(conf *configs.Config, fileName string) error {
	url := "https://firstshipper.com/api/bol?fileName=" + fileName
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	// marshall the response
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
func GetContactPersonsIds(contactPerson []books.ContactPerson) []string {
	ids := []string{}
	for _, j := range contactPerson {
		ids = append(ids, j.ContactPersonID)
	}
	return ids
}
