package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func TestDownloadImageFromURL(t *testing.T) {
	// Get the response bytes from the url

	URL := "https://images.unsplash.com/photo-1672841828271-54340a6fbcd3?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8"
	validUrl, err := url.ParseRequestURI(URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(validUrl)
	// filename := path.Base(URL)
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}

	if response.StatusCode != 200 {
		fmt.Println("Received non 200 response code")
	}
	// // Create a empty file
	// file, err := os.Create(filename)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()

	// Write the bytes to the fiel
	//readData, err := io.ReadAll(response.Request.Body)
	res, err := Uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket:             aws.String("static.menuloom.com"),
		Key:                aws.String("www.himalayen.com/images/test"),
		ContentType:        aws.String(response.Header.Get("Content-Type")),
		ContentDisposition: aws.String("inline"),
		Body:               response.Body,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v\n", err)
	}
	fmt.Println(res)

}
