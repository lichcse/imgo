package validation

import (
	"errors"
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/utils"
)

// UserValidation func
type UserValidation interface {
	Add(userAddDto entity.UserAddDTO) error
}

type userValidation struct {
	validation utils.IMValidation
}

// NewUserValidation func
func NewUserValidation() UserValidation {
	return &userValidation{validation: utils.NewIMValidation()}
}

// Add func
func (m *userValidation) Add(userAddDto entity.UserAddDTO) error {

	if len(userAddDto.FullName) < 3 || len(userAddDto.FullName) > 50 {
		return errors.New("invalid_full_name")
	}

	validUsername := m.validation.IsValidUsername(userAddDto.Username)
	if !validUsername {
		return errors.New("invalid_username")
	}

	validEmail := m.validation.IsValidEmail(userAddDto.Email)
	if !validEmail {
		return errors.New("invalid_email")
	}

	if len(userAddDto.Password) == 0 {
		return errors.New("invalid_password")
	}
	return nil
}
