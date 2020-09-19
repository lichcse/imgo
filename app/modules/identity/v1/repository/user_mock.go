package repository

import (
	"imgo/app/modules/identity/v1/entity"

	"github.com/stretchr/testify/mock"
)

// UserRepositoryMock struct
type UserRepositoryMock struct {
	mock.Mock
}

// Add func
func (u *UserRepositoryMock) Add(user *entity.User) error {
	user.Password = ""
	args := u.Called(user)
	return args.Error(0)
}

// Detail func
func (u *UserRepositoryMock) Detail(id uint64) (*entity.User, error) {
	args := u.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}

// Update func
func (u *UserRepositoryMock) Update(id string, user *entity.User) error {
	args := u.Called(id, user)
	return args.Error(0)
}

// Delete func
func (u *UserRepositoryMock) Delete(id string) error {
	args := u.Called(id)
	return args.Error(0)
}
