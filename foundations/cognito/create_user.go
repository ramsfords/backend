package cognito

import (
	"context"
	"strings"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/foundations/error"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (cognito CognitoClient) CreateUser(ctx context.Context, data *v1.User) (*cip.SignUpOutput, error) {
	signUpUser := &cip.SignUpInput{
		ClientId: &cognito.CognitoClientID,
		Username: aws.String(data.Email),
		Password: aws.String(data.Password),
		UserAttributes: []types.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(data.Email),
			},
			{
				Name:  aws.String("preferred_username"),
				Value: aws.String(data.Email),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(data.Name),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(data.PhoneNumber),
			},
		},
	}
	signUpOutPut, err := cognito.Client.SignUp(ctx, signUpUser)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "UsernameExistsException") {
			return nil, error.ErrUserAlreadyExits
		}
		return nil, err
	}
	return signUpOutPut, nil
}
