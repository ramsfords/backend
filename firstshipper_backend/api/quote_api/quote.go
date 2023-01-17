package quote_api

import (
	"context"

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
}

func New(services *services.Services, echo *echo.Group, rapid *rapid.Rapid) {
	qt := Quote{
		services: services,
		rapid:    rapid,
	}
	// quote api
	protectedQuoteGroup := echo.Group("quote")
	protectedQuoteGroup.DELETE(":id", qt.EchoDeleteQuote)
	protectedQuoteGroup.DELETE("", qt.EchoDeleteAllQuotes)
	protectedQuoteGroup.GET("", qt.EchoGetQuotes)
	protectedQuoteGroup.GET(":id", qt.EchoGetQuoteById)
	protectedQuoteGroup.POST("", qt.EchoCreateQuote)
	protectedQuoteGroup.PATCH("", qt.EchoUpdateQuote)
}
