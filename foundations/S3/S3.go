package S3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ramsfords/backend/configs"
)

type S3Client struct {
	*s3.Client
	*configs.Config
	*manager.Uploader
}

// config satisfies the CredentialsProvider interface
func New(conf *configs.Config) S3Client {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return S3Client{}
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	uploader := manager.NewUploader(s3Client)
	return S3Client{
		Client:   s3Client,
		Config:   conf,
		Uploader: uploader,
	}

}

type CannedPolicy string

const PublicRead CannedPolicy = "public-read"
const AuthenticatedRead CannedPolicy = "authenticated-read"
const BucketOwnerRead CannedPolicy = "bucket-owner-read"
const BucketOwnerFullControl CannedPolicy = "bucket-owner-full-control"
const LogDeliveryWrite CannedPolicy = "log-delivery-write"
const PUblicReadPublicWrite CannedPolicy = "public-read-write"
const AwsExecRead CannedPolicy = "aws-exec-read"
const BucketOwenrFullControl CannedPolicy = "bucket-owner-full-control"
const Private CannedPolicy = "private"

type Region string

const West1 Region = "us-west-1"

type ContentType string

const ImageJpg ContentType = "image/jpg"
const ImagePng ContentType = "image/png"
const ImageGif ContentType = "image/gif"
const ImageTiff ContentType = "image/tiff"
const ImageWebp ContentType = "image/webp"
const ImageBmp ContentType = "image/bmp"
const ImageIco ContentType = "image/vnd.microsoft.icon"
const ImageSvg ContentType = "image/svg+xml"
const ImageAvif ContentType = "image/avif"
const Jpeg ContentType = "image/jpeg"
const Pdf ContentType = "application/pdf"
const TextPlain ContentType = "text/plain"
const TextHtml ContentType = "text/html"
const TextCss ContentType = "text/css"
const TextCsv ContentType = "text/csv"
const TextJavascript ContentType = "text/javascript"
const TextXml ContentType = "text/xml"
const TextMarkdown ContentType = "text/markdown"

type WebSiteRouteFile string

const IndexFile WebSiteRouteFile = "index.html"
const ErrorFile WebSiteRouteFile = "error.html"
