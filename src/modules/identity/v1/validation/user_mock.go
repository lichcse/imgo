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
func (r *UserValidationMock) Add(userAddRequest *dto.UserAddRequest) error {
	args := r.Called(userAddRequest)
	return args.Error(0)
}
