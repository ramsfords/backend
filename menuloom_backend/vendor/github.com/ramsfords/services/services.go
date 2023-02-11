package services

import (
	"github.com/ramsfords/configs"
	"github.com/ramsfords/foundations/S3"
	"github.com/ramsfords/foundations/cloudinery"
	"github.com/ramsfords/foundations/dynamo"
	"github.com/ramsfords/foundations/logger"
	"github.com/ramsfords/foundations/zohomail"
	"github.com/scylladb/gocqlx/v2"
)

type Services struct {
	conf       configs.Config
	dbClient   dynamo.DB
	S3Client   S3.S3Client
	logger     logger.Logger
	email      zohomail.Email
	Cloudinery cloudinery.Cloudinary
	Scalladb   gocqlx.Session
}

func New(conf configs.Config, dynoClient dynamo.DB, s3Client S3.S3Client,
	logger logger.Logger,
	email zohomail.Email, cloudinery cloudinery.Cloudinary, Scalladb gocqlx.Session) Services {
	newService := Services{
		conf:       conf,
		dbClient:   dynoClient,
		S3Client:   s3Client,
		logger:     logger,
		email:      email,
		Cloudinery: cloudinery,
		Scalladb:   Scalladb,
	}
	return newService
}
func (ser Services) GetConfig() configs.Config {
	return ser.conf
}

func (ser Services) GetDbClient() dynamo.DB {
	return ser.dbClient
}

func (ser Services) GetS3Client() S3.S3Client {
	return ser.S3Client
}

func (ser Services) GetLogger() logger.Logger {
	return ser.logger
}
func (ser Services) GetEmail() zohomail.Email {
	return ser.email
}

func (ser Services) GetCloudinary() cloudinery.Cloudinary {
	return ser.Cloudinery
}
func (ser Services) GetScallaDb() gocqlx.Session {
	return ser.Scalladb
}
