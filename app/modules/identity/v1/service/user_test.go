package service

import (
	"errors"
	"imgo/app/modules/identity/v1/repository"
	schema "imgo/app/schema/identity/v1"
	"testing"

	"imgo/app/modules/identity/v1/entity"

	"github.com/stretchr/testify/assert"
)

func UserServicePrepareDataTest() (*schema.UserAddRequest, *entity.User) {
	userAddRequest := &schema.UserAddRequest{
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

func TestUserService_Add_Fail(t *testing.T) {
	userAddRequest, user := UserServicePrepareDataTest()
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Add", user).Return(errors.New("Add_Fail"))
	_, err := userService.Add(userAddRequest)
	assert.Equal(t, "Add_Fail", err.Error())
}

func TestUserService_Add_Success(t *testing.T) {
	userAddRequest, user := UserServicePrepareDataTest()
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Add", user).Return(nil)
	_, err := userService.Add(userAddRequest)
	assert.Equal(t, nil, err)
}

func TestUserService_Detail_Fail(t *testing.T) {
	_, user := UserServicePrepareDataTest()
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(user, errors.New("undefined"))
	_, err := userService.Detail(uint64(1))
	assert.Equal(t, "undefined", err.Error())
}

func TestUserService_Detail_Success(t *testing.T) {
	userAddRequest, user := UserServicePrepareDataTest()
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(user, nil)
	userDetailResponse, err := userService.Detail(uint64(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, userAddRequest.FullName, userDetailResponse.FullName)
}
