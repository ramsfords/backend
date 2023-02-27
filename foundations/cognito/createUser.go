package cognito

import (
	"context"
	"fmt"
	"strings"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/ramsfords/backend/foundations/errs"
	v1 "github.com/ramsfords/types_gen/v1"
)

func (cognito CognitoClient) CreateUser(ctx context.Context, data *v1.User) (*cip.AdminCreateUserOutput, error) {
	signUpUser := &cip.AdminCreateUserInput{
		UserPoolId:    &cognito.CognitoUserPoolID,
		Username:      aws.String(data.Email),
		MessageAction: types.MessageActionTypeSuppress,
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
				Name:  aws.String("custom:organizationId"),
				Value: aws.String(data.OrganizationId),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(data.PhoneNumber),
			},
		},
	}
	signUpOutPut, err := cognito.Client.AdminCreateUser(ctx, signUpUser)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "UsernameExistsException") {
			return nil, errs.ErrUserAlreadyExits
		}
	}
	setPasswordOutPut, err := cognito.SetUserPassword(ctx, data)
	fmt.Println(setPasswordOutPut)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "UsernameExistsException") {
			return nil, errs.ErrUserAlreadyExits
		}
		return nil, err
	}
	return signUpOutPut, nil
}
func (cognito CognitoClient) SetUserPassword(ctx context.Context, data *v1.User) (*cip.AdminSetUserPasswordOutput, error) {
	signUpUser := &cip.AdminSetUserPasswordInput{
		UserPoolId: &cognito.CognitoUserPoolID,
		Username:   aws.String(data.Email),
		Password:   aws.String(data.Password),
		Permanent:  true,
	}
	signUpOutPut, err := cognito.Client.AdminSetUserPassword(ctx, signUpUser)
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "UsernameExistsException") {
			return nil, errs.ErrUserAlreadyExits
		}
		return nil, err
	}
	return signUpOutPut, nil
}
