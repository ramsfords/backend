package cognito

import (
	"context"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cc CognitoClient) ForgotPassword(ctx context.Context, email string) (*cip.ForgotPasswordOutput, error) {
	input := cip.ForgotPasswordInput{
		ClientId: &cc.Conf.CognitoClientID,
		Username: &email,
	}
	output, err := cc.Client.ForgotPassword(ctx, &input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
