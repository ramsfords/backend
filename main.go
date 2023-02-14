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
	"github.com/ramsfords/backend/api"
	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/services"
	"github.com/ramsfords/backend/utils"
)

func Runner(services *services.Services, echoRouter *echo.Echo, app *pocketbase.PocketBase) {

	go func() error {
		err := services.Rapid.Login(&models.AuthRequestPayload{
			Username: services.Conf.SitesSettings.FirstShipper.RapidShipLTL.UserName,
			Password: services.Conf.SitesSettings.FirstShipper.RapidShipLTL.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}()

	api.SetUpAPi(echoRouter, services)
	// OR send a completely different email template
	app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(services))
	// OR send a completely different email template
	app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(services.Conf))
	// this sets auth token to cloudflare KV on every login

}

func defaultPublicDir() string {
	if strings.HasPrefix(os.Args[0], os.TempDir()) {
		// most likely ran with go run
		return "./pb_public"
	}
	return filepath.Join(os.Args[0], "../pb_public")
}

func main() {
	conf := configs.GetConfig()
	// s3 := S3.New(conf)
	// dynamodDb := dynamo.New(conf)
	// logger := logger.New("backend")
	servicesInstance := services.New(conf)
	app := pocketbase.New()
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
			c.Request().Header.Add("Access-Control-Allow-Origin", "*")
			c.Request().Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Request().Header.Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			return c.NoContent(http.StatusOK)
		})

		Runner(servicesInstance, e.Router, app)
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
