package validation

import (
	"errors"
	"imgo/app/common/identity/v1/dto"
	"imgo/app/utils"
)

// UserValidation func
type UserValidation interface {
	Add(userAdd *dto.UserAddRequest) error
}

type userValidation struct {
	validation utils.IMValidation
}

// NewUserValidation func
func NewUserValidation() UserValidation {
	return &userValidation{validation: utils.NewIMValidation()}
}

// Add func
func (u *userValidation) Add(userAddRequest *dto.UserAddRequest) error {

	if len(userAddRequest.FullName) < 3 || len(userAddRequest.FullName) > 50 {
		return errors.New("user_invalid_full_name")
	}

	validUsername := u.validation.IsValidUsername(userAddRequest.Username)
	if !validUsername {
		return errors.New("user_invalid_username")
	}

	validEmail := u.validation.IsValidEmail(userAddRequest.Email)
	if !validEmail {
		return errors.New("user_invalid_email")
	}

	if len(userAddRequest.Password) == 0 {
		return errors.New("user_invalid_password")
	}
	return nil
}
