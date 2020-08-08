package startup

import (
	"imgo/src/routes"
	"imgo/src/utils"
	"log"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	config utils.IMConfig
)

// InitApp func
func InitApp(args []string) error {

	err := loadConfig(args)
	if err != nil {
		return err
	}

	sqlConfig := config.MySQLItem("im")
	db, err := gorm.Open("mysql", sqlConfig.URL)
	defer db.Close()
	if err != nil {
		return err
	}

	rabbitMQConfig := config.RabbitMQ()
	rab := utils.NewRabbitMQ(rabbitMQConfig.URL)
	rabbitMQConn, err := rab.Connect()
	defer rabbitMQConn.Close()
	if err != nil {
		return err
	}

	go rab.HealthCheck(func(mess string) {
		log.Printf("%s", mess)
	})

	if len(args) <= 1 {
		return setRoute(args, db)
	}
	return nil
}

// loadConfig func
func loadConfig(args []string) error {
	config = utils.NewIMConfig()
	return config.Load(args)
}

// setRoute func
func setRoute(args []string, db *gorm.DB) error {
	config := utils.NewIMConfig()
	router := routes.SetupRouter(db)
	return router.Run(config.GetPort())
}

func callback(data string) error {
	// TODO
	return nil
}
