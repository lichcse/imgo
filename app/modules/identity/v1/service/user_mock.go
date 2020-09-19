package service

import (
	"imgo/app/common/identity/v1/dto"

	"github.com/stretchr/testify/mock"
)

// UserServiceMock struct
type UserServiceMock struct {
	mock.Mock
}

// Add func
func (u *UserServiceMock) Add(userAddRequest *dto.UserAddRequest) (*dto.UserDetailResponse, error) {
	args := u.Called(userAddRequest)
	return args.Get(0).(*dto.UserDetailResponse), args.Error(1)
}

// Detail func
func (u *UserServiceMock) Detail(id uint64) (*dto.UserDetailResponse, error) {
	args := u.Called(id)
	return args.Get(0).(*dto.UserDetailResponse), args.Error(1)
}

// Update func
func (u *UserServiceMock) Update(id string, userUpdateRequest *dto.UserUpdateRequest) error {
	args := u.Called(id, userUpdateRequest)
	return args.Error(0)
}

// Delete func
func (u *UserServiceMock) Delete(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
