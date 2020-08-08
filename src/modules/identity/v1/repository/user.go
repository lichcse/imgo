package repository

import (
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/resource"
	"imgo/src/utils"

	"github.com/jinzhu/gorm"
)

// UserRepository interface
type UserRepository interface {
	Add(userEntity *entity.User) error
	Detail(userID uint64) (entity.User, error)
	Update(userID string, user *entity.User) error
	Delete(userID string) error
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
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}
	return err
}

// Detail func
func (u *userRepository) Detail(userID uint64) (entity.User, error) {
	result := entity.User{}
	err := u.db.Where("id = ?", userID).First(&result).Error
	return result, err
}

// Update func
func (u *userRepository) Update(userID string, user *entity.User) error {
	return nil
}

// Delete func
func (u *userRepository) Delete(userID string) error {
	return nil
}
