package S3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (S3 S3Client) AddPublicReadBucketPolicy(bucketName string) error {
	_, err := S3.PutBucketPolicy(context.TODO(), &s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Sid": "PublicReadGetObject",
					"Effect": "Allow",
					"Principal": "*",
					"Action": "s3:GetObject",
					"Resource": "arn:aws:s3:::` + bucketName + `/*"
				}
			]
		}`),
	})
	if err != nil {
		// Handle error
		fmt.Printf("Couldn't create bucket. Here's why: %v\n", err)
	}
	return nil
}
