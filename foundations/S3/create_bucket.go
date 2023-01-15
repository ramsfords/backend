package S3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (S3 S3Client) CreateBucket(region Region, bucketName string, bocketPolicy CannedPolicy) (*s3.CreateBucketOutput, error) {
	outPut, err := S3.Client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		ACL:    types.BucketCannedACL(bocketPolicy),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(region),
		},
	})
	if err != nil {
		fmt.Printf("Couldn't create bucket. Here's why: %v\n", err)
		return nil, err
	}
	return outPut, nil
}
