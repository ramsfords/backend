package utils

import v1 "github.com/ramsfords/types_gen/v1"

func SanitizeUser(req *v1.AddStaff) *v1.User {
	newUserData := &v1.User{
		UserName: req.NewStaffEmail,
		Email:    req.NewStaffEmail,
		// Username:      req.NewStaffEmail,
		// CreatedAt:     time.Now().String(),
		// BusinessId:    []string{req.BusinessId},
		// EmailVerified: false,
	}
	// hash the user password

	return newUserData
}
