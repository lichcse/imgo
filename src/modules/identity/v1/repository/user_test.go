package repository

import (
	"database/sql"
	"errors"
	"testing"

	"imgo/src/modules/identity/v1/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func UserInit(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		assert.Equal(t, "TestUser - error sqlmock", err.Error())
	}
	db, err := gorm.Open("mysql", mockDB)
	if err != nil {
		assert.Equal(t, "TestUser - error gorm", err.Error())
	}

	return mockDB, mock, db
}

func TestUser_Add_Fail(t *testing.T) {
	mockDB, mock, db := UserInit(t)
	defer mockDB.Close()
	defer db.Close()

	user := entity.User{
		Email:    "example@example.com",
		FullName: "Test",
		Username: "example",
		Password: "123456789",
		Status:   0,
	}
	userRepository := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT*").WillReturnError(errors.New("Test"))
	mock.ExpectCommit()
	err := userRepository.Add(&user)
	assert.Equal(t, "Test", err.Error())
}

func TestUser_Add_Success(t *testing.T) {
	mockDB, mock, db := UserInit(t)
	defer mockDB.Close()
	defer db.Close()

	user := entity.User{
		Email:    "example@example.com",
		FullName: "Test",
		Username: "example",
		Password: "123456789",
		Status:   0,
	}
	userRepository := NewUserRepository(db)
	// success
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := userRepository.Add(&user)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), user.ID)
}

func TestUser_Detail(t *testing.T) {
	mockDB, mock, db := UserInit(t)
	defer mockDB.Close()
	defer db.Close()
	userRepository := NewUserRepository(db)

	// fail
	mock.ExpectQuery("^SELECT*").WillReturnError(errors.New("Test"))
	user, err := userRepository.Detail(1)
	assert.Equal(t, "Test", err.Error())

	// success
	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "password", "created_at", "modified_at", "status"}).
		AddRow(1, "Test", "test", "example@example.com", "123456789", "2020-08-09", "2020-08-09", 1)

	mock.ExpectQuery("^SELECT*").WillReturnRows(rows)
	user, err = userRepository.Detail(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), user.ID)
	assert.Equal(t, "Test", user.FullName)
}
