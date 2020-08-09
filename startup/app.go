package startup

import (
	"imgo/src/routes"
	"imgo/src/utils"

	"imgo/src/common"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// InitApp func
func InitApp(args []string) error {

	resource := common.NewIMResource(args)
	config, err := resource.Config()
	if err != nil {
		return err
	}

	mySQL, err := resource.MySQLConn()
	if err != nil {
		return err
	}
	if mySQL != nil {
		defer mySQL.Close()
	}

	rabbitMQConn, err := resource.RabbiMQConn()
	if err != nil {
		return err
	}
	if rabbitMQConn != nil {
		defer rabbitMQConn.Close()
	}

	if len(args) <= 1 {
		return setRoute(config, mySQL)
	}
	return nil
}

// setRoute func
func setRoute(config utils.IMConfig, db *gorm.DB) error {
	router := routes.SetupRouter(db)
	return router.Run(config.GetPort())
}
