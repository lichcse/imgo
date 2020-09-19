package repository

import (
	"errors"
	"testing"

	"imgo/src/modules/identity/v1/entity"

	"github.com/DATA-DOG/go-sqlmock"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func UserRepositoryPrepareDataTest() entity.User {
	user := entity.User{
		Email:    "example@example.com",
		FullName: "Test",
		Username: "example",
		Password: "123456789",
		Status:   0,
	}
	return user
}

func TestUserRepository_Add_Fail(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer mockDB.Close()
	defer db.Close()

	user := UserRepositoryPrepareDataTest()

	userRepository := NewUserRepository(db)
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT*").WillReturnError(errors.New("Test"))
	mock.ExpectCommit()
	err := userRepository.Add(&user)
	assert.Equal(t, "Test", err.Error())
}

func TestUserRepository_Add_Success(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer mockDB.Close()
	defer db.Close()

	user := UserRepositoryPrepareDataTest()
	userRepository := NewUserRepository(db)
	// success
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT*").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := userRepository.Add(&user)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), user.ID)
}

func TestUserRepository_Detail(t *testing.T) {
	mockDB, mock, db := MockDB(t)
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
