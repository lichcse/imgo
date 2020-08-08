package validation

import (
	"imgo/src/common/identity/v1/dto"

	"github.com/stretchr/testify/mock"
)

// UserValidationMock struct
type UserValidationMock struct {
	mock.Mock
}

// Add func
func (r *UserValidationMock) Add(userAdd dto.UserAdd) error {
	args := r.Called(userAdd)
	return args.Error(0)
}
