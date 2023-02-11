package zohomail

import (
	"net/smtp"

	"github.com/jhillyerd/enmime"
	"github.com/ramsfords/configs"
)

type EmailData struct {
	RedirectLink  string `json:"redirect_link"`
	ReceiverEmail string `json:"receiver_email"`
	ReceiverName  string `json:"receiver_first_name"`
	SenderEmail   string `json:"sender_email"`
	SenderName    string `json:"sender_name"`
	EmailSubject  string `json:"email_subject"`
}

type ZohoConfig struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	AccountId    string
	SmtpHost     string
	SmtpPort     string
	UserName     string
	Password     string
	FromAddress  string
	SenderName   string
}
type EmailSender interface {
	SendConfirmEmail(data EmailData) error
	SendPasswordReset(data EmailData) error
	SendWelcomeEmail(data EmailData) error
}
type Email struct {
	conf ZohoConfig
	*enmime.SMTPSender
}

func New(conf configs.Config) Email {
	// smtpHost := "smtp.zoho.com:587"
	smtpAuth := smtp.PlainAuth("", conf.ZohoConfig.UserName, conf.Password, conf.SmtpHost)
	return Email{
		conf:       conf,
		SMTPSender: enmime.NewSMTP(conf.SmtpHost+":"+conf.SmtpPort, smtpAuth),
	}
}
