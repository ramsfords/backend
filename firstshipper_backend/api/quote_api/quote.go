package quote_api

import (
	"context"
	"sync"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

type QuoteContract interface {
	DeleteQuote(context.Context, *v1.Id) (*v1.Ok, error)
	DeleteQuotes(context.Context, *v1.Ids) (*v1.Ok, error)
	GetAllQuotes(ctx context.Context, qtReq *v1.Empty) (*v1.QuotesResponse, error)
	GetNewQuotes(ctx context.Context, qtReq []*v1.QuoteRequest) (*v1.QuotesResponse, error)
	UpdateQuote(context.Context, *v1.QuoteRequest) (*v1.QuoteRequest, error)
}
type Quote struct {
	services *services.Services
	rapid    *rapid.Rapid
	*sync.Mutex
}

func New(services *services.Services, echo *echo.Echo, rapid *rapid.Rapid) {
	qt := Quote{
		services: services,
		rapid:    rapid,
		Mutex:    &sync.Mutex{},
	}
	// quote api
	protectedQuoteGroup := echo.Group("/quote")

	//GET
	protectedQuoteGroup.GET("/quotes", qt.EchoGetAllQuotes)
	protectedQuoteGroup.GET("/:quoteId", qt.EchoGetQuoteByQuoteId)
	protectedQuoteGroup.GET("/bids/:quoteId", qt.EchoGetBidsByQuoteId)
	protectedQuoteGroup.GET("/quotewithbid/:bidId", qt.EchoGetQuoteWithBidByBidId)
	protectedQuoteGroup.GET("/quotewithbids/:quoteId", qt.EchoGetQuoteWithBidsByQuoteId)
	protectedQuoteGroup.GET("/business/:businessId", qt.EchoGetQuotesByBusinessId)
	//POST
	protectedQuoteGroup.POST("", qt.EchoCreateQuote)
	//DELETE
	protectedQuoteGroup.DELETE("/:quoteId", qt.EchoDeleteQuote)
	protectedQuoteGroup.DELETE("", qt.EchoDeleteAllQuotes)
	//PATCH
	protectedQuoteGroup.PATCH("", qt.EchoUpdateQuote)

}
