package services

import (
	"fmt"

	"github.com/ramsfords/backend/business/core/model"
	"github.com/ramsfords/backend/business/rapid"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/db"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/adobe"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/cognito"
	"github.com/ramsfords/backend/foundations/email"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/foundations/zoho"
	supa "github.com/surendrakandel/supa-go"
)

type Services struct {
	Conf          *configs.Config
	S3Client      S3.S3Client
	Cloudinery    cloudinery.Cloudinery
	CognitoClient *cognito.CognitoClient
	Rapid         *rapid.Rapid
	Zoho          *zoho.Zoho
	Db            db.DB
	Email         email.Email
	Logger        *logger.AppLogger
	AdobeCli      *adobe.Adobe
	Crypto        *model.Crypto
	SupaClient    *supa.Client
}

func New(conf *configs.Config) *Services {
	logger, err := logger.New(conf, "firstshipper-api")
	if err != nil {
		fmt.Println(fmt.Errorf("could not start cognito client: %v", err))
		return nil

	}
	newService := &Services{
		Conf:   conf,
		Zoho:   &zoho.Zoho{},
		Logger: logger,
	}
	supaClient := supa.CreateClient(conf.GetSupaConfig().Url, conf.GetSupaConfig().AnonKey)
	newService.SupaClient = supaClient
	S3Client, err := S3.New(conf)
	if err != nil {
		fmt.Println("could not start s3 client", err)
		return nil
	}
	newService.S3Client = S3Client
	pdfClient, err := adobe.NewAdobe(S3Client, conf)
	if err != nil {
		fmt.Println("could not start adobe client", err)
		return nil
	}
	newService.AdobeCli = pdfClient
	emailCli := email.New(conf, S3Client.Client)
	newService.Email = emailCli
	db, err := db.New(conf)
	if err != nil {
		fmt.Println("could not start dynamodb client", err)
		return nil
	}
	newService.Db = db
	cloudinery, err := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)
	if err != nil {
		fmt.Println("could not start cloudinary client", err)
		return nil
	}
	newService.Cloudinery = cloudinery
	cogClient, err := cognito.NewClient(conf)
	if err != nil {
		logger.Error("could not start cognito client", map[string]interface{}{"error": err})
	}
	newService.CognitoClient = cogClient
	crypto := model.New(conf.FirstKey)
	newService.Crypto = crypto
	zohoClient, err := zoho.New(conf)
	if err != nil {
		logger.Error("could not start zoho client", map[string]interface{}{"error": err})
	}
	newService.Zoho = zohoClient
	rapid := rapid.New()
	newService.Rapid = rapid

	return newService
}
