package mid

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

// Middleware creates Echo middleware with provided config that
// instruments requests.
//
//	e := echo.New()
//	// Add the nrecho middleware before other middlewares or routes:
//	e.Use(nrecho.MiddlewareWithConfig(nrecho.Config{App: app}))
func Tracer(app *newrelic.Application) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			rw := c.Response().Writer
			txn := app.StartTransaction(c.RouteInfo().Path())
			defer txn.End()
			txn.SetWebRequestHTTP(c.Request())
			c.Response().Writer = txn.SetWebResponse(rw)
			// Add txn to c.Request().Context()
			c.SetRequest(c.Request().WithContext(newrelic.NewContext(c.Request().Context(), txn)))
			app.RecordLog(newrelic.LogData{
				Message:   "Hello World",
				Timestamp: time.Now().Unix(),
				Severity:  "info",
			})
			err = next(c)

			// Record the response code. The response headers are not captured
			// in this case because they are set after this middleware returns.
			// Designed to mimic the logic in echo.DefaultHTTPErrorHandler.
			if nil != err && !c.Response().Committed {

				c.Response().Writer = rw

				if httperr, ok := err.(*echo.HTTPError); ok {
					txn.SetWebResponse(nil).WriteHeader(httperr.Code)
				} else {
					txn.SetWebResponse(nil).WriteHeader(http.StatusInternalServerError)
				}
			}

			return
		}
	}
}
