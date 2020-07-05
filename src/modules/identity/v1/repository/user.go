package repository

import (
	"imgo/src/database"
	"imgo/src/modules/identity/v1/entity"
	"imgo/src/modules/identity/v1/resource"
	"imgo/src/utils"
)

// UserRepository interface
type UserRepository interface {
	Add(userEntity entity.UserEntity) error
	Detail(userID uint64) (entity.UserEntity, error)
	Update(userID string, user entity.UserEntity) error
	Delete(userID string) error
}

type userRepository struct {
	db     database.SQLDb
	imTime utils.IMTime
}

// NewUserRepository func
func NewUserRepository(db database.SQLDb) UserRepository {
	return &userRepository{db: db, imTime: utils.NewIMTime()}
}

// Add func
func (m *userRepository) Add(userEntity entity.UserEntity) error {
	userEntity.Status = resource.UserStatusDefault
	userEntity.CreatedAt = m.imTime.TimeDB()
	userEntity.ModifiedAt = m.imTime.TimeDB()
	_, err := m.db.Exec("INSERT INTO im_user (full_name, username, email, password, created_at, modified_at, status) VALUE (?, ?, ?, ?, ?, ?, ?)", userEntity.FullName, userEntity.Username, userEntity.Email, userEntity.Password, userEntity.CreatedAt, userEntity.ModifiedAt, userEntity.Status)
	return err
}

// Detail func
func (m *userRepository) Detail(userID uint64) (entity.UserEntity, error) {
	result := entity.UserEntity{}
	row := m.db.QueryRow("SELECT * FROM im_user WHERE id=?", userID)
	err := row.Scan(&result.ID, &result.FullName, &result.Username, &result.Email, &result.Password, &result.CreatedAt, &result.ModifiedAt, &result.Status)
	return result, err
}

// Update func
func (m *userRepository) Update(userID string, user entity.UserEntity) error {
	return nil
}

// Delete func
func (m *userRepository) Delete(userID string) error {
	return nil
}
