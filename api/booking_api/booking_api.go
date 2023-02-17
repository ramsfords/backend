package booking_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/books"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
)

type BookingApi struct {
	services *services.Services
	books    *books.API
}

func New(services *services.Services, echoClient *echo.Group) {
	book := BookingApi{
		services: services,
	}
	book.books = books.New(services.Zoho, services.Conf)
	protectedBolGroup := echoClient.Group("/booking", mid.Protected(services))
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
