package firstshipper_backend

import (
	"log"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend/api"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid"
	"github.com/ramsfords/backend/firstshipper_backend/business/rapid/models"
	"github.com/ramsfords/backend/firstshipper_backend/db"
	"github.com/ramsfords/backend/firstshipper_backend/services"
	"github.com/ramsfords/backend/firstshipper_backend/utils"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/adobe"
	cloundflareCLI "github.com/ramsfords/backend/foundations/cloudflare"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/dynamo"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/foundations/zohomail"
)

func FirstShipperRunner(conf *configs.Config, s3 S3.S3Client, logger logger.Logger, dbClient dynamo.DB, echoRouter *echo.Echo, app *pocketbase.PocketBase, adob *adobe.Adobe) {
	repo := db.New(dbClient, conf)
	email := zohomail.New(conf.SitesSettings.FirstShipper.Email)
	rapid := rapid.New()
	cloudinery := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)

	cloudFlareClient := cloundflareCLI.New(conf.SitesSettings.Menuloom.CloudFlareConfig)
	services := services.New(conf, s3, logger, email, repo, cloudinery, cloudFlareClient)
	go func() error {
		err := rapid.Login(&models.AuthRequestPayload{
			Username: conf.SitesSettings.FirstShipper.RapidShipLTL.UserName,
			Password: conf.SitesSettings.FirstShipper.RapidShipLTL.Password,
		})
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}()

	api.SetUpAPi(echoRouter, services, rapid, adob)
	// OR send a completely different email template
	app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(email, conf, services))
	// OR send a completely different email template
	app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(email, conf.SitesSettings.Menuloom.Email.FromName, conf.SitesSettings.Menuloom.Email.FromEmail))
	// this sets auth token to cloudflare KV on every login
	app.OnRecordAuthRequest().Add(utils.AddTokenToCloudFlareKV(conf, logger, cloudFlareClient))

}
