package S3

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

func (S3 S3Client) UploadFile(webSiteUrl string, file *os.File, key string,
	contentType ContentType) error {
	bucketName := "/" + webSiteUrl
	_, err := S3.Upload(context.Background(), &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &key,
		ContentType: aws.String(string(contentType)),
		Body:        file,
		ACL:         types.ObjectCannedACL("public-read"),
	})
	if err != nil {
		return err
	}

	return nil
}
