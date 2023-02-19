package cognito

import (
	"context"
	"strings"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/ramsfords/backend/foundations/errs"
)

func (cognito *CognitoClient) ConfirmEmail(ctx context.Context, email string, code string) error {
	confirmRequest := cip.ConfirmSignUpInput{
		ClientId:         &cognito.CognitoClientID,
		Username:         &email,
		ConfirmationCode: &code,
	}
	_, err := cognito.Client.ConfirmSignUp(ctx, &confirmRequest)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "ExpiredCodeException") {
			return &errs.ErrConfirmEmailExpired
		}
		if strings.Contains(errMsg, "LimitExceededException") {
			return &errs.ErrTooManyRequest
		}
		return &errs.ErrConfirmEmail
	}
	return nil
}
