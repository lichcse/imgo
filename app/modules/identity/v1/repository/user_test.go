package repository

import (
	"errors"
	"testing"

	"imgo/app/modules/identity/v1/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func UserRepositoryPrepareDataTest() entity.User {
	user := entity.User{
		Email:    "example@example.com",
		FullName: "Test add user",
		Username: "example",
		Password: "123456789",
		Status:   0,
	}
	return user
}

func TestUserRepository_Add_Fail(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer db.Close()

	user := UserRepositoryPrepareDataTest()

	userRepository := NewUserRepository(mockDB)
	mock.ExpectExec("^INSERT*").WillReturnError(errors.New("Test - Add_Fail"))
	err := userRepository.Add(&user)
	assert.Equal(t, "Test - Add_Fail", err.Error())
}

func TestUserRepository_Add_Success(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer db.Close()

	user := UserRepositoryPrepareDataTest()
	userRepository := NewUserRepository(mockDB)
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT*").WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()
	err := userRepository.Add(&user)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(2), user.ID)
}

func TestUserRepository_Detail_Fail(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer db.Close()
	userRepository := NewUserRepository(mockDB)

	mock.ExpectQuery("^SELECT*").WillReturnError(errors.New("Test - Detail_Fail"))
	_, err := userRepository.Detail(100)
	assert.Equal(t, "Test - Detail_Fail", err.Error())
}

func TestUserRepository_Detail_Success(t *testing.T) {
	mockDB, mock, db := MockDB(t)
	defer db.Close()
	userRepository := NewUserRepository(mockDB)

	rows := sqlmock.NewRows([]string{"id", "full_name", "username", "email", "password", "created_at", "modified_at", "status"}).
		AddRow(2, "Test - Detail_Success", "test", "example@example.com", "123456789", "2020-08-09", "2020-08-09", 1)

	mock.ExpectQuery("^SELECT*").WillReturnRows(rows)
	user, err := userRepository.Detail(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(2), user.ID)
	assert.Equal(t, "Test - Detail_Success", user.FullName)
}
