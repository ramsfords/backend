package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/db"
	"github.com/ramsfords/backend/foundations/S3"
	"github.com/ramsfords/backend/foundations/email"
)

func TestSendEmailBol(t *testing.T) {
	config := configs.GetConfig()
	db := db.New(config)
	s3client := S3.New(config)
	booking, err := db.GetBooking(context.Background(), "50099")
	if err != nil {
		fmt.Println(err)
	}
	emailCli := email.New(config, s3client.Client)
	fileName := strings.Split(booking.BookingInfo.BolUrl, ".com/")[1]
	data := email.Data{
		To:          []string{"kandelsuren@gmail.com"},
		From:        "quotes@firstshipper.com",
		Subject:     "please find your bol for pro",
		ContentType: "text/html",
		Attachments: []email.Attachment{
			{
				Path: "firstshipperbol/" + fileName,
				Type: email.AttachmentTypeS3,
			},
		},
	}
	res, err := emailCli.Send(context.Background(), data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(res)
}
