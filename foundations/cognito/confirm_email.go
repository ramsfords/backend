package cognito

import (
	"context"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cc *CognitoClient) Confirm_Email(ctx context.Context, data ConfirmEmailData) error {
	confirmRequest := cip.ConfirmSignUpInput{
		ClientId:         &data.ClientId,
		Username:         &data.UserName,
		ConfirmationCode: &data.ConfirmationCode,
	}
	_, err := cc.Client.ConfirmSignUp(ctx, &confirmRequest)

	return err
	// if err != nil {
	// 	errMsg := err.Error()
	// 	if strings.Contains(errMsg, "ExpiredCodeException") {
	// 		return &errs.ErrConfirmEmailExpired

	// 	}
	// 	if strings.Contains(errMsg, "LimitExceededException") {
	// 		return &errs.ErrTooManyRequest

	// 	}
	// 	return &errs.ErrConfirmEmail
	// }

	// return nil
}
