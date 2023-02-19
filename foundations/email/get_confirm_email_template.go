package email

import "strings"

func GetConfirmEmailHtml(receiverName string, redirectLink string) string {
	template := strings.ReplaceAll(confirmEmailTemplate, "{userName}", receiverName)
	template = strings.ReplaceAll(template, "{redirectLink}", redirectLink)
	return template
}
