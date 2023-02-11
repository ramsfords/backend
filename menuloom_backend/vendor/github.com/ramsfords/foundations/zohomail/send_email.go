package zohomail

import (
	"github.com/jhillyerd/enmime"
)

func (email Email) SendEmail(template string, data EmailData) error {
	master := enmime.Builder().
		From(data.ReceiverName, data.SenderEmail).
		Subject(data.EmailSubject).
		HTML([]byte(template))

	// master is immutable, causing each msg below to have a single recipient.
	msg := master.To(data.ReceiverName, data.ReceiverEmail)
	err := msg.Send(email.SMTPSender)
	if err != nil {
		return err
	}
	return nil
}
