package authapi

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/utils"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (auth AuthApi) EchoResetPassword(ctx echo.Context) error {
	data := &v1.ResetPassword{}
	err := ctx.Bind(data)
	if err != nil || data.NewPassword != data.ConfirmPassword || !utils.IsEmailValid(data.Email) || len(data.NewPassword) < 6 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	resetTokenEmail, err := auth.services.Crypto.Decrypt(data.Token)
	if err != nil || string(resetTokenEmail) != data.Email {
		return ctx.NoContent(http.StatusBadRequest)
	}

	auth.services.CognitoClient.ResetPassword(ctx.Request().Context(), data.Email, data.NewPassword)
	return ctx.NoContent(http.StatusOK)
}
