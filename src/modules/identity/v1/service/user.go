package service

import (
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/repository"
	"imgo/src/utils"
)

// UserService interface
type UserService interface {
	Add(userAddDto entity.UserAddDTO) error
	Detail(userID uint64) (entity.UserResponseDTO, error)
	Update(userID string, user entity.UserEntity) error
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
func (m *userService) Add(userAddDto entity.UserAddDTO) error {
	hash, _ := m.crypt.Hash(userAddDto.Password)
	return m.userRepo.Add(entity.UserEntity{
		FullName: userAddDto.FullName,
		Username: userAddDto.Username,
		Email:    userAddDto.Email,
		Password: hash,
	})
}

// Detail func
func (m *userService) Detail(userID uint64) (entity.UserResponseDTO, error) {
	userResponseDto := entity.UserResponseDTO{}
	user, err := m.userRepo.Detail(userID)
	if err != nil {
		return userResponseDto, m.convert.DatabaseError(err)
	}

	err = m.convert.Object(user, &userResponseDto)
	return userResponseDto, err
}

// Update func
func (m *userService) Update(userID string, user entity.UserEntity) error {
	return m.userRepo.Update(userID, user)
}

// Delete func
func (m *userService) Delete(userID string) error {
	return m.userRepo.Delete(userID)
}
