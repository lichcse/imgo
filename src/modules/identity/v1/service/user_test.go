package service

import (
	"errors"
	"imgo/src/common/identity/v1/dto"
	"imgo/src/modules/identity/v1/repository"
	"testing"

	"imgo/src/modules/identity/v1/entity"

	"github.com/stretchr/testify/assert"
)

func UserServicePrepareDataTest() (*dto.UserAddRequest, *entity.User) {
	userAddRequest := &dto.UserAddRequest{
		FullName: "Test",
		Username: "test",
		Email:    "example@example.com",
		Password: "123456789",
	}

	user := &entity.User{
		FullName: userAddRequest.FullName,
		Username: userAddRequest.Username,
		Email:    userAddRequest.Email,
		Password: "",
	}

	return userAddRequest, user
}

func TestUserService_Add(t *testing.T) {
	userAddRequest, user := UserServicePrepareDataTest()
	// fail
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Add", user).Return(errors.New("Add-Error"))
	_, err := userService.Add(userAddRequest)
	assert.Equal(t, "Add-Error", err.Error())

	// success
	userRepository = new(repository.UserRepositoryMock)
	userService = NewUserService(userRepository)
	userRepository.On("Add", user).Return(nil)
	_, err = userService.Add(userAddRequest)
	assert.Equal(t, nil, err)
}

func TestUserService_Detail(t *testing.T) {
	userAddRequest, user := UserServicePrepareDataTest()
	// fail
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(user, errors.New("undefined"))
	_, err := userService.Detail(uint64(1))
	assert.Equal(t, "undefined", err.Error())

	// success
	userRepository = new(repository.UserRepositoryMock)
	userService = NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(user, nil)
	userDetailResponse, err := userService.Detail(uint64(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, userAddRequest.FullName, userDetailResponse.FullName)
}
