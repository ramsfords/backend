package cognito

import (
	"context"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cc CognitoClient) ResetPassword(ctx context.Context, email string, newPassword string) (*cip.AdminSetUserPasswordOutput, error) {
	input := cip.AdminSetUserPasswordInput{
		Password:  &newPassword,
		Username:  &email,
		Permanent: true,
	}
	output, err := cc.Client.AdminSetUserPassword(ctx, &input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
