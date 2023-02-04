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

//go:embed templates/booking-confirmation.html
var bookingConfirmationEmailTemplate string
