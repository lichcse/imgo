package utils

import (
	"fmt"
	"log"

	"github.com/globalsign/mgo"
)

// global mongo session
var gMongoSession = make(map[string]*mgo.Session)

// MongoDB interface
type MongoDB interface {
	S() (*mgo.Session, error)
	DB() (*mgo.Database, func())
	C(name string) (*mgo.Collection, func())
	Conn() error
}

type mongoDB struct {
	config MongoConfig
}

// NewMongoDB func
// This function will be returned to the interface of MongoDB
func NewMongoDB(config MongoConfig) MongoDB {
	return &mongoDB{config: config}
}

// Conn func
// This function will try to connect to the database
func (m *mongoDB) Conn() error {
	var err error
	var gmsTmp *mgo.Session

	if gmsTmp, err = mgo.Dial(m.config.URL); err == nil {
		gmsTmp.SetPoolLimit(m.config.PoolLimit)
		gMongoSession[m.config.Database] = gmsTmp
	}

	if err != nil {
		log.Fatal(fmt.Sprintf("Connect to mongo database at %s fail (%s).", m.config.URL, err.Error()))
	}

	return err
}

// S func
// This function will be returned to the session
func (m *mongoDB) S() (*mgo.Session, error) {
	s, ok := gMongoSession[m.config.Database]
	if ok && s != nil {
		return s.Clone(), nil
	}

	err := m.Conn()
	if err == nil && s != nil {
		return s.Clone(), nil
	}

	return s, err
}

// DB func
// This function will be returned to the database
func (m *mongoDB) DB() (*mgo.Database, func()) {
	s, err := m.S()
	if err == nil {
		return s.DB(m.config.Database), s.Close
	}
	panic(fmt.Sprintf("Can not connect to %s (%s).", m.config.URL, err.Error()))
}

// C func
// This function will be returned to the collection
func (m *mongoDB) C(name string) (*mgo.Collection, func()) {
	db, cleanup := m.DB()
	return db.C(name), cleanup
}
