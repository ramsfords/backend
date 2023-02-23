package mid

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/authapi"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/services"
)

func Protected(services *services.Services) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var err error
			loginData := authapi.LoginData{}
			// get cookie named firtAuth
			cookie, err := ctx.Cookie("firstAuth")
			if err == nil && len(cookie.Value) > 50 {
				// decrypt cookie
				decrypted, errs := services.Crypto.Decrypt(cookie.Value)
				if errs != nil {
					err = errs
				}
				// unmarshal decrypted cookie
				errs = json.Unmarshal(decrypted, &loginData)
				if errs != nil {
					err = errs
				}
				validUntil, errs := time.Parse(time.RFC3339, loginData.ValidUntil)
				if err != nil {
					err = errs
				}
				if time.Now().After(validUntil) {
					errs = exchageRefreshTokenForAccesssToken(ctx, loginData, services)
					if errs != nil {
						err = errs
					}
					authapi.WriteCookie(ctx, loginData, services)
				}
			} else {
				authorizationHeader := ctx.Request().Header.Get("authorization")
				if len(authorizationHeader) > 50 {
					// remove cokie not found error
					err = nil
					decrypted, errs := services.Crypto.Decrypt(authorizationHeader)
					if errs != nil {
						err = errs
					}
					// unmarshal decrypted authorizationHeader
					errs = json.Unmarshal(decrypted, &loginData)
					if errs != nil {
						err = errs
					}
					validUntil, errs := time.Parse(time.RFC3339, loginData.ValidUntil)
					if errs != nil {
						err = errs
					}
					if time.Now().After(validUntil) {
						errs = exchageRefreshTokenForAccesssToken(ctx, loginData, services)
						if err != nil {
							err = errs
						}
						authapi.WriteCookie(ctx, loginData, services)
					}
				}
			}
			if err != nil {
				newErr := errors.Wrap(errors.New("users dont have access to this resource"), "error in protected middleware")
				logger.Logger.Error("error in protected middleware", map[string]interface{}{"erorr": errors.Unwrap(newErr)})
				return ctx.NoContent(http.StatusUnauthorized)
			}
			ctx.Set("authContext", loginData)
			return next(ctx)
		}
	}

}
func exchageRefreshTokenForAccesssToken(ctx echo.Context, data authapi.LoginData, services *services.Services) error {
	token, err := services.Db.GetRefreshToken(ctx.Request().Context(), data.UserId)
	if err != nil {
		return err
	}
	if token == "" {
		return errors.New("token not found")
	}
	// exchange refresh token for new token
	_, err = services.CognitoClient.LoginWithRefreshToken(ctx.Request().Context(), token)
	if err != nil {
		return err
	}
	return nil
}
