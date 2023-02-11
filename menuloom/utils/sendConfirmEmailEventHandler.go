package utils

import (
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	email "github.com/ramsfords/backend/foundations/zoho/email"
	"github.com/ramsfords/backend/menuloom/services"
)

func SendConfrimEmailEventHandler(services *services.Services, senderName string, senderEmail string) func(e *core.MailerRecordEvent) error {
	return func(e *core.MailerRecordEvent) error {
		token, ok := e.Meta["token"].(string)
		if !ok {
			return errors.New("could not find token for confirm email")
		}
		emailID := e.Record.GetString("email")
		if emailID == "" {
			return errors.New("could not find email for confirm email")
		}

		name := e.Record.GetString("name")

		redirectLink := fmt.Sprintf("http://127.0.0.1:3000/confirm-email?email=%s&token=%s", token, emailID)
		data := email.EmailData{
			ReceiverEmail: emailID,
			ReceiverName:  name,
			EmailSubject:  "FirstShipper: Please Confirm your email!",
			RedirectLink:  redirectLink,
			SenderEmail:   "quotes@firstshipper.com",
			SenderName:    "FirstShipper",
		}

		err := services.Email.SendConfirmEmail(data)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
