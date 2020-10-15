package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

// MockDB func sql mock
func MockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New()
	if err != nil {
		assert.Equal(t, "TestUser - error sqlmock", err.Error())
	}

	mockDB, err := gorm.Open(mysql.Dialector{Config: &mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}}, &gorm.Config{})
	if err != nil {
		assert.Equal(t, "TestUser - error gorm", err.Error())
	}

	return mockDB, mock, db
}
