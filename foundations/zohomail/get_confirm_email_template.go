package zohomail

import (
	"fmt"
	"strings"
)

func (email Email) getConfirmTemplate(data EmailData) string {
	template := strings.ReplaceAll(confirmEmailTemplate, "{user_name}", data.ReceiverName)
	template = strings.ReplaceAll(template, "{redirect_link}", data.RedirectLink)
	fmt.Println(template)
	return template
}
