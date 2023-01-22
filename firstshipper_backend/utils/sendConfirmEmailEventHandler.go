package utils

import (
	"context"
	"errors"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
	"github.com/ramsfords/backend/configs"
	"github.com/ramsfords/backend/firstshipper_backend/services"
	"github.com/ramsfords/backend/foundations/zohomail"
	v1 "github.com/ramsfords/types_gen/v1"
)

func SendConfrimEmailEventHandler(email *zohomail.Email, conf *configs.Config, services *services.Services) func(e *core.MailerRecordEvent) error {
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

		redirectLink := fmt.Sprintf(conf.GetFirstShipperFontEndURL()+"/confirm-email?token=%s&email=%s", token, emailID)
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
		createdAt := e.Record.GetString("createAt")
		if emailID == "" {
			return errors.New("could not find email for confirm email")
		}
		businessData := v1.Business{
			BusinessName:                      name,
			BusinessId:                        emailID,
			AccountingEmail:                   emailID,
			CustomerServiceEmail:              emailID,
			AdminEmail:                        emailID,
			CreateAt:                          createdAt,
			NeedsAddressUpdate:                true,
			NeedsDefaultPickupAddressUpdate:   true,
			NeedsDefaultDeliveryAddressUpdate: true,
		}
		err = services.SaveBusiness(context.Background(), businessData)
		if err != nil {
			return err
		}
		return hook.StopPropagation
	}

}
