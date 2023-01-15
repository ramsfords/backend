package utils

import (
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/foundations/zohomail"
)

func SendResetPasswordLinkEventHandler(email *zohomail.Email, senderName string, senderEmail string) func(e *core.MailerRecordEvent) error {
	return func(e *core.MailerRecordEvent) error {
		token, ok := e.Meta["token"].(string)
		if !ok {
			return errors.New("could not find token for confirm email")
		}
		userName := e.Record.GetString("username")
		id := e.Record.Id
		redirectLink := fmt.Sprintf("http://localhost:3000/forgot-password?token=%s&id=%s", token, id)
		data := zohomail.EmailData{
			ReceiverEmail: e.Record.Email(),
			ReceiverName:  userName,
			EmailSubject:  "Menuloom: Reset Your Password",
			RedirectLink:  redirectLink,
			SenderEmail:   senderEmail,
			SenderName:    senderName,
		}
		err := email.SendPasswordReset(data)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
