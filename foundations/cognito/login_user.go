package cognito

import (
	"context"
	"strings"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/ramsfords/backend/foundations/errs"
)

func (cognito *CognitoClient) LoginUser(ctx context.Context, email string, password string) (*cip.AdminInitiateAuthOutput, error) {
	confirmRequest := cip.AdminInitiateAuthInput{
		AuthFlow:   types.AuthFlowTypeAdminUserPasswordAuth,
		UserPoolId: &cognito.CognitoUserPoolID,
		ClientId:   &cognito.CognitoClientID,
		AuthParameters: map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		},
	}
	loginRes, err := cognito.Client.AdminInitiateAuth(ctx, &confirmRequest)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "ExpiredCodeException") {
			return nil, &errs.ErrConfirmEmailExpired
		}
		if strings.Contains(errMsg, "LimitExceededException") {
			return nil, &errs.ErrTooManyRequest
		}
		return nil, &errs.ErrConfirmEmail
	}
	return loginRes, nil
}
func (cognito *CognitoClient) LoginWithRefreshToken(ctx context.Context, refreshToken string) (*cip.AdminInitiateAuthOutput, error) {
	confirmRequest := cip.AdminInitiateAuthInput{
		AuthFlow:   types.AuthFlowTypeRefreshTokenAuth,
		UserPoolId: &cognito.CognitoUserPoolID,
		ClientId:   &cognito.CognitoClientID,
		AuthParameters: map[string]string{
			"REFRESH_TOKEN": refreshToken,
		},
	}
	loginRes, err := cognito.Client.AdminInitiateAuth(ctx, &confirmRequest)
	if err != nil {
		errMsg := err.Error()
		if strings.Contains(errMsg, "ExpiredCodeException") {
			return nil, &errs.ErrConfirmEmailExpired
		}
		if strings.Contains(errMsg, "LimitExceededException") {
			return nil, &errs.ErrTooManyRequest
		}
		return nil, &errs.ErrConfirmEmail
	}
	return loginRes, nil
}
