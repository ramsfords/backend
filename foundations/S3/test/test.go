package test

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ramsfords/backend/configs"
)

var Uploader *manager.Uploader
var conf configs.Config
var S3Client *s3.Client

func init() {
	Uploader, S3Client, conf = getS3Client()
}
func getS3Client() (S3Manager *manager.Uploader, client *s3.Client, conf configs.Config) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	uploader := manager.NewUploader(s3Client)
	return uploader, s3Client, conf
}
