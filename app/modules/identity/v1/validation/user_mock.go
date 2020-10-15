package validation

import (
	schema "imgo/app/schema/identity/v1"

	"github.com/stretchr/testify/mock"
)

// UserValidationMock struct
type UserValidationMock struct {
	mock.Mock
}

// Add func
func (r *UserValidationMock) Add(userAddRequest *schema.UserAddRequest) error {
	args := r.Called(userAddRequest)
	return args.Error(0)
}
