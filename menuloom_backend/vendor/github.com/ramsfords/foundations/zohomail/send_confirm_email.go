package zohomail

func (email Email) SendConfirmEmail(data EmailData) error {
	return email.SendEmail(email.getConfirmTemplate(data), data)
}
