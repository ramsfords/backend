package bids_api

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/services"
)

type Bids struct {
	services *services.Services
	rapid    *rapid.Rapid
}

func New(services *services.Services, echo *echo.Group, rapid *rapid.Rapid) {
	bids := Bids{
		services: services,
		rapid:    rapid,
	}
	protectedBidsGroup := echo.Group("/bids")
	protectedBidsGroup.GET("", bids.EchoGetQuoteWithBids)
	protectedBidsGroup.GET("/bid", bids.EchoGetQuoteWithBid)
	protectedBidsGroup.DELETE("", bids.EchoDeleteBids)
}
