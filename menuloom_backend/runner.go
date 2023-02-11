package menuloom_backend

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/ramsfords/backend/menuloom_backend/api"
	"github.com/ramsfords/backend/menuloom_backend/services"
	"github.com/ramsfords/backend/menuloom_backend/utils"
)

func MenuloomRunner(services *services.Services, echoRouter *echo.Echo, app *pocketbase.PocketBase) {
	api.SetUpAPi(echoRouter, services)
	// OR send a completely different email template
	app.OnMailerBeforeRecordVerificationSend().Add(utils.SendConfrimEmailEventHandler(services, services.Conf.SitesSettings.Menuloom.Email.FromName, services.Conf.SitesSettings.Menuloom.Email.FromEmail))
	// OR send a completely different email template
	app.OnMailerBeforeRecordResetPasswordSend().Add(utils.SendResetPasswordLinkEventHandler(services.Email, services.Conf.SitesSettings.Menuloom.Email.FromName, services.Conf.SitesSettings.Menuloom.Email.FromEmail))

}
