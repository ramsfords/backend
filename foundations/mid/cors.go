package mid

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/configs"
)

func AddHeaders(ctx echo.Context, url string) echo.Context {
	ctx.Request().Header.Set("Content-Type", "*")
	ctx.Request().Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request().Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max, Access-Control-Allow-Origin,Access-Control-Allow-Methods,custom_header, pragma")
	ctx.Request().Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Request().Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
	return ctx
}

func CORS(conf *configs.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			//fontEndUrl := conf.GetFontEndURL()
			origin := ctx.Request().Header.Get("origin")
			valid := strings.Contains(origin, "firstshipper.com") || strings.Contains(origin, "127.0.0.1") || origin == ""
			if !valid {
				origin = ""
			}
			fmt.Println("origin was", origin)
			ctx.Request().Header.Set("Access-Control-Allow-Origin", origin)
			ctx.Request().Header.Set("Access-Control-Allow-Credentials", "true")
			ctx.Request().Header.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, access-control-allow-methods, access-control-allow-origin, first-shipper-token")
			ctx.Request().Header.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

			if ctx.Request().Method == "OPTIONS" {
				return ctx.NoContent(204)
			}
			return nil

		}
	}
}
