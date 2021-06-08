package main

import (
	"github.com/bugg123/rest-golang-microservices-udemy/app"
	"github.com/bugg123/rest-golang-microservices-udemy/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
