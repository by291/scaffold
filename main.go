package main

import (
	"by291.fun/scaffold/global"
	initialize "by291.fun/scaffold/init"
	"by291.fun/scaffold/router"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	initialize.Init()

	conf := &global.Config
	r := router.Routers()

	go func() {
		err := r.Run(":" + strconv.Itoa(conf.App.Port))

		// if you use tls, uncomment code bene following
		// you must add cert and key file path in config.yaml

		// err := r.RunTLS(
		// 	":"+strconv.Itoa(conf.App.Port),
		// 	conf.App.CertFile,
		// 	conf.App.KeyFile,
		// )
		if err != nil {
			panic(err)
		}
	}()

	// let the application quit when received a signal
	quit := make(chan os.Signal)
	// corresponding to CTRL + C in Win and kill in Linux
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
