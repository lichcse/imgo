package service

import (
	"imgo/app/modules/identity/v1/entity"
	"imgo/app/modules/identity/v1/repository"
	schema "imgo/app/schema/identity/v1"
	"imgo/app/utils"
)

// UserService interface of user service object
type UserService interface {
	Add(userAddRequest *schema.UserAddRequest) (*schema.UserDetailResponse, error)
	Detail(id uint64) (*schema.UserDetailResponse, error)
	Update(id string, userUpdate *schema.UserUpdateRequest) error
	Delete(id string) error
}

type userService struct {
	userRepo repository.UserRepository
	crypt    utils.IMCrypt
	convert  utils.IMConvert
}

// NewUserService func new user service object
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		crypt:    utils.NewIMCrypt(),
		convert:  utils.NewIMConvert(),
	}
}

// Add func add new user
func (u *userService) Add(userAddRequest *schema.UserAddRequest) (*schema.UserDetailResponse, error) {
	result := &schema.UserDetailResponse{}
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

// Detail func get detail user info
func (u *userService) Detail(id uint64) (*schema.UserDetailResponse, error) {
	result := &schema.UserDetailResponse{}
	user, err := u.userRepo.Detail(id)
	if err != nil {
		return result, u.convert.DatabaseError(err)
	}

	err = u.convert.Object(user, &result)
	return result, err
}

// Update func update user info
func (u *userService) Update(id string, userUpdate *schema.UserUpdateRequest) error {
	// TODO
	return u.userRepo.Update(id, &entity.User{})
}

// Delete func delete user info
func (u *userService) Delete(id string) error {
	// TODO
	return u.userRepo.Delete(id)
}
