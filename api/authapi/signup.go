package authapi

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/email"
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (auth AuthApi) EchoSignUp(ctx echo.Context) error {
	data := &v1.User{}
	if err := ctx.Bind(data); err != nil || data.Email == "" || data.Password == "" || data.Password != data.ConfirmPassword {
		return ctx.NoContent(http.StatusBadRequest)
	}
	data.UserName = strings.ToLower(data.Email)
	data.Email = strings.ToLower(data.Email)
	data.Name = strings.ToLower(data.Name)
	orgId := auth.services.Db.GetBusinessCount()
	orgIdStr := fmt.Sprintf("%d", orgId)
	data.OrgId = orgIdStr
	userId, err := auth.services.CognitoClient.CreateUser(ctx.Request().Context(), data)
	if err != nil {
		if err == errs.ErrUserAlreadyExits {
			return ctx.NoContent(http.StatusConflict)
		}
	}
	data.Id = *userId.UserSub
	go auth.services.Db.IncreaseBusinessCount()
	confirmEmailToken, err := auth.services.Crypto.Encrypt(data)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	validUntil := time.Now().Add(24 * time.Hour)
	validUntilStr, err := auth.services.Crypto.Encrypt(validUntil)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	redirectLink := auth.services.Conf.GetFirstShipperFontEndURL() + "/confirm-email?token=" + confirmEmailToken + "next" + validUntilStr
	emailData := email.Data{
		To:          []string{data.Email},
		CC:          &[]string{"noreply@firstshipper.com"},
		From:        "noreply@firstshipper.com",
		Subject:     "Please Confirm your email",
		ContentType: "text/html",
		Body:        email.GetConfirmEmailHtml(data.Name, redirectLink),
	}
	_, err = auth.services.Email.Send(ctx.Request().Context(), emailData)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}
	return ctx.NoContent(http.StatusCreated)
}
