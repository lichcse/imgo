package repository

import (
	"imgo/app/modules/identity/v1/entity"
	"imgo/app/utils"

	"gorm.io/gorm"
)

// UserRepository interface of user repository object
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

// NewUserRepository func new user repository object
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db, imTime: utils.NewIMTime()}
}

// Add func add new user
func (u *userRepository) Add(user *entity.User) error {
	user.Status = entity.UserStatusDefault
	user.CreatedAt = u.imTime.TimeDB()
	user.ModifiedAt = u.imTime.TimeDB()
	return u.db.Create(&user).Error
}

// Detail func get detail user info
func (u *userRepository) Detail(id uint64) (*entity.User, error) {
	result := &entity.User{}
	err := u.db.Where("id = ?", id).First(result).Error
	return result, err
}

// Update func update user info
func (u *userRepository) Update(id string, user *entity.User) error {
	// TODO
	user.ModifiedAt = u.imTime.TimeDB()
	return nil
}

// Delete func delete user info
func (u *userRepository) Delete(id string) error {
	// TODO
	return nil
}
