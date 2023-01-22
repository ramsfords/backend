package utils

import (
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/foundations/zohomail"
	"github.com/ramsfords/backend/menuloom_backend/db"
)

func SendConfrimEmailEventHandler(email *zohomail.Email, senderName string, senderEmail string, businessDb db.Repository) func(e *core.MailerRecordEvent) error {
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
		data := zohomail.EmailData{
			ReceiverEmail: emailID,
			ReceiverName:  name,
			EmailSubject:  "FirstShipper: Please Confirm your email!",
			RedirectLink:  redirectLink,
			SenderEmail:   "quotes@firstshipper.com",
			SenderName:    "FirstShipper",
		}

		err := email.SendConfirmEmail(data)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
