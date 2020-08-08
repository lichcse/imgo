package service

import (
	"errors"
	"imgo/src/common/identity/v1/dto"
	"imgo/src/modules/identity/v1/repository"
	"testing"

	"imgo/src/modules/identity/v1/entity"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Add(t *testing.T) {
	userAddDTO := dto.UserAdd{
		FullName: "Test",
		Username: "test",
		Email:    "example@example.com",
		Password: "123456789",
	}

	userEntity := entity.User{
		FullName: userAddDTO.FullName,
		Username: userAddDTO.Username,
		Email:    userAddDTO.Email,
		Password: "",
	}

	// fail
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Add", &userEntity).Return(errors.New("Test"))
	err := userService.Add(userAddDTO)
	assert.Equal(t, "Test", err.Error())

	// success
	userRepository = new(repository.UserRepositoryMock)
	userService = NewUserService(userRepository)
	userRepository.On("Add", &userEntity).Return(nil)
	err = userService.Add(userAddDTO)
	assert.Equal(t, nil, err)
}

func TestUserService_Detail(t *testing.T) {
	userAddDTO := dto.UserResponse{
		FullName: "Test",
		Username: "test",
		Email:    "example@example.com",
	}

	userEntity := entity.User{
		FullName: userAddDTO.FullName,
		Username: userAddDTO.Username,
		Email:    userAddDTO.Email,
		Password: "",
	}

	// fail
	userRepository := new(repository.UserRepositoryMock)
	userService := NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(userEntity, errors.New("undefined"))
	_, err := userService.Detail(uint64(1))
	assert.Equal(t, "undefined", err.Error())

	// success
	userRepository = new(repository.UserRepositoryMock)
	userService = NewUserService(userRepository)
	userRepository.On("Detail", uint64(1)).Return(userEntity, nil)
	user, err := userService.Detail(uint64(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, "Test", user.FullName)
}
