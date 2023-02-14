package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/email"
	template "github.com/ramsfords/backend/email"
	"github.com/ramsfords/backend/services"
	v1 "github.com/ramsfords/types_gen/v1"
)

func SendConfrimEmailEventHandler(services *services.Services) func(e *core.MailerRecordEvent) error {
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

		redirectLink := fmt.Sprintf(services.Conf.GetFirstShipperFontEndURL()+"/confirm-email?token=%s&email=%s", token, emailID)
		templeHtml := email.GetConfirmEmailHtml(name, redirectLink)
		data := template.Data{
			To:          []string{emailID},
			Subject:     "FirstShipper: Please Confirm your email!",
			From:        services.Conf.SitesSettings.FirstShipper.Prod.EmailId,
			ContentType: template.ContentTypeTextHTML,
			Body:        templeHtml,
		}
		// Send the email.
		if _, err := template.Send(context.Background(), data, services.Conf); err != nil {
			panic(err)
		}
		originData := e.Record.OriginalCopy()
		userBytes, err := originData.MarshalJSON()
		if err != nil {
			return err
		}
		fmt.Println(originData)
		userData := v1.User{}
		err = json.Unmarshal(userBytes, &userData)
		if err != nil {
			return err
		}
		userData.Email = emailID
		userData.PasswordHash = e.Record.GetString("passwordHash")
		userData.Token = token
		userData.Type = "user"
		err = services.Db.SaveUser(context.Background(), userData, emailID)
		if err != nil {
			return err
		}
		// err = services.Email.SendConfirmEmail(data)
		// if err != nil {
		// 	return err
		// }
		createdAt := e.Record.GetString("createAt")
		if emailID == "" {
			return errors.New("could not find email for confirm email")
		}
		recordData := e.Record.Get("passwordHash")
		fmt.Print(recordData)
		businessData := v1.Business{
			BusinessId:                        emailID,
			AccountingEmail:                   emailID,
			CustomerServiceEmail:              emailID,
			AdminEmail:                        emailID,
			CreateAt:                          createdAt,
			NeedsAddressUpdate:                true,
			NeedsDefaultPickupAddressUpdate:   true,
			NeedsDefaultDeliveryAddressUpdate: true,
			Address:                           &v1.Address{},
			AdminUser:                         &userData,
		}

		err = services.Db.SaveBusiness(context.Background(), businessData, emailID)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
