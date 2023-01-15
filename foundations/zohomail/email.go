package zohomail

import (
	"net/smtp"

	"github.com/jhillyerd/enmime"
	"github.com/ramsfords/backend/configs"
)

type EmailData struct {
	RedirectLink  string `json:"redirect_link"`
	ReceiverEmail string `json:"receiver_email"`
	ReceiverName  string `json:"receiver_first_name"`
	SenderEmail   string `json:"sender_email"`
	SenderName    string `json:"sender_name"`
	EmailSubject  string `json:"email_subject"`
}

type EmailSender interface {
	SendConfirmEmail(data EmailData) error
	SendPasswordReset(data EmailData) error
	SendWelcomeEmail(data EmailData) error
}
type Email struct {
	conf configs.Email
	*enmime.SMTPSender
}

func New(conf configs.Email) *Email {
	// smtpHost := "smtp.zoho.com:587"
	smtpAuth := smtp.PlainAuth("", conf.UserName, conf.Password, conf.SmtpHost)
	return &Email{
		conf:       conf,
		SMTPSender: enmime.NewSMTP(conf.SmtpHost+":"+conf.SmtpPort, smtpAuth),
	}
}
