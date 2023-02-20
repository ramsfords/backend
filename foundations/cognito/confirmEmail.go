package cognito

import (
	"context"
	"strings"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/ramsfords/backend/foundations/errs"
)

func (cognito *CognitoClient) ConfirmEmail(ctx context.Context, email string) error {
	confirmRequest := cip.AdminConfirmSignUpInput{
		UserPoolId: &cognito.CognitoUserPoolID,
		Username:   &email,
	}
	_, err := cognito.Client.AdminConfirmSignUp(ctx, &confirmRequest)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "ExpiredCodeException") {
			return &errs.ErrConfirmEmailExpired
		}
		if strings.Contains(errMsg, "LimitExceededException") {
			return &errs.ErrTooManyRequest
		}
		if strings.Contains(errMsg, "Current status is CONFIRMED") {
			return &errs.ErrUserAlreadyExits
		}
		return &errs.ErrConfirmEmail
	}
	return nil
}
