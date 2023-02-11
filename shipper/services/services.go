package services

import (
	"fmt"
	"log"

	"github.com/ramsfords/backend/configs"
	email "github.com/ramsfords/backend/email"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/auth"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/shipper/business/rapid"
	"github.com/ramsfords/backend/shipper/db"
	"go.uber.org/zap"
)

type Services struct {
	Conf       *configs.Config
	S3Client   S3.S3Client
	Logger     logger.Logger
	Cloudinery *cloudinery.Cloudinery
	Rapid      *rapid.Rapid
	Zoho       *auth.Zoho
	Db         db.DB
}

func New(conf *configs.Config) *Services {
	S3Client := S3.New(conf)
	logger := logger.New(conf.GetFirstShipperServiceName())
	email.New(conf)
	db := db.New(conf)
	cloudinery := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)
	zohoClient := auth.New(conf)
	// z.AuthorizationCodeRequest("1000.MC1OXITE0VDZA0E996T3VQKKQKTEFM", "7700f6385a550ec7fc38d451da76e1147877e315e9", []ScopeString{"ZohoBooks.fullaccess.all"}, "https://www.google.Com")
	// to start oAuth2 flow
	// scopes := []zoho.ScopeString{
	// 	zoho.BuildScope(zoho.Crm, zoho.ModulesScope, zoho.AllMethod, zoho.NoOp),
	// }
	if err := zohoClient.GenerateTokenRequest(conf.Zoho.ZohoClientId, conf.Zoho.ZohoClientSecret, conf.Zoho.ZohoCode, "https://api.firstshipper.Com"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("access Token \n", zohoClient.Oauth.Token.AccessToken)
	rapid := rapid.New()
	newService := &Services{
		Conf:       conf,
		S3Client:   S3Client,
		Logger:     logger,
		Cloudinery: cloudinery,
		Db:         db,
		Zoho:       zohoClient,
		Rapid:      rapid,
	}
	return newService
}
func (ser Services) GetLogger() *zap.SugaredLogger {
	return ser.Logger.SugaredLogger
}
