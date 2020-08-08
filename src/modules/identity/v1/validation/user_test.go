package validation

import (
	"imgo/src/common/identity/v1/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Add(t *testing.T) {
	user := dto.UserAdd{}
	userValidation := NewUserValidation()
	user.FullName = ""
	err := userValidation.Add(user)
	assert.Equal(t, "invalid_full_name", err.Error())

	user.FullName = "Test"
	err = userValidation.Add(user)
	assert.Equal(t, "invalid_username", err.Error())

	user.Username = "test"
	err = userValidation.Add(user)
	assert.Equal(t, "invalid_email", err.Error())

	user.Email = "example@example.com"
	err = userValidation.Add(user)
	assert.Equal(t, "invalid_password", err.Error())

	user.Password = "123456789"
	err = userValidation.Add(user)
	assert.Equal(t, nil, err)
}
