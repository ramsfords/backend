package emails

func (email Email) SendPasswordReset(data EmailData) error {
	return email.SendEmail(email.getResetPasswordTemplate(data), data)
}
