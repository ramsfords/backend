package menuloom_backend

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	cloundflareCLI "github.com/ramsfords/backend/foundations/cloudflare"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/dynamo"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/foundations/zohomail"
	"github.com/ramsfords/backend/menuloom_backend/api"
	"github.com/ramsfords/backend/menuloom_backend/db"
	"github.com/ramsfords/backend/menuloom_backend/services"
	"github.com/ramsfords/backend/menuloom_backend/utils"
)

func MenuloomRunner(conf *configs.Config, s3 S3.S3Client, logger logger.Logger, dbRepo dynamo.DB, echoRouter *echo.Echo, app *pocketbase.PocketBase) {
	repo := db.New(conf, dbRepo)
	email := zohomail.New(conf.SitesSettings.Menuloom.Email)
	cloudinery := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)
	cloudFlareClient := cloundflareCLI.New(conf.SitesSettings.Menuloom.CloudFlareConfig)
	services := services.New(conf, s3, logger, email, repo, cloudinery, cloudFlareClient)
	api.SetUpAPi(echoRouter, services)
	// OR send a completely different email template
	app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(email, conf.SitesSettings.Menuloom.Email.FromName, conf.SitesSettings.Menuloom.Email.FromEmail))
	// OR send a completely different email template
	app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(email, conf.SitesSettings.Menuloom.Email.FromName, conf.SitesSettings.Menuloom.Email.FromEmail))
	// this sets auth token to cloudflare KV on every login
	app.OnRecordAuthRequest().Add(utils.AddTokenToCloudFlareKV(conf, logger, cloudFlareClient))

}
