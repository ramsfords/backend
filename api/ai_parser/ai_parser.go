package ai_parser

import (
	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
)

type AIParser struct {
	services *services.Services
}

func New(services *services.Services, echo *echo.Echo) {
	aiParser := AIParser{
		services: services,
	}
	protectedBolGroup := echo.Group("/aiparser", mid.Protected(services))
	protectedBolGroup.POST("/invoice", aiParser.ParseInvoice)

}
