package main

import (
	"boiler/src/api"
	"boiler/src/api/rest"
	conf "boiler/src/config"
	"boiler/src/db"
	"boiler/src/db/mysql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kowiste/config"
)

func main() {
	err := config.New[conf.BoilerConfig](config.GetPathEnv())
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	_, err = config.Get[conf.BoilerConfig]()
	if err != nil {
		fmt.Println("Error getting config:", err)
		return
	}
	var service api.IAPI = rest.New()
	err = service.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.New(mysql.New())
	database, err := db.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = database.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wait()

}
func wait() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

}
