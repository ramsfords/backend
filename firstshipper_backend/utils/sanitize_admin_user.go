package utils

import v1 "github.com/ramsfords/types_gen/v1"

func SanitizeAdminUser(req *v1.User) (userDb v1.User, err error) {
	// hash the user password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return v1.SignUp{}, err
	// }
	// userId := uuid.New().String()[0:6]
	// if err != nil {
	// 	return v1.SignUp{}, err
	// }
	// businessId, err := bcrypt.GenerateFromPassword([]byte(req.CompanyName), bcrypt.MinCost)
	// if err != nil {
	// 	return err
	// }
	dbUser := v1.User{
		Name:  req.Name,
		Email: req.Email,
	}
	// dbv1.BusinessIds = []string{}
	// if req.PhoneNumber != "" {
	// 	dbv1.PhoneNumbers = append(dbv1.PhoneNumbers, req.PhoneNumber)
	// }
	// emailVerfiedStr := ""
	// // fix for now, While not using cognito. email_verified = false
	// if emailVerfiedStr == "UNCONFIRMED" || emailVerfiedStr == "" {
	// 	dbv1.EmailVerified = false
	// } else {
	// 	dbv1.EmailVerified = true
	// }
	return dbUser, nil
}
