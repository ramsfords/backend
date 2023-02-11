package email

import (
	"strings"
)

func (email Email) getResetPasswordTemplate(data EmailData) string {
	template := strings.ReplaceAll(passwordResetEmailTemplate, "{user_name}", data.ReceiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", data.RedirectLink)
	return template
}
func GetResetPasswordTemplate(receiverName string, redirectLink string) string {
	template := strings.ReplaceAll(passwordResetEmailTemplate, "{user_name}", receiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", redirectLink)
	return template
}
