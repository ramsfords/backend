package services

import (
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/foundations/zoho/email"
	"github.com/ramsfords/backend/menuloom_backend/db"
	"go.uber.org/zap"
)

type Services struct {
	Conf       *configs.Config
	S3Client   S3.S3Client
	Logger     logger.Logger
	Email      *email.Email
	Cloudinery *cloudinery.Cloudinery
	Db         db.Db
}

func New(conf *configs.Config) *Services {
	S3Client := S3.New(conf)
	logger := logger.New(conf.GetMenuloomServiceName())
	db := db.New(conf)
	cloudinery := cloudinery.New(conf.SitesSettings.FirstShipper.CloudinaryConfig)
	newService := &Services{
		Conf:       conf,
		S3Client:   S3Client,
		Logger:     logger,
		Cloudinery: cloudinery,
		Db:         db,
	}
	return newService
}
func (ser Services) GetLogger() *zap.SugaredLogger {
	return ser.Logger.SugaredLogger
}
