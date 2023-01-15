package mid

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/configs"
)

func ReverseProxy(conf *configs.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			remote, _ := url.Parse("http://localhost:3000")
			proxy := httputil.NewSingleHostReverseProxy(remote)
			proxy.Director = func(req *http.Request) {
				req.Header = ctx.Request().Header
				req.Host = remote.Host
				req.URL = ctx.Request().URL
				req.URL.Scheme = remote.Scheme
				req.URL.Host = remote.Host
			}
			proxy.ServeHTTP(ctx.Response().Writer, ctx.Request())
			return next(ctx)
		}

	}
}
