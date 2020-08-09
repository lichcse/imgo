package common

import (
	"imgo/src/utils"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var resource imResource

// IMResource interface
type IMResource interface {
	RabbiMQConn() (*amqp.Connection, error)
	MySQLConn() (*gorm.DB, error)
	Config() (utils.IMConfig, error)
}

type imResource struct {
	args         []string
	config       utils.IMConfig
	rabbitMQConn *amqp.Connection
	mySQL        *gorm.DB
}

// NewIMResource func
func NewIMResource(args []string) IMResource {
	resource.args = args
	return &imResource{}
}

func (r *imResource) Config() (utils.IMConfig, error) {
	if resource.config != nil {
		return resource.config, nil
	}

	resource.config = utils.NewIMConfig()
	err := resource.config.Load(resource.args)
	if err != nil {
		return nil, nil
	}

	resource.config.Load(resource.args)
	return resource.config, nil
}

func (r *imResource) RabbiMQConn() (*amqp.Connection, error) {
	if resource.rabbitMQConn != nil && !resource.rabbitMQConn.IsClosed() {
		return resource.rabbitMQConn, nil
	}

	rabbitMQConfig := resource.config.RabbitMQ()
	if rabbitMQConfig.URL == "" {
		return nil, nil
	}

	rab := utils.NewRabbitMQ(rabbitMQConfig.URL)
	rabbitMQConn, err := rab.Connect()
	if err != nil {
		return nil, err
	}

	resource.rabbitMQConn = rabbitMQConn

	go rab.HealthCheck(func(message string) {
		log.Printf("%s", message)
	})
	return rabbitMQConn, nil
}

func (r *imResource) MySQLConn() (*gorm.DB, error) {
	if resource.mySQL != nil {
		return resource.mySQL, nil
	}

	sqlConfig := resource.config.MySQLItem("im")
	conn, err := gorm.Open("mysql", sqlConfig.URL)
	conn.DB().SetMaxOpenConns(sqlConfig.PoolLimit)
	conn.DB().SetMaxIdleConns(sqlConfig.PoolLimit)
	conn.DB().SetConnMaxLifetime(time.Duration(sqlConfig.MaxLifetime) * time.Minute)
	if err != nil {
		return nil, err
	}

	resource.mySQL = conn
	return conn, nil
}
