package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v5"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IdToken      string `json:"id_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}
type LoginResponse struct {
	Token      `json:"token"`
	LoggedUser string `json:"loggedUser"`
}

func (auth Auth) EchoLogin(ctx echo.Context) error {
	// get cookie name firstRefreshToken
	code := ctx.QueryParam("code")
	tkn := Token{}
	if code == "" || code == "null" {
		_, err := ctx.Cookie("firstAccessToken")
		if err != nil {
			// get refresh token
			_, tokenErr := ctx.Cookie("firstAccessToken")
			refreshToken, refreshErr := ctx.Cookie("firstRefreshToken")
			if refreshErr != nil && tokenErr != nil {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			return auth.ExchageRefreshTokenForToken(ctx, refreshToken.Value)
		}
		// get cookie name firstRefreshToken
		if _, err := ctx.Cookie("firstRefreshToken"); err == nil {
			// cookie exists
			refreshToken, refreshErr := ctx.Cookie("firstRefreshToken")
			if refreshErr != nil {
				return ctx.NoContent(http.StatusUnauthorized)
			}
			return auth.ExchageRefreshTokenForToken(ctx, refreshToken.Value)
		}
	} else {
		var err error
		tkn, err = auth.ExchageAuthCodeForToken(code)
		if err != nil || tkn.AccessToken == "" {
			return ctx.NoContent(http.StatusUnauthorized)
		}
	}
	return writeCookie(ctx, tkn, auth)

}
func (auth Auth) ExchageAuthCodeForToken(code string) (Token, error) {
	client := &http.Client{}
	encodedBody := fmt.Sprintf("code=%s&grant_type=authorization_code&redirect_uri=%s", code, auth.services.CognitoClient.RedirectUrl)
	var data = strings.NewReader(encodedBody)
	req, err := http.NewRequest("POST", auth.services.CognitoClient.AuthUrl+"/oauth2/token", data)
	if err != nil {
		return Token{}, err
	}
	req.Header.Set("Authorization", "Basic "+auth.services.CognitoClient.BaseAuth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return Token{}, err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return Token{}, err
	}
	fmt.Print(string(bodyText))
	tkn := &Token{}
	err = json.Unmarshal(bodyText, tkn)
	if err != nil {
		return *tkn, err
	}
	return *tkn, nil
}
func (auth Auth) ExchageRefreshTokenForToken(ctx echo.Context, refreshToken string) error {
	client := &http.Client{}
	encodedBody := fmt.Sprintf("refresh_token=%s&grant_type=refresh_token&client_id=%s", refreshToken, auth.services.CognitoClient.CognitoClientID)
	var data = strings.NewReader(encodedBody)
	req, err := http.NewRequest("POST", auth.services.CognitoClient.AuthUrl+"/oauth2/token", data)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Basic "+auth.services.CognitoClient.BaseAuth)
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
	tkn := &Token{}
	err = json.Unmarshal(bodyText, tkn)
	if err != nil {
		return err
	}
	tkn.RefreshToken = refreshToken
	return writeCookie(ctx, *tkn, auth)
}
func writeCookie(ctx echo.Context, token Token, auth Auth) error {
	secure := false
	url := "127.0.0.1"
	if auth.services.Conf.Env == "prod" {
		secure = true
		url = "https://firstshipper.com"
	}
	time.Sleep(1 * time.Second)
	tokens, err := auth.services.CognitoClient.Validate(ctx.Request().Context(), token.IdToken)
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
	tokenResponse := LoginResponse{
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
	baseAuth := base64.RawStdEncoding.EncodeToString(tokenStr)
	cookie := new(http.Cookie)
	cookie.Name = "firstAuth"
	cookie.Value = baseAuth
	cookie.Path = "/"
	cookie.Domain = url
	cookie.Secure = secure
	cookie.Expires = time.Now().Add(60 * time.Minute)
	ctx.SetCookie(cookie)
	auth.services.CloudFlare.AddTokenToCloudFlareKV(token.AccessToken, token.RefreshToken)
	return ctx.JSON(http.StatusOK, tokenResponse)
}
