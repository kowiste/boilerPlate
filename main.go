package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	conf "github.com/kowiste/boilerplate/src/config"
	"github.com/kowiste/boilerplate/src/core"

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
		return
	}
	err = core.Init()
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
