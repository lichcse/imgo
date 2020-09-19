package service

import (
	"imgo/src/common/identity/v1/dto"
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/repository"
	"imgo/src/utils"
)

// UserService interface
type UserService interface {
	Add(userAddRequest *dto.UserAddRequest) (*dto.UserDetailResponse, error)
	Detail(id uint64) (*dto.UserDetailResponse, error)
	Update(id string, userUpdate *dto.UserUpdateRequest) error
	Delete(id string) error
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
func (u *userService) Add(userAddRequest *dto.UserAddRequest) (*dto.UserDetailResponse, error) {
	result := &dto.UserDetailResponse{}
	hash, _ := u.crypt.Hash(userAddRequest.Password)
	user := &entity.User{
		FullName: userAddRequest.FullName,
		Username: userAddRequest.Username,
		Email:    userAddRequest.Email,
		Password: hash,
	}

	err := u.userRepo.Add(user)
	if err != nil {
		return result, err
	}

	err = u.convert.Object(user, &result)
	return result, err
}

// Detail func
func (u *userService) Detail(id uint64) (*dto.UserDetailResponse, error) {
	result := &dto.UserDetailResponse{}
	user, err := u.userRepo.Detail(id)
	if err != nil {
		return result, u.convert.DatabaseError(err)
	}

	err = u.convert.Object(user, &result)
	return result, err
}

// Update func
func (u *userService) Update(id string, userUpdate *dto.UserUpdateRequest) error {
	// TODO
	return u.userRepo.Update(id, &entity.User{})
}

// Delete func
func (u *userService) Delete(id string) error {
	// TODO
	return u.userRepo.Delete(id)
}
