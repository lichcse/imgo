package validation

import (
	"errors"
	"imgo/src/common/identity/v1/dto"
	"imgo/src/utils"
)

// UserValidation func
type UserValidation interface {
	Add(userAdd dto.UserAdd) error
}

type userValidation struct {
	validation utils.IMValidation
}

// NewUserValidation func
func NewUserValidation() UserValidation {
	return &userValidation{validation: utils.NewIMValidation()}
}

// Add func
func (u *userValidation) Add(userAdd dto.UserAdd) error {

	if len(userAdd.FullName) < 3 || len(userAdd.FullName) > 50 {
		return errors.New("invalid_full_name")
	}

	validUsername := u.validation.IsValidUsername(userAdd.Username)
	if !validUsername {
		return errors.New("invalid_username")
	}

	validEmail := u.validation.IsValidEmail(userAdd.Email)
	if !validEmail {
		return errors.New("invalid_email")
	}

	if len(userAdd.Password) == 0 {
		return errors.New("invalid_password")
	}
	return nil
}
