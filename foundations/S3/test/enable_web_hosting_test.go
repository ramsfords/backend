package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func TestEnableWebHosting(t *testing.T) {
	bucketName := "test-my-bucket-create"
	_, err := S3Client.PutBucketWebsite(context.TODO(), &s3.PutBucketWebsiteInput{
		Bucket: aws.String(bucketName),
		WebsiteConfiguration: &types.WebsiteConfiguration{
			IndexDocument: &types.IndexDocument{
				Suffix: aws.String("index.html"),
			},
		},
	})
	if err != nil {
		// Handle error
		fmt.Printf("Couldn't create bucket. Here's why: %v\n", err)
	}
}
