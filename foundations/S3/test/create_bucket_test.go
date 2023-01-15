package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func TestAll(t *testing.T) {
	t.Run("TestCreateBucket", TestCreateBucket)
	t.Run("TestEnableWebHosting", TestEnableWebHosting)
	t.Run("TestAddBucketPolicy", TestAddBucketPolicy)

}
func TestCreateBucket(t *testing.T) {
	bucketName := "test-my-bucket-create"
	readpolicy := "public-read"
	region := "us-west-1"
	_, err := S3Client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		ACL:    types.BucketCannedACL(readpolicy),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't create bucket. Here's why: %v\n", err)
	}

}
