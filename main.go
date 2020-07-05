package main

import (
	"imgo/startup"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	args := os.Args[1:]
	err := startup.InitApp(args)
	if err != nil {
		panic(err)
	}
}
