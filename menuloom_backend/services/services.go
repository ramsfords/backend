package services

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/cloudinery"
	"github.com/ramsfords/backend/foundations/logger"
	"github.com/ramsfords/backend/foundations/zohomail"
	"github.com/ramsfords/backend/menuloom_backend/db"
	"github.com/scylladb/gocqlx/v2"
)

type Services struct {
	Conf             *configs.Config
	S3Client         S3.S3Client
	Logger           logger.Logger
	Email            *zohomail.Email
	Cloudinery       *cloudinery.Cloudinery
	Scalladb         gocqlx.Session
	CloudFlareClient *cloudflare.API
	Repository       db.Repository
}

func New(conf *configs.Config,
	S3Client S3.S3Client,
	Logger logger.Logger,
	Email *zohomail.Email,
	Db db.Repository, cloudinery *cloudinery.Cloudinery, cloudFlareClient *cloudflare.API) Services {
	newService := Services{
		Conf:             conf,
		S3Client:         S3Client,
		Logger:           Logger,
		Email:            Email,
		Cloudinery:       cloudinery,
		Repository:       Db,
		CloudFlareClient: cloudFlareClient,
	}
	return newService
}
