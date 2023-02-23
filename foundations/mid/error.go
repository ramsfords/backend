package mid

import (
	"log"

	"github.com/labstack/echo/v5"
)

func ErroHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("executing middleware for route named: %s", c.RouteInfo().Name())
		return next(c)
	}
}
