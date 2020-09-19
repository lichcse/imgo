package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

// MockDB func sql mock
func MockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *gorm.DB) {
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
