package service

import (
	"imgo/src/common/identity/v1/dto"

	"github.com/stretchr/testify/mock"
)

// UserServiceMock struct
type UserServiceMock struct {
	mock.Mock
}

// Add func
func (u *UserServiceMock) Add(userAddDTO dto.UserAdd) error {
	args := u.Called(userAddDTO)
	return args.Error(0)
}

// Detail func
func (u *UserServiceMock) Detail(userID uint64) (dto.UserResponse, error) {
	args := u.Called(userID)
	return args.Get(0).(dto.UserResponse), args.Error(1)
}

// Update func
func (u *UserServiceMock) Update(userID string, userUpdateDTO dto.UserUpdate) error {
	args := u.Called(userID, userUpdateDTO)
	return args.Error(0)
}

// Delete func
func (u *UserServiceMock) Delete(userID string) error {
	args := u.Called(userID)
	return args.Error(0)
}
