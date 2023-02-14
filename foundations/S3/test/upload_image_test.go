package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestUploadImage(t *testing.T) {
	fileName := "order.jpg"
	bucketName := "test-my-bucket-create"
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, "objectKey", err)
		}
		_, err = Uploader.Upload(context.Background(), &s3.PutObjectInput{
			Bucket:      aws.String(bucketName),
			Key:         aws.String("images/alina"),
			ContentType: aws.String("image/jpg"),
			Body:        file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, "objectKey", err)
		}
	}
	fmt.Print(err)
}
