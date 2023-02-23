package services

import (
	"fmt"

	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/business/rapid"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/db"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/auth"
	"github.com/ramsfords/backend/foundations/cloudflare"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/cognito"
	"github.com/ramsfords/backend/foundations/email"
	"github.com/ramsfords/backend/foundations/logger"
)

type Services struct {
	Conf          *configs.Config
	S3Client      S3.S3Client
	Cloudinery    *cloudinery.Cloudinery
	CloudFlare    *cloudflare.Cloudflare
	CognitoClient *cognito.CognitoClient
	Rapid         *rapid.Rapid
	Zoho          *auth.Zoho
	Db            db.DB
	Email         *email.Email
	Logger        *logger.AppLogger
	Crypto        *model.Crypto
}

func New(conf *configs.Config) *Services {
	S3Client := S3.New(conf)
	logger, err := logger.New(conf, "firstshipper-api")
	if err != nil {
		fmt.Println(fmt.Errorf("could not start cognito client: %v", err))
		return nil

	}
	emailCli := email.New(conf, S3Client.Client)
	db := db.New(conf)
	cloudinery := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)
	cloudflar := cloudflare.New(conf.CloudFlareConfig)
	cogClient, err := cognito.NewClient(conf)
	if err != nil {
		logger.Error("could not start cognito client", map[string]interface{}{"error": err})
	}
	crypto := model.New(conf.FirstKey)
	// zohoClient := auth.New(conf)
	// if err := zohoClient.GenerateTokenRequest(conf.Zoho.ZohoClientId, conf.Zoho.ZohoClientSecret, conf.Zoho.ZohoCode, "https://api.firstshipper.Com"); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("access Token \n", zohoClient.Oauth.Token.AccessToken)
	rapid := rapid.New()
	newService := &Services{
		Conf:          conf,
		S3Client:      S3Client,
		Cloudinery:    cloudinery,
		CloudFlare:    cloudflar,
		CognitoClient: cogClient,
		Db:            db,
		Zoho:          &auth.Zoho{},
		Rapid:         rapid,
		Crypto:        crypto,
		Logger:        logger,
		Email:         emailCli,
	}
	return newService
}
