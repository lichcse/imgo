package validation

import (
	"imgo/app/common/identity/v1/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Add(t *testing.T) {
	userAddRequest := &dto.UserAddRequest{}
	userValidation := NewUserValidation()
	userAddRequest.FullName = ""
	err := userValidation.Add(userAddRequest)
	assert.Equal(t, "user_invalid_full_name", err.Error())

	userAddRequest.FullName = "Test"
	err = userValidation.Add(userAddRequest)
	assert.Equal(t, "user_invalid_username", err.Error())

	userAddRequest.Username = "test"
	err = userValidation.Add(userAddRequest)
	assert.Equal(t, "user_invalid_email", err.Error())

	userAddRequest.Email = "example@example.com"
	err = userValidation.Add(userAddRequest)
	assert.Equal(t, "user_invalid_password", err.Error())

	userAddRequest.Password = "123456789"
	err = userValidation.Add(userAddRequest)
	assert.Equal(t, nil, err)
}
