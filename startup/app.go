package startup

import (
	"imgo/src/database"
	"imgo/src/routes"
	"imgo/src/utils"
)

// InitApp func
func InitApp(args []string) error {

	err := loadConfig(args)
	if err != nil {
		return err
	}

	config := utils.NewIMConfig()
	sqlConfig := config.MySQLItem("im")
	mySQL := database.NewMySQL(sqlConfig)
	mySQL.Conn()
	defer mySQL.Close()

	if len(args) <= 1 {
		return setRoute(args, mySQL)
	}
	return nil
}

// loadConfig func
func loadConfig(args []string) error {
	config := utils.NewIMConfig()
	return config.Load(args)
}

// setRoute func
func setRoute(args []string, db database.SQLDb) error {
	config := utils.NewIMConfig()
	router := routes.SetupRouter(db)
	return router.Run(config.GetPort())
}
