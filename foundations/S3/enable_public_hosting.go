package S3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func (S3 S3Client) EnableWebHosting(bucketName string, indexFileName WebSiteRouteFile, errorFileName WebSiteRouteFile) error {
	_, err := S3.PutBucketWebsite(context.TODO(), &s3.PutBucketWebsiteInput{
		Bucket: aws.String(bucketName),
		WebsiteConfiguration: &types.WebsiteConfiguration{
			IndexDocument: &types.IndexDocument{
				Suffix: aws.String(string(indexFileName)),
			},
			ErrorDocument: &types.ErrorDocument{
				Key: aws.String(string(errorFileName)),
			},
		},
	})
	if err != nil {
		// Handle error
		fmt.Printf("Couldn't create bucket. Here's why: %v\n", err)
	}
	return nil
}
