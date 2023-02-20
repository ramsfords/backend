package main

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/ramsfords/backend/api"
	"github.com/ramsfords/backend/business/rapid/models"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/mid"
	"github.com/ramsfords/backend/services"
	"golang.org/x/time/rate"
)

func Runner(services *services.Services, echoRouter *echo.Echo) {
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

	// OR send a completely different email template
	// app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(services))
	// // OR send a completely different email template
	// app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(services.Conf))
	// this sets auth token to cloudflare KV on every login

}

func main() {
	conf := configs.GetConfig()
	servicesInstance := services.New(conf)
	echos := echo.New()
	echos.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(200, echo.Map{
			"message": "pong",
			"status":  "ok",
			"code":    200,
		})
	})
	echos.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	echos.Use(middleware.Logger())
	var reqLimit rate.Limit = 20
	if conf.Env == "dev" {
		reqLimit = 1000
	}
	echos.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(reqLimit)))
	echos.Use(middleware.Recover())
	echos.Use(mid.CORS())
	api.SetUpAPi(echos, servicesInstance)
	go func() error {
		err := servicesInstance.Rapid.Login(&models.AuthRequestPayload{
			Username: servicesInstance.Conf.SitesSettings.FirstShipper.RapidShipLTL.UserName,
			Password: servicesInstance.Conf.SitesSettings.FirstShipper.RapidShipLTL.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}()
	if err := echos.Start(":8090"); err != nil {
		log.Fatal(err)
	}
}
