package zohomail

import (
	_ "embed"
)

//go:embed templates/email-confirmation.html
var confirmEmailTemplate string

//go:embed templates/password-reset.html
var passwordResetEmailTemplate string

//go:embed templates/welcome.html
var welcomeEmailTemplate string
