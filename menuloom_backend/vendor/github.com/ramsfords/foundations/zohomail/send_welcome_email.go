package zohomail

func (email Email) SendWelcomeEmail(data EmailData) error {
	return email.SendEmail(email.getWelcomeTemplate(data), data)
}
