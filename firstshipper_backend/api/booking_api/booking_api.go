package booking_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"
	"github.com/ramsfords/backend/foundations/adobe"
	"github.com/ramsfords/backend/menuloom_backend/api/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

type BookingApi struct {
	services *services.Services
	adobe    *adobe.Adobe
	rapid    *rapid.Rapid
}

func New(services *services.Services, echoClient *echo.Echo, rapid *rapid.Rapid) {
	book := BookingApi{
		services: services,
		rapid:    rapid,
	}
	protectedBolGroup := echoClient.Group("/booking")
	protectedBolGroup.GET("/:bookingId", book.EchoGetBooking)
	protectedBolGroup.POST("", book.EchoCreateBooking)
}

type Script struct {
	CustomScript string `json:"CustomScript,omitempty" dynamodbav:"CustomScript,omitempty"`
}
type BOlGenerateReq struct {
	Url             string `json:"url,omitempty" dynamodbav:"url,omityempty"`
	Margins         string `json:"margins,omitempty" dynamodbav:"margins,omityempty"`
	PaperSize       string `json:"paperSize,omitempty" dynamodbav:"paperSize,omityempty"`
	Orientation     string `json:"orientation,omitempty" dynamodbav:"orientation,omityempty"`
	PrintBackground bool   `json:"printBackground,omitempty" dynamodbav:"printBackground,omityempty"`
	Header          string `json:"header,omitempty" dynamodbav:"header,omityempty"`
	Footer          string `json:"footer,omitempty" dynamodbav:"footer,omityempty"`
	MediaType       string `json:"mediaType,omitempty" dynamodbav:"mediaType,omityempty"`
	Async           bool   `json:"async,omitempty" dynamodbav:"async,omityempty"`
	Encrypt         bool   `json:"encrypt,omitempty" dynamodbav:"encrypt,omityempty"`
	Profiles        string `json:"profiles,omitempty" dynamodbav:"profiles,omityempty"`
}
type BolGenerateResponse struct {
	Url              string `json:"url,omitempty" dynamodbav:"url,omitempty"`
	PageCount        int    `json:"pageCount,omitempty" dynamodbav:"pageCount,omitempty"`
	Error            bool   `json:"error,omitempty" dynamodbav:"error,omitempty"`
	Status           int    `json:"status,omitempty" dynamodbav:"status,omitempty"`
	Name             string `json:"name,omitempty" dynamodbav:"name,omitempty"`
	Credits          int32  `json:"credits,omitempty" dynamodbav:"credits,omitempty"`
	Duration         int32  `json:"duration,omitempty" dynamodbav:"duration,omitempty"`
	RemainingCredits int32  `json:"remainingCredits,omitempty" dynamodbav:"remainingCredits,omitempty"`
}

func (booking *BookingApi) EchoCreateBookingV0(ctx echo.Context) error {
	//booking id
	id := &v1.Id{}
	err := ctx.Bind(id)
	if err != nil {
		return errs.ErrInvalidInputs

	}
	bookingData, err := booking.services.GetBooking(ctx.Request().Context(), id.Id)
	if err != nil {
		return errs.ErrInvalidInputs
	}
	bookingData.BookingInfo.BolUrl = fmt.Sprintf("https://api.firstshipper.com/bol/v1/%s", bookingData.Bid.ShipmentId)
	bolGen := BOlGenerateReq{
		Url:             fmt.Sprintf("https://firstshipper.com/documents/booking_id=%s&token=ADMINuEvTLNWGJlJOKTENawuMQYBfboqOASkurvgmQUamNREeWdDTnh", bookingData.Bid.ShipmentId),
		Margins:         "5mm",
		PaperSize:       "A4",
		Orientation:     "Portrait",
		PrintBackground: true,
		MediaType:       "print",
		Async:           false,
	}
	jsonStr, err := json.Marshal(bolGen)
	if err != nil {
		return errs.ErrInvalidInputs
	}
	b := bytes.NewBuffer(jsonStr)
	pdfRenderReq, err := http.NewRequest("POST", booking.services.Conf.PdfRenderer.Url, b)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	pdfRenderReq.Header.Add("content-type", "application/json")
	pdfRenderReq.Header.Add("x-api-key", booking.services.Conf.PdfRenderer.ApiKey)
	pdfRenderRes, err := http.DefaultClient.Do(pdfRenderReq)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	resObj := &BolGenerateResponse{}
	err = json.NewDecoder(pdfRenderRes.Body).Decode(resObj)
	if err != nil {

		return errs.ErrInternal
	}
	// gets pdf from provided url
	reqs, err := http.Get(resObj.Url)
	if err != nil {
		return errs.ErrInternal
	}
	pdfBytes, err := io.ReadAll(reqs.Body)
	if err != nil {
		return errs.ErrInternal
	}
	pdfRenderRes.Body.Close()
	s3Input := &s3.PutObjectInput{
		Bucket:             aws.String("firstshipperbol"),
		Key:                aws.String("bol" + id.Id + ".pdf"),
		CacheControl:       aws.String(""),
		ContentType:        aws.String("application/pdf"),
		ContentDisposition: aws.String("inline"),
		Body:               strings.NewReader(string(pdfBytes)),
		Metadata: map[string]string{
			"businessId": bookingData.QuoteRequest.QuoteId,
		},
	}
	s3res, err := booking.services.S3Client.Client.PutObject(context.Background(), s3Input)
	if err != nil {
		return errs.ErrInternal
	}
	// user, err := bol.services.GetUser(ctx.Request().Context(), bookingData.Pickup.Contact.EmailAddress)
	// if err != nil {
	// 	return errs.ErrInternal
	// }
	// emailData := zohomail.EmailData{
	// 	ReceiverEmail: bookingData.Pickup.Contact.EmailAddress,
	// 	ReceiverName:  v1.Name,
	// 	SenderName:    "firstshipper bill of lading",
	// 	SenderEmail:   "quotes@firstshipper.com",
	// 	EmailSubject:  "please find BILL OF LADING for booking bol#" + bookingData.QuoteId,
	// 	RedirectLink:  bookingData.BolUrl,
	// }
	// err = bol.services.Email.Send(ctx.Request().Context(), bookingData, pdfBytes, emailData)
	// if err != nil {
	// 	bol.services.Logger.Debug("could not send bol in email")
	// }
	fmt.Println(s3res)
	return ctx.NoContent(http.StatusCreated)

}
