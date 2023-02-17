package mid

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/api/auth"
	"github.com/ramsfords/backend/services"
)

func Protected(services *services.Services) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			letGo := false
			cookie, err := ctx.Cookie("firstAuth")
			authorizationKey := ctx.Request().Header.Get("authorization")
			if err != nil && len(authorizationKey) < 10 {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			if authorizationKey == "" && cookie.Value != "" {
				authorizationKey = cookie.Value
				letGo = true
			}
			token := &auth.LoginResponse{}
			authorizationKey = strings.Split(authorizationKey, "Bearer ")[1]
			if err != nil {
				return fmt.Errorf("could not decode base64 string: %v", err)
			}
			err = json.Unmarshal([]byte(authorizationKey), token)
			if err != nil {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			if token.Token.AccessToken == "" && token.Token.RefreshToken == "" {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			if token.Token.AccessToken != "" {
				tkn, err := services.CognitoClient.Validate(ctx.Request().Context(), token.Token.AccessToken)
				if tkn.Valid && err == nil {
					letGo = true
				} else if token.Token.RefreshToken != "" {
					err := ExchageRefreshTokenForToken(ctx, token.Token.RefreshToken, services)
					if err != nil {
						return ctx.NoContent(http.StatusUnauthorized)
					}
					letGo = true
				}
				letGo = true
			} else if token.Token.RefreshToken != "" {
				err := ExchageRefreshTokenForToken(ctx, token.Token.RefreshToken, services)
				if err != nil {
					return ctx.NoContent(http.StatusUnauthorized)
				}
			}
			if !letGo {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			return next(ctx)
		}
	}

}
func ExchageRefreshTokenForToken(ctx echo.Context, refreshToken string, services *services.Services) error {
	cognitoConf := services.Conf.GetAwsConfig()
	client := &http.Client{}
	encodedBody := fmt.Sprintf("refresh_token=%s&grant_type=refresh_token&client_id=%s", refreshToken, cognitoConf.CognitoClientID)
	var data = strings.NewReader(encodedBody)
	req, err := http.NewRequest("POST", cognitoConf.CognitoUrl+"/oauth2/token", data)
	if err != nil {
		return err
	}
	baseAuth := base64.RawStdEncoding.EncodeToString([]byte(cognitoConf.CognitoClientID + ":" + cognitoConf.CognitoClientSecret))
	req.Header.Set("Authorization", "Basic "+baseAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Print(string(bodyText))
	tkn := &auth.Token{}
	err = json.Unmarshal(bodyText, tkn)
	if err != nil {
		return err
	}
	tkn.RefreshToken = refreshToken
	return writeCookie(ctx, *tkn, services)
}
func writeCookie(ctx echo.Context, token auth.Token, services *services.Services) error {
	secure := false
	url := "127.0.0.1"
	if services.Conf.Env == "prod" {
		secure = true
		url = "https://firstshipper.com"
	}
	time.Sleep(1 * time.Second)
	tokens, err := services.CognitoClient.Validate(ctx.Request().Context(), token.IdToken)
	if err != nil {
		return err
	}
	type Data struct {
		Email string `json:"email"`
	}
	emailData := Data{}
	byts, err := json.Marshal(tokens.Claims)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byts, &emailData)
	if err != nil {
		return err
	}
	tokenResponse := auth.LoginResponse{
		Token:      token,
		LoggedUser: emailData.Email,
	}
	if emailData.Email == "" {
		return ctx.JSON(http.StatusOK, map[string]string{"loggedUser": ""})
	}

	tokenStr, err := json.Marshal(tokenResponse)
	if err != nil {
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "firstAuth"
	cookie.Value = string(tokenStr)
	cookie.Path = "/"
	cookie.Domain = url
	cookie.Secure = secure
	cookie.Expires = time.Now().Add(60 * time.Minute)
	ctx.SetCookie(cookie)
	services.CloudFlare.AddTokenToCloudFlareKV(token.AccessToken, token.RefreshToken)
	return nil
}
