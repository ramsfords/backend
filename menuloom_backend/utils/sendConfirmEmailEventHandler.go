package utils

import (
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/foundations/zohomail"
)

func SendConfrimEmailEventHandler(email *zohomail.Email, senderName string, senderEmail string) func(e *core.MailerRecordEvent) error {
	return func(e *core.MailerRecordEvent) error {
		token, ok := e.Meta["token"].(string)
		if !ok {
			return errors.New("could not find token for confirm email")
		}
		userName := e.Record.GetString("username")
		redirectLink := fmt.Sprintf("http://localhost:3000/confirm-email?token=%s", token)
		data := zohomail.EmailData{
			ReceiverEmail: e.Record.Email(),
			ReceiverName:  userName,
			EmailSubject:  "Menuloom: Confirm your email",
			RedirectLink:  redirectLink,
			SenderEmail:   senderName,
			SenderName:    senderName,
		}
		err := email.SendConfirmEmail(data)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
