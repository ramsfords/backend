package authapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

func (auth AuthApi) RedirectLogin(ctx echo.Context) error {
	token := ctx.QueryParam("token")
	if len(token) < 50 {
		return ctx.NoContent(http.StatusBadRequest)
	}
	data, err := auth.services.Crypto.Decrypt(token)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	userLogin := &RedirectLoginData{}
	// json unmarshal data to user
	err = json.Unmarshal(data, userLogin)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	userLogin.Token = token
	loginRes, err := auth.services.CognitoClient.LoginUser(ctx.Request().Context(), userLogin.Email, userLogin.Password)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = auth.services.Db.SaveRefreshToken(ctx.Request().Context(), userLogin.UserId, *loginRes.AuthenticationResult.RefreshToken)
	if err != nil {
		auth.services.Logger.Errorf("RedirectLogin SaveRefreshToken : error in inserting refresh token into the database: %s", err)
	}
	loginData := &LoginData{
		UserId:         userLogin.UserId,
		Email:          userLogin.Email,
		ValidUntil:     time.Now().Add(time.Hour * 1).Format(time.RFC3339),
		OrganizationId: userLogin.OrganizationId,
	}
	// encrypt the data
	token, err = auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		auth.services.Logger.Errorf("RedirectLogin SaveRefreshToken : error in encrypting login data: %s", err)
	}
	// encrypt the data
	loginData.Token = token
	token, err = auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		auth.services.Logger.Errorf("RedirectLogin SaveRefreshToken : error in encrypting login data: %s", err)
	}

	writeCookie(ctx, token, auth)
	return ctx.JSON(http.StatusOK, userLogin)
}
