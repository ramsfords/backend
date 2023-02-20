package authapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
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
		auth.services.Logger.Error("could not validate jwt")
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
	err = auth.services.Db.SaveRefreshToken(ctx.Request().Context(), cognitoUserData.Sub, *loginRes.AuthenticationResult.RefreshToken)
	if err != nil {
		auth.services.Logger.Errorf("RedirectLogin SaveRefreshToken : error in inserting refresh token into the database: %s", err)
	}
	validUntil := time.Now().Add(24 * time.Hour).Format(time.RFC3339)
	loginData := LoginData{
		Email:          loginReq.Email,
		OrganizationId: cognitoUserData.OrganizationId,
		UserId:         cognitoUserData.Sub,
		ValidUntil:     validUntil,
	}
	token, err := auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		// log error
		auth.services.Logger.Errorf("EchoLogin Encrypt : error in encrypting login data: %s", err)
	}
	loginData.Token = token
	// encrypt with token again
	token, err = auth.services.Crypto.Encrypt(loginData)
	if err != nil {
		// log error
		auth.services.Logger.Errorf("EchoLogin Encrypt : error in encrypting login data: %s", err)
	}
	writeCookie(ctx, token, auth)
	return ctx.JSON(http.StatusOK, loginData)

}

func writeCookie(ctx echo.Context, token string, auth AuthApi) error {
	secure := false
	url := "127.0.0.1"
	if auth.services.Conf.Env == "prod" {
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
	cookie.Expires = time.Now().Add(60 * time.Minute)
	ctx.SetCookie(cookie)
	return nil
}
