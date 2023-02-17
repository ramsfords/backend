package mid

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "https://localhost:3000", "https://firstshipper.com", "https://www.firstshipper.com", "https://localhost:3001", "https://127.0.0.1:3000", "https://127.0.0.1:8787", "https://api.firstshipper.com"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"firstAuth",
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAccessControlRequestHeaders,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderAccessControlAllowCredentials,
		},
		AllowMethods: []string{
			http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodHead, http.MethodPatch},
		AllowCredentials: true,
	})
}

func getValidOrigin(origin string) bool {
	origins := []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://firstshipper.com", "https://localhost:3000", "https://127.0.0.1:3000", "https://firstshipper.com"}
	for _, v := range origins {
		if v == origin {
			return true
		}
	}
	return false
}
