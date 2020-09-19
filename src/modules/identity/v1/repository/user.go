package repository

import (
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/resource"
	"imgo/src/utils"

	"github.com/jinzhu/gorm"
)

// UserRepository interface
type UserRepository interface {
	Add(user *entity.User) error
	Detail(id uint64) (*entity.User, error)
	Update(id string, user *entity.User) error
	Delete(id string) error
}

type userRepository struct {
	db     *gorm.DB
	imTime utils.IMTime
}

// NewUserRepository func
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db, imTime: utils.NewIMTime()}
}

// Add func
func (u *userRepository) Add(user *entity.User) error {
	user.Status = resource.UserStatusDefault
	user.CreatedAt = u.imTime.TimeDB()
	user.ModifiedAt = u.imTime.TimeDB()
	return u.db.Create(&user).Error
}

// Detail func
func (u *userRepository) Detail(id uint64) (*entity.User, error) {
	result := &entity.User{}
	err := u.db.Where("id = ?", id).First(result).Error
	return result, err
}

// Update func
func (u *userRepository) Update(id string, user *entity.User) error {
	// TODO
	user.ModifiedAt = u.imTime.TimeDB()
	return nil
}

// Delete func
func (u *userRepository) Delete(id string) error {
	// TODO
	return nil
}
