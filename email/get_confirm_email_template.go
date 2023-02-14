package email

import (
	"strings"
)

//	func (email Email) getConfirmTemplate(data EmailData) string {
//		template := strings.ReplaceAll("confirmEmailTemplate", "{user_name}", data.ReceiverName)
//		template = strings.ReplaceAll(template, "{redirect_link}", data.RedirectLink)
//		fmt.Println(template)
//		return template
//	}
func GetConfirmEmailHtml(receiverName string, redirectLink string) string {
	template := strings.ReplaceAll("confirmEmailTemplate", "{user_name}", receiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", redirectLink)
	return template
}
