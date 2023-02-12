package emails

import (
	"strings"
)

func (email Email) getWelcomeTemplate(data EmailData) string {
	template := strings.ReplaceAll("welcomeEmailTemplate", "{user_name}", data.ReceiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", data.RedirectLink)
	return template
}
