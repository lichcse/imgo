package resources

import (
	"imgo/app/utils"
	"log"
	"time"

	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var resource imResource

// IMResource interface of resource object
type IMResource interface {
	RabbiMQConn() (*amqp.Connection, error)
	MySQLConn() (*gorm.DB, error)
	Config(args []string) (utils.IMConfig, error)
}

type imResource struct {
	args         []string
	config       utils.IMConfig
	rabbitMQConn *amqp.Connection
	mySQL        *gorm.DB
}

// NewIMResource func new resource object
func NewIMResource() IMResource {
	return &imResource{}
}

// Config func get config data
func (r *imResource) Config(args []string) (utils.IMConfig, error) {
	resource.args = args
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

// RabbiMQConn func get rabbit connection
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

// MySQLConn func get mysql connection
func (r *imResource) MySQLConn() (*gorm.DB, error) {
	if resource.mySQL != nil {
		return resource.mySQL, nil
	}

	sqlConfig := resource.config.MySQLItem("im")
	config := &gorm.Config{}
	if sqlConfig.LogMode {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	conn, err := gorm.Open(mysql.Open(sqlConfig.URL), config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := conn.DB()
	sqlDB.SetMaxOpenConns(sqlConfig.PoolLimit)
	sqlDB.SetMaxIdleConns(sqlConfig.PoolLimit)
	sqlDB.SetConnMaxLifetime(time.Duration(sqlConfig.MaxLifetime) * time.Minute)
	resource.mySQL = conn

	return conn, nil
}
