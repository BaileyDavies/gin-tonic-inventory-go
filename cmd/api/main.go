package main

import (
	"gin-tonic-inventory-go/cmd/api/router"
	"gin-tonic-inventory-go/pkg/application"
	"gin-tonic-inventory-go/pkg/exithandler"
	"gin-tonic-inventory-go/pkg/logger"
	"gin-tonic-inventory-go/pkg/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// godotenv gets the env vars we defined in our local .env file
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load env vars")
	}

	app, err := application.Get()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.
		Get().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(router.Get(app)).
		WithErrLogger(logger.Error)

	go func() {
		logger.Info.Printf("starting server on %s", app.Cfg.GetAPIPort())
		if err := srv.Start(); err != nil {
			logger.Error.Fatal(err.Error())
		}
	}()

	exithandler.Init(func() {
		if err := app.DB.Close(); err != nil {
			log.Println(err.Error())
		}
	})
}
