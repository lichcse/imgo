package utils

import (
	"testing"

	"github.com/globalsign/mgo"
	. "github.com/stretchr/testify/assert"
)

func initTestMongoConfig(db string, url string) MongoConfig {
	if len(db) <= 0 {
		db = "test"
	}
	if len(url) <= 0 {
		url = "mongodb://localhost:27017"
	}
	return MongoConfig{
		URL:       url,
		Database:  db,
		PoolLimit: 200,
	}
}

func TestMongoDB_ConnSuccess(t *testing.T) {
	mongoDB := NewMongoDB(initTestMongoConfig("test_conn", ""))
	err := mongoDB.Conn()
	Equal(t, nil, err)
}

func TestMongoDB_S(t *testing.T) {
	mongoConfig := initTestMongoConfig("test_s", "")
	mongoDB := NewMongoDB(mongoConfig)
	s, err := mongoDB.S()
	Equal(t, nil, err)
	NotEqual(t, nil, s)
}

func TestMongoDB_DB(t *testing.T) {
	mongoConfig := initTestMongoConfig("test_db", "")
	mongoDB := NewMongoDB(mongoConfig)
	s, close := mongoDB.DB()
	closeFunc := func() {

	}
	db := mgo.Database{}
	IsType(t, &db, s)
	IsType(t, closeFunc, close)
}

func TestMongoDB_C(t *testing.T) {
	mongoConfig := initTestMongoConfig("test_c", "")
	mongoDB := NewMongoDB(mongoConfig)
	c, close := mongoDB.C("test")
	closeFunc := func() {

	}
	collection := mgo.Collection{}
	IsType(t, &collection, c)
	IsType(t, closeFunc, close)
}
