package mid

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/authapi"
	"github.com/ramsfords/backend/services"
)

func Protected(services *services.Services) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			letGo := false
			// get cookie named firtAuth
			cookie, err := ctx.Cookie("firstAuth")
			if err == nil && len(cookie.Value) > 50 {
				// decrypt cookie
				decrypted, err := services.Crypto.Decrypt(cookie.Value)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				// unmarshal decrypted cookie
				loginData := &authapi.LoginData{}
				err = json.Unmarshal(decrypted, loginData)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				// set loginData to context
				ctx.Set("authContext", loginData)
				letGo = true
			} else {
				authorizationHeader := ctx.Request().Header.Get("authorization")
				if authorizationHeader == "" || len(authorizationHeader) < 50 {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				// decrypt cookie
				decrypted, err := services.Crypto.Decrypt(cookie.Value)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				// unmarshal decrypted cookie
				loginData := &authapi.LoginData{}
				err = json.Unmarshal(decrypted, loginData)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				// check if token is validUntil now
				validUntil, err := time.Parse(time.RFC3339, loginData.ValidUntil)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
				if time.Now().After(validUntil) {
					// request new token with refresh token
					err := ExchageRefreshTokenForToken(ctx, loginData, services)
					if err != nil {
						return ctx.NoContent(http.StatusUnauthorized)
					}
					loginData.ValidUntil = time.Now().Add(time.Hour * 1).Format(time.RFC3339)
					// encrypt loginData
					encrypted, err := services.Crypto.Encrypt(loginData)
					if err != nil {
						services.Logger.Errorf("Protected Encrypt : error in encrypting login data: %s", err)
					}
					// set cookie
					writeCookie(ctx, encrypted, services.Conf.Env)
					letGo = true
				}
				letGo = true
			}
			if !letGo {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			return next(ctx)
		}
	}

}
func ExchageRefreshTokenForToken(ctx echo.Context, data *authapi.LoginData, services *services.Services) error {
	token, err := services.Db.GetRefreshToken(ctx.Request().Context(), data.UserId)
	if err != nil {
		return err
	}
	if token == "" {
		return errors.New("token not found")
	}
	// exchange refresh token for new token
	cognitLoginRes, err := services.CognitoClient.LoginWithRefreshToken(ctx.Request().Context(), token)
	if err != nil {
		return err
	}
	// set refresh token to db
	err = services.Db.SaveRefreshToken(ctx.Request().Context(), data.Email, *cognitLoginRes.AuthenticationResult.RefreshToken)
	if err != nil {
		services.Logger.Errorf("ExchageRefreshTokenForToken SaveRefreshToken : error in inserting refresh token into the database: %s", err)
		// dont return error here because it should not block the user
		return nil
	}
	return nil
}
func writeCookie(ctx echo.Context, token string, env string) error {
	secure := false
	url := "127.0.0.1"
	if env == "prod" {
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
