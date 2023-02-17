package quote_api

import (
	"context"
	"sync"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

type QuoteContract interface {
	GetAllQuotes(ctx context.Context, qtReq *v1.Empty) (*v1.QuotesResponse, error)
	GetNewQuotes(ctx context.Context, qtReq []*v1.QuoteRequest) (*v1.QuotesResponse, error)
	UpdateQuote(context.Context, *v1.QuoteRequest) (*v1.QuoteRequest, error)
}
type Quote struct {
	services *services.Services
	*sync.Mutex
}

func New(services *services.Services, echo *echo.Group) {
	qt := Quote{
		services: services,
		Mutex:    &sync.Mutex{},
	}
	// quote api
	protectedQuoteGroup := echo.Group("/quote", mid.Protected(services))

	//GET
	protectedQuoteGroup.GET("/quotes", qt.EchoGetAllQuotes)
	protectedQuoteGroup.GET("/:quoteId", qt.EchoGetQuoteByQuoteId)
	protectedQuoteGroup.GET("/bids/:quoteId", qt.EchoGetBidsByQuoteId)
	protectedQuoteGroup.GET("/quotewithbid/:bidId", qt.EchoGetQuoteWithBidByBidId)
	protectedQuoteGroup.GET("/quotewithbids/:quoteId", qt.EchoGetQuoteWithBidsByQuoteId)
	protectedQuoteGroup.GET("/business/:businessId", qt.EchoGetQuotesByBusinessId)
	//POST
	protectedQuoteGroup.POST("", qt.EchoCreateQuote)
	//PATCH
	protectedQuoteGroup.PATCH("", qt.EchoUpdateQuote)

}
