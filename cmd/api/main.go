package main

import (
	"os"
	"runtime"

	"imgo/app/routes"

	"imgo/app/resources"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args[1:]

	resource := resources.NewIMResource()
	config, err := resource.Config(args, "api.yaml")
	if err != nil {
		panic(err)
	}

	mySQL, err := resource.MySQLConn()
	if err != nil {
		panic(err)
	}

	rabbitMQConn, err := resource.RabbiMQConn()
	if err != nil {
		panic(err)
	}

	if rabbitMQConn != nil {
		defer rabbitMQConn.Close()
	}

	router := routes.SetupRouter(mySQL)
	router.Run(config.GetPort())
}
