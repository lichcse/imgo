package service

import (
	schema "imgo/app/schema/identity/v1"

	"github.com/stretchr/testify/mock"
)

// UserServiceMock struct
type UserServiceMock struct {
	mock.Mock
}

// Add func
func (u *UserServiceMock) Add(userAddRequest *schema.UserAddRequest) (*schema.UserDetailResponse, error) {
	args := u.Called(userAddRequest)
	return args.Get(0).(*schema.UserDetailResponse), args.Error(1)
}

// Detail func
func (u *UserServiceMock) Detail(id uint64) (*schema.UserDetailResponse, error) {
	args := u.Called(id)
	return args.Get(0).(*schema.UserDetailResponse), args.Error(1)
}

// Update func
func (u *UserServiceMock) Update(id string, userUpdateRequest *schema.UserUpdateRequest) error {
	args := u.Called(id, userUpdateRequest)
	return args.Error(0)
}

// Delete func
func (u *UserServiceMock) Delete(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
