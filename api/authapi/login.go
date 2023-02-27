package authapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}
type ExtractUserId struct {
	Sub            string `json:"sub"`
	OrganizationId string `json:"custom:organizationId"`
}

func (auth AuthApi) EchoLogin(ctx echo.Context) error {
	// get cookie name firstRefreshToken
	loginReq := &v1.Login{}
	err := ctx.Bind(loginReq)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	loginRes, err := auth.services.CognitoClient.LoginUser(ctx.Request().Context(), loginReq.Email, loginReq.Password)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// jwtBytes
	getClaim, err := auth.services.CognitoClient.Validate(ctx.Request().Context(), *loginRes.AuthenticationResult.IdToken)
	if err != nil {
		logger.Error(err, "could not validate jwt")
	}
	loginResBytes, err := json.Marshal(getClaim.Claims)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// extract user id from jwt
	cognitoUserData := &ExtractUserId{}
	err = json.Unmarshal(loginResBytes, cognitoUserData)
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	err = auth.services.Db.SaveRefreshToken(ctx.Request().Context(), cognitoUserData.OrganizationId, cognitoUserData.Sub, *loginRes.AuthenticationResult.RefreshToken)
	if err != nil {
		logger.Error(err, "RedirectLogin SaveRefreshToken : error in inserting refresh token into the database")
	}

	loginData := model.LoginData{
		Email:          loginReq.Email,
		OrganizationId: cognitoUserData.OrganizationId,
		UserId:         cognitoUserData.Sub,
	}
	token, err := auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		// log error
		logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	}
	loginData.Token = token
	// encrypt with token again
	token, err = auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		// log error
		logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	}
	WriteCookie(ctx, loginData, auth.services)
	return ctx.JSON(http.StatusOK, loginData)

}

func WriteCookie(ctx echo.Context, loginData model.LoginData, services *services.Services) error {
	validUntil := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	loginData.ValidUntil = validUntil
	token, err := services.Crypto.Encrypt(loginData)
	if err != nil {
		// log error
		logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	}
	secure := false
	url := "127.0.0.1"
	if services.Conf.Env == "prod" {
		secure = true
		url = "https://firstshipper.com"
	}
	ctx.Response().Header().Set(echo.HeaderAuthorization, token)

	cookie := new(http.Cookie)
	cookie.Name = "firstAuth"
	cookie.Value = token
	cookie.Path = "/"
	cookie.Domain = url
	cookie.Secure = secure
	cookie.Expires = time.Now().Add(24 * time.Hour * 365)
	ctx.SetCookie(cookie)
	return nil
}
