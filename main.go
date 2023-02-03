package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/jsvm"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/adobe"
	"github.com/ramsfords/backend/foundations/dynamo"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/menuloom_backend"
)

func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}
	return filepath.Join(os.Args[0], "../pb_public")
}

func main() {
	conf := configs.GetConfig()
	s3 := S3.New(conf)
	dynamodDb := dynamo.New(conf)
	logger := logger.New("backend")
	app := pocketbase.New()
	adobe := adobe.NewAdobe(s3, &logger)
	var publicDirFlag string

	// add "--publicDir" option flag
	app.RootCmd.PersistentFlags().StringVar(
		&publicDirFlag,
		"publicDir",
		defaultPublicDir(),
		"the directory to serve static files",
	)
	migrationsDir := "" // default to "pb_migrations" (for js) and "migrations" (for go)

	// load js files to allow loading external JavaScript migrations
	jsvm.MustRegisterMigrations(app, &jsvm.MigrationsOptions{
		Dir: migrationsDir,
	})

	// register the `migrate` command
	migratecmd.MustRegister(app, app.RootCmd, &migratecmd.Options{
		TemplateLang: migratecmd.TemplateLangJS, // or migratecmd.TemplateLangGo (default)
		Dir:          migrationsDir,
		Automigrate:  true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"https://firstshipper.com", "https://www.firstshipper.com", "https://menuloom.com", "http://localhost:3000", "http://localhost:3001", "http://127.0.0.1:3000", "http://127.0.0.1:8787", "https://api.firstshipper.com"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "auth-guard", echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlRequestHeaders},
		}))
		e.Router.OPTIONS("/*", func(c echo.Context) error {
			return c.NoContent(http.StatusOK)
		})
		firstshipper_backend.FirstShipperRunner(conf, s3, logger, dynamodDb, e.Router, app, adobe)
		menuloom_backend.MenuloomRunner(conf, s3, logger, dynamodDb, e.Router, app)
		// serves static files from the provided public dir (if exists)
		e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/ping",
			Handler: func(c echo.Context) error {
				obj := map[string]interface{}{"message": "pong"}
				return c.JSON(http.StatusOK, obj)
			},
		})

		return nil

	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
