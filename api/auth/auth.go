package auth

import (
	"encoding/base64"
	"log"

	"github.com/labstack/echo/v5"
	"github.com/ramsfords/backend/foundations/cognito"
	"github.com/ramsfords/backend/services"
)

type Auth struct {
	cognitoClient       cognito.CognitoClient
	services            *services.Services
	jwkURL              string
	cognitoRegion       string
	cognitoUserPoolID   string
	cognitoClientSecret string
	cognitoClientID     string
	loginUrl            string
	redirectUrl         string
	baseAuth            string
	authUrl             string
}

func New(services *services.Services, echo *echo.Echo) {
	awsConf := services.Conf.GetAwsConfig()
	baseAuth := base64.RawStdEncoding.EncodeToString([]byte(services.Conf.GetAwsConfig().CognitoClientID + ":" + services.Conf.GetAwsConfig().CognitoClientSecret))
	cogCient, err := cognito.NewClient(services.Conf)
	if err != nil {
		log.Fatal(err)
	}
	auth := Auth{
		services:            services,
		cognitoRegion:       awsConf.CognitoRegion,
		cognitoUserPoolID:   awsConf.CognitoUserPoolID,
		baseAuth:            baseAuth,
		cognitoClientSecret: awsConf.CognitoClientSecret,
		cognitoClientID:     awsConf.CognitoClientID,
		redirectUrl:         awsConf.CognitoRedirectUri,
		authUrl:             awsConf.CognitoUrl,
		jwkURL:              awsConf.JWKUrl,
		loginUrl:            awsConf.LoginUrl,
		cognitoClient:       cogCient,
	}
	protectedBolGroup := echo.Group("/auth-callback")
	protectedBolGroup.GET("/:code", auth.EchoLogin)
	protectedBolGroup.GET("/logout", auth.EchoLogout)
}
