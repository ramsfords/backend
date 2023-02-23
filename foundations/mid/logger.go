package mid

import (
	"time"

	"github.com/labstack/echo/v5"
	"github.com/sirupsen/logrus"
)

// Skipper defines a function to skip middleware.
type Skipper func(c echo.Context) bool

// Config defines the config for Logger middleware.
type Config struct {
	Logger  *logrus.Logger
	Skipper Skipper
	Fields  []string
	Status  int
}

// Middleware returns a Logger middleware with config.
func Middleware(config Config, logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			if config.Skipper(ctx) {
				return next(ctx)
			}
			fields := logrus.Fields{}

			req := ctx.Request()
			res := ctx.Response()
			start := time.Now()
			if err = next(ctx); err != nil {
				echo.ErrUnauthorized.WithInternal(err)
				fields["error"] = err
			}
			stop := time.Now()

			path := req.URL.Path
			if path == "" {
				path = "/"
			}

			for _, field := range config.Fields {
				switch field {
				case "id":
					id := req.Header.Get(echo.HeaderXRequestID)
					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}
					fields[field] = id
				case "ip":
					fields[field] = ctx.RealIP()
				case "host":
					fields[field] = req.Host
				case "referer":
					fields[field] = req.Referer()
				case "user_agent":
					fields[field] = req.UserAgent()
				case "status":
					fields[field] = res.Status
				case "latency":
					fields[field] = stop.Sub(start).String()
				case "headers":
					fields[field] = req.Header
				}
			}

			if res.Status < config.Status {
				return
			}
			switch {
			case res.Status >= 500:
				config.Logger.WithFields(fields).Errorf("%s %s", req.Method, path)
			case res.Status >= 400:
				config.Logger.WithFields(fields).Warnf("%s %s", req.Method, path)
			default:
				config.Logger.WithFields(fields).Debugf("%s %s", req.Method, path)
			}
			return
		}
	}
}
