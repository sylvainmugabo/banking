package main

import (
	"github.com/sylvainmugabo/microservices-lib/logger"
	"github.com/sylvainmugabo/microservices/banking/app"
	"github.com/sylvainmugabo/microservices/banking/config"
)

func main() {
	config := config.LoadEnvVariables()

	logger.Info("Starting the application")
	app.Start(config)
}
