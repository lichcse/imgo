package database

import (
	"context"
	"database/sql"
	"imgo/src/utils"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type mySQL struct {
	config utils.MySQLConfig
	db     *sql.DB
}

// NewMySQL func
func NewMySQL(config utils.MySQLConfig) SQLDb {
	return &mySQL{config: config}
}

// DB func
func (m *mySQL) Conn() {
	db, err := sql.Open("mysql", m.config.URL)
	db.SetConnMaxLifetime(time.Second * time.Duration(60))
	db.SetMaxIdleConns(m.config.PoolLimit)
	db.SetMaxOpenConns(m.config.PoolLimit)
	if err != nil {
		panic(err.Error())
	}
	m.db = db
}

// Close func
func (m *mySQL) Close() {
	m.db.Close()
}

// Context func
func (m *mySQL) Context() (context.Context, func()) {
	return context.WithTimeout(context.Background(), 15*time.Millisecond)
}

// Query func
func (m *mySQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.QueryContext(context.Background(), query, args...)
}

// Exec func
func (m *mySQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.db.ExecContext(context.Background(), query, args...)
}

// QueryRow func
func (m *mySQL) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.db.QueryRowContext(context.Background(), query, args...)
}
