package zohomail

import (
	"strings"
)

func (email Email) getResetPasswordTemplate(data EmailData) string {
	template := strings.ReplaceAll(passwordResetEmailTemplate, "{user_name}", data.ReceiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", data.RedirectLink)
	return template
}
