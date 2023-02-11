package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/email"
	template "github.com/ramsfords/backend/foundations/email"
)

func SendResetPasswordLinkEventHandler(conf *configs.Config) func(e *core.MailerRecordEvent) error {
	return func(e *core.MailerRecordEvent) error {
		token, ok := e.Meta["token"].(string)
		if !ok {
			return errors.New("could not find token for confirm email")
		}
		userName := e.Record.GetString("username")
		id := e.Record.Id
		redirectLink := fmt.Sprintf("http://localhost:3000/forgot-password?token=%s&id=%s", token, id)
		emailID := e.Record.GetString("email")
		if emailID == "" {
			return errors.New("could not find email for confirm email")
		}
		resetEmailTemplate := template.GetResetPasswordTemplate(userName, redirectLink)
		data := email.Data{
			To:          []string{emailID},
			Subject:     "FirstShipper: Please Confirm your email!",
			From:        conf.SitesSettings.FirstShipper.Prod.EmailId,
			ContentType: email.ContentTypeTextHTML,
			Body:        resetEmailTemplate,
		}
		sentRes, err := email.Send(context.Background(), data)
		fmt.Println("email sent res", sentRes)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
