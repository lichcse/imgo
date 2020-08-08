package service

import (
	"imgo/src/common/identity/v1/dto"
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/repository"
	"imgo/src/utils"
)

// UserService interface
type UserService interface {
	Add(userAddDTO dto.UserAdd) error
	Detail(userID uint64) (dto.UserResponse, error)
	Update(userID string, userUpdateDTO dto.UserUpdate) error
	Delete(userID string) error
}

type userService struct {
	userRepo repository.UserRepository
	crypt    utils.IMCrypt
	convert  utils.IMConvert
}

// NewUserService func
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		crypt:    utils.NewIMCrypt(),
		convert:  utils.NewIMConvert(),
	}
}

// Add func
func (u *userService) Add(userAddDTO dto.UserAdd) error {
	hash, _ := u.crypt.Hash(userAddDTO.Password)

	return u.userRepo.Add(&entity.User{
		FullName: userAddDTO.FullName,
		Username: userAddDTO.Username,
		Email:    userAddDTO.Email,
		Password: hash,
	})
}

// Detail func
func (u *userService) Detail(userID uint64) (dto.UserResponse, error) {
	userResponseDto := dto.UserResponse{}
	user, err := u.userRepo.Detail(userID)
	if err != nil {
		return userResponseDto, u.convert.DatabaseError(err)
	}

	err = u.convert.Object(user, &userResponseDto)
	return userResponseDto, err
}

// Update func
func (u *userService) Update(userID string, userUpdateDTO dto.UserUpdate) error {
	return u.userRepo.Update(userID, &entity.User{})
}

// Delete func
func (u *userService) Delete(userID string) error {
	return u.userRepo.Delete(userID)
}
