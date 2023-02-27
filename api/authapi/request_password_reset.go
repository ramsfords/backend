package authapi

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/utils"
)

func (auth AuthApi) EchoRequestResetPassword(ctx echo.Context) error {
	emailAddress := ctx.PathParam("email")
	fmt.Println("emailAddress", emailAddress)
	if emailAddress == "" || !utils.IsEmailValid(emailAddress) {
		return ctx.NoContent(http.StatusBadRequest)
	}
	usr, err := auth.services.Db.Getuser(ctx.Request().Context(), emailAddress)
	if err != nil {
		if err.Error() == "no data found" {
			return ctx.NoContent(http.StatusNotFound)
		}
		return ctx.NoContent(http.StatusInternalServerError)
	}
	if usr == nil {
		return ctx.NoContent(http.StatusNotFound)
	}
	err = auth.services.SupaClient.Auth.ResetPasswordForEmail(ctx.Request().Context(), emailAddress)
	// encodedStr, err := auth.services.Crypto.Encrypt(emailAddress)
	// if err != nil {
	// 	auth.services.Logger.Error("Error encrypting email", map[string]interface{}{"err": err})
	// 	return ctx.NoContent(500)
	// }

	// redirectLink := auth.services.Conf.GetFirstShipperFontEndURL() + "/reset-password?token=" + encodedStr + "&email=" + emailAddress
	// emailData := email.Data{
	// 	To:          []string{emailAddress},
	// 	Subject:     "Your password reset link",
	// 	From:        "noreply@firstshipper.com",
	// 	Body:        email.GetResetPasswordTemplate("", redirectLink),
	// 	ContentType: "text/html",
	// }
	// _, err = auth.services.Email.Send(ctx.Request().Context(), emailData)
	// if err != nil {
	// 	auth.services.Logger.Error("Error sending email", map[string]interface{}{"err": err})
	// 	return ctx.NoContent(500)
	// }
	if err != nil {
		auth.services.Logger.Error("Error password reset request", map[string]interface{}{"err": err})
		return ctx.NoContent(500)
	}
	return ctx.NoContent(200)
}
