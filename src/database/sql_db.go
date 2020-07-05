package database

import (
	"context"
	"database/sql"
)

// SQLDb interface
type SQLDb interface {
	Conn()
	Close()
	Context() (context.Context, func())
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
