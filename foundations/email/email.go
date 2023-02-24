package email

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/ramsfords/backend/configs"
)

type Email struct {
	S3Client   *s3.Client
	SESClient  *ses.Client
	HTTPClient httpClient
	Conf       *configs.Config
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(conf *configs.Config, S3Client *s3.Client) Email {
	cfg := aws.Config{
		Region:           "us-west-1",
		Credentials:      conf,
		RetryMaxAttempts: 10,
	}
	SESClient := ses.NewFromConfig(cfg)
	emailCli := Email{
		S3Client:   S3Client,
		SESClient:  SESClient,
		Conf:       conf,
		HTTPClient: &http.Client{},
	}
	return emailCli
}
