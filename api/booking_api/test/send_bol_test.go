package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/email"
)

func TestSendEmailBol(t *testing.T) {
	config := configs.GetConfig()
	// db := db.New(config)
	// booking, err := db.GetBooking(context.Background(), "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	email.New(config)
	data := email.Data{
		To:          []string{"kandelsuren@gmail.com"},
		From:        "quotes@firstshipper.com",
		Subject:     "please find your bol for pro",
		ContentType: "text/html",
		Attachments: []email.Attachment{
			{
				Path: "firstshipperbol/30015-1-asdfd.pdf",
				Type: email.AttachmentTypeS3,
			},
		},
	}
	res, err := email.Send(context.Background(), data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(res)
}
