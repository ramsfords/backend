package firstshipper_backend

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/ramsfords/backend/shipper/api"
	"github.com/ramsfords/backend/shipper/business/rapid/models"
	"github.com/ramsfords/backend/shipper/services"
	"github.com/ramsfords/backend/shipper/utils"
)

func FirstShipperRunner(services *services.Services, echoRouter *echo.Echo, app *pocketbase.PocketBase) {

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
