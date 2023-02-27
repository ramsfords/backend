package authapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/configs"
	v1 "github.com/ramsfords/types_gen/v1"
	supabase "github.com/surendrakandel/supa-go"
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
	loginRes, err := auth.services.SupaClient.Auth.SignIn(ctx.Request().Context(), supabase.UserCredentials{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})
	if err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}
	// loginRes, err := auth.services.CognitoClient.LoginUser(ctx.Request().Context(), loginReq.Email, loginReq.Password)
	// if err != nil {
	// 	return ctx.NoContent(http.StatusBadRequest)
	// }
	// // jwtBytes
	// getClaim, err := auth.services.CognitoClient.Validate(ctx.Request().Context(), *loginRes.AuthenticationResult.IdToken)
	// if err != nil {
	// 	logger.Error(err, "could not validate jwt")
	// }
	// loginResBytes, err := json.Marshal(getClaim.Claims)
	// if err != nil {
	// 	return ctx.NoContent(http.StatusBadRequest)
	// }
	// // extract user id from jwt
	// cognitoUserData := &ExtractUserId{}
	// err = json.Unmarshal(loginResBytes, cognitoUserData)
	// if err != nil {
	// 	return ctx.NoContent(http.StatusBadRequest)
	// }
	// err = auth.services.Db.SaveRefreshToken(ctx.Request().Context(), cognitoUserData.OrganizationId, cognitoUserData.Sub, *loginRes.AuthenticationResult.RefreshToken)
	// if err != nil {
	// 	logger.Error(err, "RedirectLogin SaveRefreshToken : error in inserting refresh token into the database")
	// }

	// loginData := model.LoginData{
	// 	Email:          loginReq.Email,
	// 	OrganizationId: cognitoUserData.OrganizationId,
	// 	UserId:         cognitoUserData.Sub,
	// }
	// token, err := auth.services.Crypto.Encrypt(loginData)
	// if err != nil {
	// 	// log error
	// 	logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	// }
	// loginData.Token = token
	// // encrypt with token again
	// token, err = auth.services.Crypto.Encrypt(loginData)
	// if err != nil {
	// 	// log error
	// 	logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	// }
	WriteCookie(ctx, loginRes, auth.services.Conf)
	return ctx.JSON(http.StatusOK, loginRes)

}

func WriteCookie(ctx echo.Context, loginDetail *supabase.AuthenticatedDetails, conf *configs.Config) error {
	// validUntil := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	// loginData.ValidUntil = validUntil
	// token, err := services.Crypto.Encrypt(loginData)
	// if err != nil {
	// 	// log error
	// 	logger.Error(err, "EchoLogin Encrypt : error in encrypting login data")
	// }
	secure := false
	url := "127.0.0.1"
	if conf.Env == "prod" {
		secure = true
		url = "https://firstshipper.com"
	}
	// mashall token into string
	token, err := json.Marshal(loginDetail)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set(echo.HeaderAuthorization, loginDetail.AccessToken)

	cookie := new(http.Cookie)
	cookie.Name = "firstAuth"
	cookie.Value = string(token)
	cookie.Path = "/"
	cookie.Domain = url
	cookie.Secure = secure
	cookie.Expires = time.Now().Add(time.Second * time.Duration(loginDetail.ExpiresIn))
	ctx.SetCookie(cookie)
	return nil
}
