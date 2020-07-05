package database

import (
	"context"
	"imgo/src/utils"
	"testing"
	"time"

	. "github.com/stretchr/testify/assert"
)

// func TestMySQL_DB(t *testing.T) {
// 	mySQL := NewMySQL(utils.MySQLConfig{URL: "root:root123456@/im", Database: "im", PoolLimit: 200})
// 	db1 := mySQL.DB()
// 	defer db1.Close()
// 	dbx := &sql.DB{}
// 	IsType(t, dbx, db1)
// 	db2 := mySQL.DB()
// 	IsType(t, dbx, db2)
// 	db2.Close()
// }

func TestMySQL_Context(t *testing.T) {
	mySQL := NewMySQL(utils.MySQLConfig{URL: "root:root123456@/im", Database: "im", PoolLimit: 200})
	ctx, cancel := mySQL.Context()
	defer cancel()
	ctxTest, cancelTest := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelTest()
	IsType(t, ctxTest, ctx)
	cancelTestX := func() {

	}
	IsType(t, cancelTestX, cancel)
}
