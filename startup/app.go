package startup

import (
	"imgo/app/routes"
	"imgo/app/utils"

	"imgo/app/resources"

	"gorm.io/gorm"
)

// InitApp func init app resource
func InitApp(args []string) error {

	resource := resources.NewIMResource()
	config, err := resource.Config(args)
	if err != nil {
		return err
	}

	mySQL, err := resource.MySQLConn()
	if err != nil {
		return err
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

// setRoute func setup app route
func setRoute(config utils.IMConfig, db *gorm.DB) error {
	router := routes.SetupRouter(db)
	return router.Run(config.GetPort())
}
