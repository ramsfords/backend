package cognito

import (
	"context"

	cip "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func (cc CognitoClient) Create_User(signUpReq interface{}, ctx context.Context) (*cip.SignUpOutput, error) {
	// signUpUser := &cip.SignUpInput{
	// 	ClientId: &cc.Conf.ClientId,
	// 	Username: aws.String(signUpReq.Email),
	// 	Password: aws.String(signUpReq.Password),
	// 	UserAttributes: []types.AttributeType{
	// 		{
	// 			Name:  aws.String("email"),
	// 			Value: aws.String(signUpReq.Email),
	// 		},
	// 		{
	// 			Name:  aws.String("preferred_username"),
	// 			Value: aws.String(signUpReq.Email),
	// 		},
	// 		{
	// 			Name:  aws.String("given_name"),
	// 			Value: aws.String(signUpReq.GivenName()),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:org_id"),
	// 			Value: aws.String("org#" + signUpReq.Email),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:sort_id"),
	// 			Value: aws.String("user#" + signUpReq.Email),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:phone_numbers"),
	// 			Value: aws.String("+1-713-516-2836"),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:roles"),
	// 			Value: aws.String(models.Roles.String()),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:issued_by"),
	// 			Value: aws.String("userservice"),
	// 		},
	// 		{
	// 			Name:  aws.String("custom:user_updated_locally"),
	// 			Value: aws.String("false"),
	// 		},
	// 	},
	// }
	// for _, j := range signUpUser.UserAttributes {
	// 	if *j.Name == "custom:phone_numbers" && signUpReq.PhoneNumber != "" {
	// 		*j.Value = signUpReq.PhoneNumber
	// 		break
	// 	}

	// }
	// signUpOutPut, err := cc.Client.SignUp(ctx, signUpUser)
	// if err != nil {
	// 	// msg := err.Error()
	// 	// if strings.Contains(msg, "UsernameExistsException") {
	// 	// 	return nil, errs.ErrUserAlreadyExits
	// 	// }
	// 	return nil, err
	// }
	// return signUpOutPut, nil
	return nil, nil
}
